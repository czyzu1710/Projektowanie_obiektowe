package controllers

import (
	"Zadanie03v2/dao"
	"Zadanie03v2/model"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func GetStudents(c echo.Context) error {
	students, err := dao.GetStudents()
	if err != nil {
		log.Println("Unable to fetch students")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, students)

}

func GetStudent(c echo.Context) error {
	student, err := dao.GetStudent(c.Param("id"))
	if err != nil {
		log.Println("No such student")
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, student)

}

func CreateStudent(c echo.Context) error {
	student := new(model.Student)
	err := c.Bind(student)
	if err != nil {
		log.Println("Unable to create student.")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	dao.Save(student)
	return c.JSON(http.StatusOK, student)
}

func UpdateStudent(c echo.Context) error {
	student := new(model.Student)
	err := c.Bind(student)
	if err != nil {
		log.Println("Unable to update student.")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = dao.Update(student)
	if err != nil {
		log.Println("Error trying update student. No such student.")
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, student)
}

func DeleteStudent(c echo.Context) error {
	err := dao.Delete(c.Param("id"))
	if err != nil {
		log.Println("Unable to delete student")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, "")
}
