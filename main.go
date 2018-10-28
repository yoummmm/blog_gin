package main

import (
	db "blog_gin/database"
	"blog_gin/pkg/setting"
	"fmt"
	"net/http"
	"blog_gin/router"
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
