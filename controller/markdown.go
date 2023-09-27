package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"os"
)

func RenderMarkdown(c *gin.Context) {
	// 读取 Markdown 文件
	markdownBytes, err := os.ReadFile("./static/test.md")
	if err != nil {
		c.String(500, "Failed to read Markdown file")
		return
	}

	// 将 Markdown 转换为 HTML
	html := blackfriday.Run(markdownBytes)

	// 返回 HTML 响应
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, string(html))
}
