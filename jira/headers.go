package jira

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	. "github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"strconv"
)

type heading struct {
}

// NewHeadingParser returns a new BlockParser that parses
// jira Headings
func NewHeadingParser() BlockParser {
	return &heading{}
}

func (h *heading) Trigger() []byte {
	return []byte{'h'}
}

func (h *heading) Open(parent ast.Node, reader text.Reader, pc Context) (ast.Node, State) {
	line, segment := reader.PeekLine()
	pos := pc.BlockOffset()
	if pos < 0 {
		return nil, NoChildren
	}
	pos++
	if pos > len(line)-1 {
		return nil, NoChildren
	}
	level, _ := strconv.Atoi(string(line[pos]))
	if level > 6 {
		return nil, NoChildren
	}

	node := ast.NewHeading(level)

	i := pos
	for ; i < len(line) && line[i] != ' '; i++ {
	}

	l := util.TrimLeftSpaceLength(line[i:])
	if l == 0 {
		return nil, NoChildren
	}

	start := i + l
	stop := len(line) - util.TrimRightSpaceLength(line)

	node.Lines().Append(text.NewSegment(segment.Start+start, segment.Start+stop))
	return node, NoChildren
}

func (h *heading) Continue(node ast.Node, reader text.Reader, pc Context) State {
	return Close
}

func (h *heading) Close(node ast.Node, reader text.Reader, pc Context) {
}

func (h *heading) CanInterruptParagraph() bool {
	return true
}

func (h *heading) CanAcceptIndentedLine() bool {
	return false
}

var Heading = &heading{}

func (h *heading) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		WithBlockParsers(
			util.Prioritized(NewHeadingParser(), 100),
		),
	)
}
