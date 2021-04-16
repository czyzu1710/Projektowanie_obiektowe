package main

import (
	"Zadanie03v2/controllers"
	"Zadanie03v2/database"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	database.Init()
	db := database.GetInstance()
	db.Open()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", controllers.GetStudents)
	e.POST("/users", controllers.CreateStudent)
	e.GET("/users/:id", controllers.GetStudent)
	e.PUT("/users/:id", controllers.UpdateStudent)
	e.DELETE("/users/:id", controllers.DeleteStudent)

	e.Logger.Fatal(e.Start(":1323"))
	db.Close()
}
