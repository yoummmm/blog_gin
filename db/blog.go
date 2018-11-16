package db

import (
	"blog_gin/models"
	"github.com/goinggo/mapstructure"
	"log"
)

func Listblog() ([]models.Blog,error) {
	res, err := engin.Query("select * from blog")
	if err != nil{
		log.Fatal("listblog",err.Error())
	}
	var result Result
	errr := mapstructure.Decode(res,&result)
	
	if errr != nil{
		log.Fatal("listblog",err.Error())
	}
	return result,err
}
