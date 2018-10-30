package main

import (
	db "blog_gin/database"
	"blog_gin/pkg/setting"
	"blog_gin/router"
	"fmt"
	"net/http"
)

func main() {
	defer db.SqlDB.Close()

	router := router.InitRouter()

	ser := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.HttpPort),
		Handler:           router,
		ReadHeaderTimeout: setting.ReadTimeOut,
		WriteTimeout:      setting.WriteTimeOut,
	}
	ser.ListenAndServe()
}
