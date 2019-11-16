package jira

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	. "github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type strongDelimiter struct {
}

func (s strongDelimiter) IsDelimiter(b byte) bool {
	return b == '*'
}

func (s strongDelimiter) CanOpenCloser(opener, closer *Delimiter) bool {
	return opener.Char == closer.Char
}

func (s strongDelimiter) OnMatch(consumes int) ast.Node {
	return ast.NewEmphasis(2)
}

var defaultStrongDelimiter = &strongDelimiter{}

type strong struct {
}

func (s *strong) Trigger() []byte {
	return []byte{'*'}
}

func (s *strong) Parse(parent ast.Node, block text.Reader, pc Context) ast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := ScanDelimiter(line, before, 1, defaultStrongDelimiter)
	if node == nil {
		return nil
	}
	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}

// NewStrongParser returns a InlineParser that parses strong
func NewStrongParser() InlineParser {
	return &strong{}
}

var Strong = &strong{}

func (s *strong) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		WithInlineParsers(
			util.Prioritized(NewStrongParser(), 200)))
}
