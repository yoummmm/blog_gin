package models

import (
	db "gin_blog/database"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson() (id int64, err error) {

	rs, err := db.SqlDB.Exec("insert into person(first_name,last_name) values(?,?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}

	id, err = rs.LastInsertId()
	return
}

func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("select id,first_name,last_name from person")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}

	if err = rows.Err(); err != nil {
		return
	}
	return
}
