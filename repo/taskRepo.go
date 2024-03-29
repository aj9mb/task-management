package repo

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/aj9mb/task-management/constants"
	"github.com/aj9mb/task-management/dbmg"
	"github.com/aj9mb/task-management/logging"
	"github.com/aj9mb/task-management/model"
)

func TaskAdd(t *model.Task) (*model.Task, error) {
	db := dbmg.GetDb()
	stmt, err := db.Prepare(constants.TASK_ADD)
	if err != nil {
		logging.GetLogger().Print(err)
		return nil, err
	}
	res, err := stmt.Exec(t.BoardId, t.AddedBy, t.Assignee, t.TaskDesc)
	if err != nil {
		logging.GetLogger().Print(err)
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	t.Id = id
	return t, err
}

func GetTaskList(boardId int64, userId int64) (*[]model.Task, error) {
	taskList := make([]model.Task, 0)
	db := dbmg.GetDb()
	stmt, err := db.Prepare(constants.GET_TASK_LIST)
	if err != nil {
		logging.GetLogger().Print(err)
		return nil, err
	}
	rows, err := stmt.Query(boardId, userId)
	if err != nil {
		logging.GetLogger().Print(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		task := new(model.Task)
		var createDt sql.NullTime
		var lastUpt sql.NullTime
		err := rows.Scan(&task.Id, &task.BoardId, &task.AddedBy, &task.AddedByName, &task.Assignee, &task.AssigneeName, &task.TaskDesc, &task.Status, &createDt, &lastUpt)
		if err != nil {
			logging.GetLogger().Print(err)
		} else {
			if createDt.Valid {
				task.CreateDt = &createDt.Time
			}
			if lastUpt.Valid {
				task.LastUpt = &lastUpt.Time
			}
			taskList = append(taskList, *task)
		}
	}
	err = rows.Err()
	if err != nil {
		logging.GetLogger().Print(err)
	}
	return &taskList, err
}

func UpdateTask(task *model.Task) (bool, error) {
	db := dbmg.GetDb()
	stmt, err := db.Prepare(createTaskUpdateQuery(task))
	if err != nil {
		logging.GetLogger().Print(err)
		return false, err
	}
	res, err := stmt.Exec(task.Id)
	if err != nil {
		logging.GetLogger().Print(err)
		return false, err
	}
	cnt, err := res.RowsAffected()
	if err != nil {
		logging.GetLogger().Print(err)
		return false, err
	}
	return cnt > 0, err
}

func createTaskUpdateQuery(task *model.Task) string {
	placeholder := ""
	if task.Assignee != nil && *task.Assignee > 0 {
		placeholder += " assigned_to = " + strconv.FormatInt(*task.Assignee, 10) + ","
	}
	if task.TaskDesc != nil && *task.TaskDesc != "" {
		placeholder += " task_desc = '" + *task.TaskDesc + "',"
	}
	if task.Status != nil {
		placeholder += " status = " + strconv.FormatBool(*task.Status) + ","
	}
	if placeholder != "" && placeholder[len(placeholder)-1:] == "," {
		placeholder = placeholder[:len(placeholder)-1]
	}
	return strings.Replace(constants.TASK_UPDATE, "$PLACEHOLDER", placeholder, 1)
}
