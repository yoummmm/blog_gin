package v1

import (
	"blog_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//博客列表
func ListBlog(c *gin.Context)  {
		B := models.Blog{}
		res,err := B.ListBlogs(0)
		if err != nil{
			c.JSON(http.StatusOK, gin.H{
				"res": res,
			})
		}
}
