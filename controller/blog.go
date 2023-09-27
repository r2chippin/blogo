package controller

import "github.com/gin-gonic/gin"

func HandleBlogs(c *gin.Context) {
	param := c.Param("param")

	html := RenderMarkdown(param)

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, string(html))
}
