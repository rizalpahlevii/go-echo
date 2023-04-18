package route

import (
	"github.com/labstack/echo/v4"
	"go-echo/controller"
)

func ApiRoutes(
	e *echo.Echo,
	userController controller.UserController,
) {
	g := e.Group("/users")
	g.GET("", userController.FindAll)
	g.GET("/:userId", userController.FindById)
	g.POST("", userController.Create)
	g.PUT("/:userId", userController.Update)
	g.DELETE("/:userId", userController.Delete)
}
