package main

import (
	"github.com/aj9mb/task-management/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/board", service.BoardAdd)
	e.POST("/board/people", service.BoardUserAdd)
	e.GET("/board/list/get", service.BoardListGet)
	e.Logger.Fatal(e.Start(":8080"))
}
