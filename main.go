package main

import (
	"blog_gin/pkg/setting"
	"blog_gin/router"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)


func main() {
	
	router := router.InitRouter()
	
	ser := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.HttpPort),
		Handler:           router,
		ReadHeaderTimeout: setting.ReadTimeOut,
		WriteTimeout:      setting.WriteTimeOut,
	}
	ser.ListenAndServe()
}
