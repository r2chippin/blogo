package controller

import "github.com/gin-gonic/gin"

func HandleBlogs(c *gin.Context) {
	path := "./resources/blogFiles/"
	param := c.Param("param")

	html := RenderMarkdown(path, param)

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, string(html))
}
