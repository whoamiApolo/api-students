package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	/*
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
	*/

	// Routes
	e.GET("/students", getStudents)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler - gerenciar a rota raiz "/"
func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "List of all students")
}
