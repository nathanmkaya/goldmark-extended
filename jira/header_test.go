package jira

import (
	"github.com/nathanmkaya/goldmark-extended/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHeadingParser(t *testing.T) {
	assert := assert.New(t)
	header1 := "h1. Biggest heading"
	assert.Equal("<h1>Biggest heading</h1>\n", testutil.ParseExtension(header1, Jira))
}
