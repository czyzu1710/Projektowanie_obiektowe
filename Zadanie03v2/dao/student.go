package dao

import (
	"Zadanie03v2/database"
	"Zadanie03v2/model"
	"errors"
	"fmt"
	"log"
)

func GetStudent(id string) (*model.Student, error) {
	db := database.GetInstance()

	row, err := db.Query(fmt.Sprintf("SELECT * FROM STUDENTS WHERE ID = %v LIMIT 1", id))
	if err != nil {
		log.Println("Error trying fetch student")
		return nil, err
	}

	if row.Next() {
		student := new(model.Student)
		err := row.Scan(&student.Id, &student.Name, &student.Surname, &student.Birthday)
		if err != nil {
			return nil, err
		}
		return student, nil
	}

	return nil, errors.New("No such student.")
}

func GetStudents() ([]*model.Student, error) {
	db := database.GetInstance()
	row, err := db.Query("SELECT * FROM STUDENTS s")
	if err != nil {
		log.Println("Error trying fetch students")
	}

	var students = make([]*model.Student, 0)
	for row.Next() {
		student := new(model.Student)
		err := row.Scan(&student.Id, &student.Name, &student.Surname, &student.Birthday)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if len(students) == 0 {
		return students, errors.New("Error trying fetch students or no students are available")
	}

	return students, nil
}

func Save(s *model.Student) error {
	db := database.GetInstance()
	id, err := db.Execute(fmt.Sprintf("INSERT INTO STUDENTS(NAME, SURNAME, BIRTHDAY) VALUES ('%v', '%v', '%v')", s.Name, s.Surname, s.Birthday))
	if err != nil {
		log.Println("Error trying insert student.")
		return err
	}
	s.Id = id
	return nil
}

func Update(s *model.Student) error {
	db := database.GetInstance()
	_, err := db.Execute(fmt.Sprintf("UPDATE STUDENTS SET NAME = '%v', SURNAME = '%v', BIRTHDAY = '%v' WHERE ID = %v", s.Name, s.Name, s.Birthday, s.Id))
	if err != nil {
		log.Println("Error trying update student.")
		return err
	}
	return nil
}

func Delete(id string) error {
	db := database.GetInstance()
	_, err := db.Execute(fmt.Sprintf("DELETE FROM STUDENTS WHERE ID = %v", id))
	if err != nil {
		log.Println("Error trying to delete student")
		return err
	}
	return nil
}
