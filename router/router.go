package router

import (
	"blog_gin/apis/v1"
	"blog_gin/apis/v2"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
		
	//创建路由器组建 组合相关的相同前缀的接口
	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/",v1.HomeApi)
		apiv1.GET("listblog",v1.ListBlog)
	}
		
	//创建不同版本的接口
	apiv2 := router.Group("/api/v2")
	{
		apiv2.GET("/",v2.HomeApi)
	}
	
	return router
}
