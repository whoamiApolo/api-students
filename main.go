package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/whoamiApolo/api-students/db"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/students", getStudents)
	e.POST("/students", createStudents)
	e.GET("/students/:id", getStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}

// Handler - gerenciar a rota raiz "/"
func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "List of all students")
}

func createStudents(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}

	if err := db.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error creating student")
	}

	return c.String(http.StatusOK, "Create a new student")
}

func getStudent(c echo.Context) error {
	id := c.Param("id")
	getStud := fmt.Sprintf("Get student with ID: %s", id)
	return c.String(http.StatusOK, getStud)
}

func updateStudent(c echo.Context) error {
	id := c.Param("id")
	updateStud := fmt.Sprintf("Update student with ID: %s", id)
	return c.String(http.StatusOK, updateStud)
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("Delete student with ID: %s", id)
	return c.String(http.StatusOK, deleteStud)
}
