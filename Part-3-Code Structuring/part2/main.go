package main

import (
	"fmt"
	"part1/config"
	"part1/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitDB()
	config.InitPort()
	routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
}
