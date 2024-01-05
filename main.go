package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	e.GET("/board/list/get", nil)
	e.Logger.Fatal(e.Start(":8080"))
}
