package service

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/aj9mb/task-management/model"
	"github.com/aj9mb/task-management/repo"
	"github.com/labstack/echo/v4"
)

func BoardAdd(c echo.Context) error {
	url := c.Request().URL.RequestURI()
	userId, err := strconv.ParseInt(c.Request().Header.Get("x-user-id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid user id", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	board := new(model.Board)
	c.Bind(board)
	if board.Name == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "bad request", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	board.CreatedBy = userId
	board, err = repo.BoardAdd(board)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "error creating board", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	boardUser := &model.BoardUser{BoardId: board.Id, UserId: userId, AddedBy: userId}
	id, err := repo.BoardUserAdd(boardUser)
	if id <= 0 || err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "error adding user to board", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	return c.JSON(http.StatusOK, board)
}

func BoardUserAdd(c echo.Context) error {
	url := c.Request().URL.RequestURI()
	userId, err := strconv.ParseInt(c.Request().Header.Get("x-user-id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid user id", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	boardUser := new(model.BoardUser)
	c.Bind(boardUser)
	log.Print(boardUser)
	if boardUser.BoardId <= 0 || boardUser.UserId <= 0 {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "bad request", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	boardUser.AddedBy = userId
	id, err := repo.BoardUserAdd(boardUser)
	if id <= 0 || err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "error adding user to board", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	return c.JSON(http.StatusOK, boardUser)
}

func BoardListGet(c echo.Context) error {
	url := c.Request().URL.RequestURI()
	userId, err := strconv.ParseInt(c.Request().Header.Get("x-user-id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid user id", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	boardList, err := repo.BoardListGet(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "error", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	return c.JSON(http.StatusOK, boardList)
}
