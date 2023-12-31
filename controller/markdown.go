package controller

import (
	"github.com/russross/blackfriday/v2"
	"os"
)

func RenderMarkdown(path string, param string) []byte {
	// Read markdown file
	markdownBytes, err := os.ReadFile(path + param + ".md")
	if err != nil {
		return []byte("Failed to read Markdown file")
	}

	// Turn markdown file to html
	html := blackfriday.Run(markdownBytes)

	// 返回 HTML 响应
	return html
}
