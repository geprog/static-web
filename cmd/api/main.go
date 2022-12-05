package main

import (
	"github.com/geprog/static-web/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/deploy", api.Deploy)
	e.GET("/api/list", api.List)
	e.POST("/api/teardown", api.Teardown)

	e.Logger.Fatal(e.Start(":1323"))
}
