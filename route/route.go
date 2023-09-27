package route

import (
	"blogo/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Static("/static/", "./static")

	router.GET("/blogs/:param", controller.HandleBlogs)

	return router
}
