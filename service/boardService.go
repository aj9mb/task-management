package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aj9mb/task-management/model"
	"github.com/aj9mb/task-management/repo"
	"github.com/labstack/echo/v4"
)

func BoardAdd(c echo.Context) error {
	userId, err := strconv.ParseInt(c.Request().Header.Get("x-user-id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid user id")
	}
	board := new(model.Board)
	c.Bind(board)
	if board.Name == "" {
		return c.String(http.StatusBadRequest, "bad request")
	}
	board.CreatedBy = userId
	board, err = repo.BoardAdd(board)
	if err != nil {
		return c.String(http.StatusBadRequest, "error creating board")
	}
	boardUser := &model.BoardUser{BoardId: board.Id, UserId: userId, AddedBy: userId}
	id, err := repo.BoardUserAdd(boardUser)
	if id <= 0 || err != nil {
		return c.String(http.StatusBadRequest, "error adding user to board")
	}
	return c.JSON(http.StatusOK, board)
}

func BoardUserAdd(c echo.Context) error {
	userId, err := strconv.ParseInt(c.Request().Header.Get("x-user-id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid user id")
	}
	boardUser := new(model.BoardUser)
	c.Bind(boardUser)
	log.Print(boardUser)
	if boardUser.BoardId <= 0 || boardUser.UserId <= 0 {
		return c.String(http.StatusBadRequest, "bad request")
	}
	boardUser.AddedBy = userId
	id, err := repo.BoardUserAdd(boardUser)
	if id <= 0 || err != nil {
		return c.String(http.StatusBadRequest, "error adding user to board")
	}
	return c.String(http.StatusOK, "User added to board")
}

func BoardListGet(c echo.Context) error {
	userId, err := strconv.ParseInt(c.Request().Header.Get("x-user-id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid user id")
	}
	boardList, err := repo.BoardListGet(userId)
	if err != nil {
		return c.String(http.StatusBadRequest, "error")
	}
	return c.JSON(http.StatusOK, boardList)
}
