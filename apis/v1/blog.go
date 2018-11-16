package v1

import (
	"blog_gin/db"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func ListBlog(c *gin.Context){
	db.Listblog()
}
