package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListBlog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   "hi list",
	})
}
