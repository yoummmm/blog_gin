package main

import (
	db "gin_blog/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8000")
}
