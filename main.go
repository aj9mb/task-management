package main

import (
	"crypto/subtle"
	"strings"

	"github.com/aj9mb/task-management/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{Validator: func(username, pwd string, ctx echo.Context) (bool, error) {
		return (subtle.ConstantTimeCompare([]byte(username), []byte("user")) == 1 && subtle.ConstantTimeCompare([]byte(pwd), []byte("pwd")) == 1), nil
	}, Skipper: func(c echo.Context) bool {
		url := c.Request().URL.RequestURI()
		return strings.Contains(url, "/login") || strings.Contains(url, "/signup")
	}}))

	e.POST("/board", service.BoardAdd)
	e.POST("/board/people", service.BoardUserAdd)
	e.GET("/board/list/get", service.BoardListGet)
	e.POST("/signup", service.AddUser)
	e.POST("/login", service.LoginUser)
	e.POST("/:board_id/task", service.TaskAdd)
	e.GET("/:board_id/task/list/get", service.TaskListGet)

	e.Logger.Fatal(e.Start(":8080"))
}
