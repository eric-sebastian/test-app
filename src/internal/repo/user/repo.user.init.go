package user

import (
	"log"
	"strconv"
	"test-app/src/internal/entity/user_model"
	"test-app/src/internal/repo"
)

var thisUser user_model.User
var thisUsers []user_model.User

//New :
func New() {
	repo.New()
}

//GetUsers :
func GetUsers() []user_model.User {
	New()
	db := repo.DB
	defer db.Close()

	thisUsers = nil
	rows, err := db.Query("SELECT * FROM users")
	handleError(err)
	for rows.Next() {
		err = rows.Scan(&thisUser.ID, &thisUser.Name)
		handleError(err)

		thisUsers = append(thisUsers, thisUser)
	}

	return thisUsers
}

//GetUser :
func GetUser(id string) user_model.User {
	New()
	db := repo.DB
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE id = ?", id)
	handleError(err)
	for rows.Next() {
		err = rows.Scan(&thisUser.ID, &thisUser.Name)
		handleError(err)
	}

	return thisUser
}

//CreateUser :
func CreateUser(newUser user_model.UserForm) user_model.User {
	New()
	db := repo.DB
	defer db.Close()

	tx, err := db.Begin()
	handleError(err)

	res, err := tx.Exec("INSERT INTO users (name) VALUES (?)", newUser.Name)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	handleError(err)

	thisUser.ID = int(id)
	thisUser.Name = newUser.Name
	handleError(tx.Commit())

	return thisUser
}

//UpdateUser :
func UpdateUser(id string, updateUser user_model.UserForm) user_model.User {
	New()
	db := repo.DB
	defer db.Close()

	tx, err := db.Begin()
	handleError(err)

	res, err := tx.Exec("UPDATE users SET name = ? WHERE id = ?", updateUser.Name, id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	if res == nil {

	}

	sid, err := strconv.Atoi(id)
	handleError(err)

	thisUser.ID = sid
	thisUser.Name = updateUser.Name
	handleError(tx.Commit())

	return thisUser
}

//DeleteUser :
func DeleteUser(id string) {
	New()
	db := repo.DB
	defer db.Close()

	tx, err := db.Begin()
	handleError(err)

	res, err := tx.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	if res == nil {

	}
	handleError(tx.Commit())
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
