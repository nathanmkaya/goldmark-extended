package jira

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/yuin/goldmark"
	"log"
	"testing"
)

func result(sourceString string) string {
	var buf bytes.Buffer
	var source = []byte(sourceString)
	md := goldmark.New(goldmark.WithExtensions(Jira))
	if err := md.Convert(source, &buf); err != nil {
		log.Fatalln(err)
	}
	return buf.String()
}

func TestNewHeadingParser(t *testing.T) {
	assert := assert.New(t)
	header1 := "h1. Biggest heading"
	assert.Equal("<h1>Biggest heading</h1>\n", result(header1))
}
