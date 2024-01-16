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

func AddTaskComment(c echo.Context) error {
	url := c.Request().URL.RequestURI()
	userId, err := strconv.ParseInt(c.Request().Header.Get("x-user-id"), 10, 64)
	if err != nil || userId < 1 {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid user id", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	taskId, err := strconv.ParseInt(c.Param("task_id"), 10, 64)
	if err != nil || taskId < 1 {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "Invalid task id", Url: url, StatusCode: http.StatusInternalServerError, Time: time.Now()})
	}
	taskComment := new(model.TaskComment)
	c.Bind(taskComment)
	if taskComment.Comment == nil || *taskComment.Comment == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Comment cannot be empty or null", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	taskComment.TaskId = taskId
	taskComment.UserId = userId
	taskComment, err = repo.TaskCommentAdd(taskComment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "Error adding comment", Url: url, StatusCode: http.StatusInternalServerError, Time: time.Now()})
	}
	return c.JSON(http.StatusOK, taskComment)
}

func GetTaskCommentList(c echo.Context) error {
	url := c.Request().URL.RequestURI()
	taskId, err := strconv.ParseInt(c.Param("task_id"), 10, 64)
	if err != nil || taskId < 1 {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "Invalid task id", Url: url, StatusCode: http.StatusInternalServerError, Time: time.Now()})
	}
	taskCommentList, err := repo.TaskCommentListGet(taskId)
	if err != nil {
		logging.GetLogger().Print(err)
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "Internal Server error", Url: url, StatusCode: http.StatusInternalServerError, Time: time.Now()})
	}
	return c.JSON(http.StatusOK, taskCommentList)
}
