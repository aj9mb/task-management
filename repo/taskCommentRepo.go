package repo

import (
	"database/sql"

	"github.com/aj9mb/task-management/constants"
	"github.com/aj9mb/task-management/dbmg"
	"github.com/aj9mb/task-management/logging"
	"github.com/aj9mb/task-management/model"
)

func TaskCommentAdd(taskComment *model.TaskComment) (*model.TaskComment, error) {
	db := dbmg.GetDb()
	stmt, err := db.Prepare(constants.TASK_COMMENT_ADD)
	if err != nil {
		logging.GetLogger().Print(err)
		return nil, err
	}
	res, err := stmt.Exec(taskComment.TaskId, taskComment.UserId, *taskComment.Comment)
	if err != nil {
		logging.GetLogger().Print(err)
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		logging.GetLogger().Print(err)
		return nil, err
	}
	taskComment.Id = id
	return taskComment, err
}

func TaskCommentListGet(taskId int64) (*[]model.TaskComment, error) {
	db := dbmg.GetDb()
	stmt, err := db.Prepare(constants.TASK_COMMENT_LIST_GET)
	if err != nil {
		logging.GetLogger().Print(err)
		return nil, err
	}
	rows, err := stmt.Query(taskId)
	if err != nil {
		logging.GetLogger().Print(err)
		return nil, err
	}
	defer rows.Close()
	taskCommentList := make([]model.TaskComment, 0)
	for rows.Next() {
		taskComment := new(model.TaskComment)
		var createDt sql.NullTime
		var lastUpt sql.NullTime
		err := rows.Scan(&taskComment.Id, &taskComment.TaskId, &taskComment.UserId, &taskComment.UserName, &taskComment.Comment, &createDt, &lastUpt)
		if err != nil {
			logging.GetLogger().Print(err)
		} else {
			if createDt.Valid {
				taskComment.CreateDt = &createDt.Time
			}
			if lastUpt.Valid {
				taskComment.LastUpt = &lastUpt.Time
			}
			taskCommentList = append(taskCommentList, *taskComment)
		}
	}
	err = rows.Err()
	if err != nil {
		logging.GetLogger().Print(err)
	}
	return &taskCommentList, err
}
