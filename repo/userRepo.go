package repo

import (
	"log"

	"github.com/aj9mb/task-management/constants"
	"github.com/aj9mb/task-management/dbmg"
	"github.com/aj9mb/task-management/model"
)

func UserAdd(u *model.User) (*model.User, error) {
	db := dbmg.GetDb()
	stmt, err := db.Prepare(constants.USER_ADD)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	res, err := stmt.Exec(u.UserName, u.Password, u.Name)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	u.Id = id
	u.Password = ""
	return u, err
}

func GetUserByUserName(username string) (*model.User, error) {
	db := dbmg.GetDb()
	stmt, err := db.Prepare(constants.USER_GET)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	rows, err := stmt.Query(username)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	user := new(model.User)
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func CheckUserNameExist(usernames []string) ([]string, error) {
	query := constants.USER_NAME_CHECK
	for i, v := range usernames {
		query += "'" + v + "'"
		if i < len(usernames)-1 {
			query += ","
		}
	}
	query += ")"

	db := dbmg.GetDb()
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	existUserNames := make([]string, 0)
	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
		existUserNames = append(existUserNames, username)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return existUserNames, err
}
