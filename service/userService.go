package service

import (
	"net/http"
	"time"

	"github.com/aj9mb/task-management/helpers"
	"github.com/aj9mb/task-management/logging"
	"github.com/aj9mb/task-management/model"
	"github.com/aj9mb/task-management/repo"
	"github.com/labstack/echo/v4"
)

func AddUser(c echo.Context) error {
	url := c.Request().URL.RequestURI()
	user := new(model.User)
	c.Bind(user)
	if user.Name == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "bad request", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	username, err := helpers.GenerateUserName(user.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "unexpected error", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	user.UserName = username
	pwd, err := helpers.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "unexpected error", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	user.Password = pwd
	user, err = repo.UserAdd(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "error creating user", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	return c.JSON(http.StatusCreated, user)
}

func LoginUser(c echo.Context) error {
	url := c.Request().URL.RequestURI()
	user := new(model.User)
	c.Bind(user)
	if user.UserName == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "bad request", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	user1, err := repo.GetUserByUserName(user.UserName)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "incorrect username/password", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
	}
	if helpers.ComparePassword(user1.Password, user.Password) {
		user1.Password = ""
		user1 = userAuthAdd(user1)
		return c.JSON(http.StatusOK, user1)
	}
	return c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "incorrect username/password", Url: url, StatusCode: http.StatusBadRequest, Time: time.Now()})
}

func userAuthAdd(user *model.User) *model.User {
	if user.Id < 1 {
		logging.GetLogger().Print("invalid user id")
		return user
	}
	username := helpers.GenerateRandomString(30)
	pwd := helpers.GetUuid()
	err := repo.UserAuthAdd(user.Id, username, pwd)
	if err != nil {
		logging.GetLogger().Print(err)
		return user
	}
	user.BasicAuthUserName = &username
	user.BasicAuthPwd = &pwd
	return user
}

func UserAuthCheck(username, pwd string) (bool, error) {
	if username == "" || pwd == "" {
		return false, nil
	}
	return repo.UserAuthCheck(username, pwd)
}
