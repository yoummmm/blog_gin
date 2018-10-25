package main

import (
	"github.com/gin-gonic/gin"
	."blog_gin/apis"
	)


func initRouter() *gin.Engine  {
	router:=gin.Default()
	router.GET("/",IndexApi)
	router.POST("/person", AddPersonApi)
	router.GET("/blog", ListBlog)
/*	router.GET("/persons", GetPersonsApi)
	router.GET("/person/:id", GetPersonApi)
	router.PUT("/person/:id", ModPersonApi)
	router.DELETE("/person/:id", DelPersonApi)*/
	return router
}