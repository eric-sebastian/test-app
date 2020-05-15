package repo

import (
	"database/sql"
	//go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql"
)

//DB :
var DB *sql.DB

//New :
func New() {
	InitDB("root:1234567890@tcp(127.0.0.1:3306)/Phonebook")
}

//InitDB :
func InitDB(source string) {
	var err error
	DB, err = sql.Open("mysql", source)
	if err != nil {
		panic(err.Error())
	}
}
