package ui

import (
	markdown "github.com/MichaelMure/go-term-markdown"
)

func RenderMarkdown(source string) string {
	return string(markdown.Render(source, 80, 0))
}
