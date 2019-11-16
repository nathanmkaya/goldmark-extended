package jira

import (
	"github.com/nathanmkaya/goldmark-extended/jira/ast"
	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	. "github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type citationDelimiter struct {
}

func (c citationDelimiter) IsDelimiter(b byte) bool {
	return b == '?'
}

func (c citationDelimiter) CanOpenCloser(opener, closer *Delimiter) bool {
	return opener.Char == closer.Char
}

func (c citationDelimiter) OnMatch(consumes int) gast.Node {
	return ast.NewCitation()
}

var defaultCitationDelimiter = &citationDelimiter{}

type citationParser struct {
}

func (c *citationParser) Trigger() []byte {
	return []byte{'?'}
}

func (c *citationParser) Parse(parent gast.Node, block text.Reader, pc Context) gast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := ScanDelimiter(line, before, 2, defaultCitationDelimiter)
	if node == nil {
		return nil
	}
	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}

var defaultCitationParser = &citationParser{}

func NewCitationParser() InlineParser {
	return defaultCitationParser
}

type CitationHTMLRenderer struct {
	html.Config
}

func (c *CitationHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindCitation, c.renderCitation)
}

func (r *CitationHTMLRenderer) renderCitation(w util.BufWriter, souce []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		w.WriteString("<cite>")
	} else {
		w.WriteString("</cite>")
	}
	return gast.WalkContinue, nil
}

func NewCitationHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &CitationHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

type citation struct {
}

var Citation = &citation{}

func (c *citation) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(WithInlineParsers(
		util.Prioritized(NewCitationParser(), 500)))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewCitationHTMLRenderer(), 500)))
}
