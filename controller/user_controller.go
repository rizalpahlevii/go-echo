package controller

import (
	"github.com/labstack/echo/v4"
)

type UserController interface {
	Create(e echo.Context) error
	Update(e echo.Context) error
	Delete(e echo.Context) error
	FindById(e echo.Context) error
	FindAll(e echo.Context) error
}
