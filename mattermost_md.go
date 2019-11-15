package goldmark_extended

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

type mattermost struct {
}

var MM  = &mattermost{}

func (mm *mattermost) Extend(md goldmark.Markdown) {
	extension.Linkify.Extend(md)
	extension.Table.Extend(md)
	extension.Strikethrough.Extend(md)
	extension.TaskList.Extend(md)
}