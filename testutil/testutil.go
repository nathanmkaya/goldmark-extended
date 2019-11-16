package testutil

import (
	"bytes"
	"github.com/yuin/goldmark"
	"log"
)

func ParseExtension(sourceString string, ext ...goldmark.Extender) string {
	var buf bytes.Buffer
	var source = []byte(sourceString)
	md := goldmark.New(goldmark.WithExtensions(ext...))
	if err := md.Convert(source, &buf); err != nil {
		log.Fatalln(err)
	}
	return buf.String()
}
