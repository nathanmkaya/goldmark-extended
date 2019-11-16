package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

type Citation struct {
	gast.BaseInline
}

func (c *Citation) Kind() gast.NodeKind {
	return KindCitation
}

func (c *Citation) Dump(source []byte, level int) {
	gast.DumpHelper(c, source, level, nil, nil)
}

var KindCitation = gast.NewNodeKind("Citation")

func NewCitation() *Citation {
	return &Citation{}
}
