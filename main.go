package main

import (
	db "blog_gin/database"
	"blog_gin/pkg/setting"
	"fmt"
	"net/http"
)

func main() {
	defer db.SqlDB.Close()

	router := initRouter()

	ser := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.HttpPort),
		Handler:           router,
		ReadHeaderTimeout: setting.ReadTimeOut,
		WriteTimeout:      setting.WriteTimeOut,
	}
	ser.ListenAndServe()
}
