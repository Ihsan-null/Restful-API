package routes

import (
	"part1/controllers"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	e.GET("/books", controllers.GetBooksControllers)
	e.GET("/book/:id", controllers.GetBookController)
	e.DELETE("/book/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	e.POST("/books/:id", controllers.CreateBookController)
}
