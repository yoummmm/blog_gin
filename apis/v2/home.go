package v2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeApi(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  nil,
		"data":   "if you get this message,the apis for version v2 is ok",
	})
}
