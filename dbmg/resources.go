package dbmg

import "database/sql"

var db *sql.DB

func GetDb() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/task_management?parseTime=true")
		if err != nil {
			panic(err)
		}
	}
	return db
}
