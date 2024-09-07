package main

import (
	"go-mini-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
