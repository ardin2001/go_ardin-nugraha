package main

import (
	"net/http"
	"orm_mvc/route"
	"orm_mvc/utils"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	// create a new echo instance

	utils.UserMigrate()

	e := route.Routers()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hallo user!",
		})
	})
	e.Logger.Fatal(e.Start(":8000"))
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
