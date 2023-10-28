package utils

import "database/sql"

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang?parseTime=true")
	if err != nil {
		panic(err)
	}

	return db
}
