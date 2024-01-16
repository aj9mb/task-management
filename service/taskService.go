package service

import (
	"net/http"
	"strconv"
	"time"

	"github.com/aj9mb/task-management/logging"
	"github.com/aj9mb/task-management/model"
	"github.com/aj9mb/task-management/repo"
	"github.com/labstack/echo/v4"
)

func TaskAdd(c echo.Context) error {
	url := c.Request().URL.RequestURI()
	userId, err := strconv.ParseInt(c.Request().Header.Get("x-user-id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid user id", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	boardId, err := strconv.Atoi(c.Param("board_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid board id", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	if exist, err := repo.CheckBoardExist(int64(boardId)); err != nil || !exist {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "board does not exist", Url: url, StatusCode: http.StatusInternalServerError, Time: time.Now()})
	}
	task := new(model.Task)
	c.Bind(task)
	if task.Assignee == nil || *task.Assignee == 0 || task.TaskDesc == nil || *task.TaskDesc == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	task.BoardId = int64(boardId)
	task.AddedBy = userId
	task, err = repo.TaskAdd(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "Error adding task", Url: url, StatusCode: http.StatusInternalServerError, Time: time.Now()})
	}
	return c.JSON(http.StatusOK, task)
}

func TaskListGet(c echo.Context) error {
	url := c.Request().URL.RequestURI()
	userId, err := strconv.ParseInt(c.Request().Header.Get("x-user-id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid user id", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	boardId, err := strconv.Atoi(c.Param("board_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid board id", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	if exist, err := repo.CheckBoardExist(int64(boardId)); err != nil || !exist {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "board does not exist", Url: url, StatusCode: http.StatusInternalServerError, Time: time.Now()})
	}
	taskList, err := repo.GetTaskList(int64(boardId), userId)
	if err != nil || taskList == nil || len(*taskList) < 1 {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "No tasks found", Url: url, StatusCode: http.StatusInternalServerError, Time: time.Now()})
	}
	return c.JSON(http.StatusOK, taskList)
}

func TaskUpdate(c echo.Context) error {
	url := c.Request().URL.RequestURI()
	taskId, err := strconv.ParseInt(c.Param("task_id"), 10, 64)
	if err != nil || taskId < 1 {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "Invalid task id", Url: url, StatusCode: http.StatusInternalServerError, Time: time.Now()})
	}
	task := new(model.Task)
	c.Bind(task)
	if task.Assignee == nil && task.TaskDesc == nil && task.Status == nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "bad request", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	task.Id = taskId
	res, err := repo.UpdateTask(task)
	if err != nil {
		logging.GetLogger().Print(err)
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "Internal Server error", Url: url, StatusCode: http.StatusInternalServerError, Time: time.Now()})
	}
	if !res {
		return c.JSON(http.StatusNotModified, model.ErrorResponse{Message: "no update", Url: url, StatusCode: http.StatusNotModified, Time: time.Now()})
	}
	return c.JSON(http.StatusOK, task)
}
