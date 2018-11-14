package main

import (
	"blog_gin/pkg/setting"
	"blog_gin/router"
	"fmt"
	"github.com/go-xorm/xorm"
	"net/http"
)

var Engine *xorm.Engine

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
