package router

import (
	"github.com/gin-gonic/gin"
	. "blog_gin/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexApi)
	router.GET("/persons", GetPersonsApi)
	router.POST("/person", AddPersonApi)
	router.GET("/listBlog", ListBlog)
	return router
}
