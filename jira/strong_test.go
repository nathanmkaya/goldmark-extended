package jira

import (
	"github.com/nathanmkaya/goldmark-extended/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStrongParser(t *testing.T) {
	assert := assert.New(t)
	strong := "*strong*"
	assert.Equal("<p><strong>strong</strong></p>\n", testutil.ParseExtension(strong, Jira))
}
