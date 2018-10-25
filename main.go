package main

import (
	db "blog_gin/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8001")
}
