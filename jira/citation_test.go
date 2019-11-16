package jira

import (
	"github.com/nathanmkaya/goldmark-extended/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCitationParser(t *testing.T) {
	assert := assert.New(t)
	citation := `??citation??`
	assert.Equal("<p><cite>citation</cite></p>\n", testutil.ParseExtension(citation, Jira))
}
