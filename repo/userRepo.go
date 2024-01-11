package repo

import (
	"github.com/aj9mb/task-management/constants"
	"github.com/aj9mb/task-management/dbmg"
	"github.com/aj9mb/task-management/logging"
	"github.com/aj9mb/task-management/model"
)

func UserAdd(u *model.User) (*model.User, error) {
	db := dbmg.GetDb()
	stmt, err := db.Prepare(constants.USER_ADD)
	if err != nil {
		logging.GetLogger().Print(err)
		return nil, err
	}
	res, err := stmt.Exec(u.UserName, u.Password, u.Name)
	if err != nil {
		logging.GetLogger().Print(err)
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		logging.GetLogger().Print(err)
		return nil, err
	}
	u.Id = id
	u.Password = ""
	return u, err
}

func GetUserByUserName(username string) (*model.User, error) {
	db := dbmg.GetDb()
	user := new(model.User)
	err := db.QueryRow(constants.USER_GET, username).Scan(&user.Id, &user.UserName, &user.Password, &user.Name)
	if err != nil {
		logging.GetLogger().Print(err)
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
		logging.GetLogger().Print(err)
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		logging.GetLogger().Print(err)
		return nil, err
	}
	defer rows.Close()
	existUserNames := make([]string, 0)
	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			logging.GetLogger().Print(err)
		} else {
			existUserNames = append(existUserNames, username)
		}
	}
	err = rows.Err()
	if err != nil {
		logging.GetLogger().Print(err)
	}
	return existUserNames, err
}
