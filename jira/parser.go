package jira

import "github.com/yuin/goldmark"

type jira struct {
}

var Jira = &jira{}

func (j *jira) Extend(m goldmark.Markdown) {
	Heading.Extend(m)
	Strong.Extend(m)
	Citation.Extend(m)
}
