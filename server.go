package main

import (
	"github.com/labstack/echo/v4"
	"go-echo/app"
	"go-echo/controller"
	"go-echo/route"
)

func main() {
	e := echo.New()
	DB := app.NewDB()
	userControllerImpl := controller.NewUserController(DB)
	route.ApiRoutes(e, userControllerImpl)
	e.Logger.Fatal(e.Start(":1323"))
}
