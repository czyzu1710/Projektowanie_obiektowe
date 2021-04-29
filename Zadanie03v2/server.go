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
	e.GET("/students", controllers.GetStudents)
	e.POST("/students", controllers.CreateStudent)
	e.GET("/students/:id", controllers.GetStudent)
	e.PUT("/students/:id", controllers.UpdateStudent)
	e.DELETE("/students/:id", controllers.DeleteStudent)

	e.Logger.Fatal(e.Start(":1323"))
	db.Close()

