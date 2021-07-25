package routes

import (
	"part1/controllers"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	e.GET("/users", controllers.GetUsersControllers)
	e.GET("/user/:id", controllers.GetUserController)
	e.DELETE("/user/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.POST("/users/:id", controllers.CreateUserController)
}
