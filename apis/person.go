package apis

import (
	"blog_gin/models"
	. "blog_gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "it is works")
}
func GetPersonsApi(c *gin.Context) {
	p := models.Person{}
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}
func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	p := Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
