package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func init() {
	var err error
<<<<<<< HEAD
	SqlDB, err := sql.Open("mysql", "test:Test123!@#@tcp(119.3.11.44:3306)/blog?parseTime=true")
=======
	SqlDB, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/blog?parseTime=true")
>>>>>>> jikeSen
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
