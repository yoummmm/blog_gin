package main

import (
	"github.com/gin-gonic/gin"
	. "blog_gin/apis"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexApi)
	router.GET("/persons", GetPersonsApi)
	router.POST("/person", AddPersonApi)
	return router
}
