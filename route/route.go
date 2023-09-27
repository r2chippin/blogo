package route

import (
	"blogo/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Static("/statics/", "./resources/statics")

	router.GET("/blogs/:param", controller.HandleBlogs)

	return router
}
