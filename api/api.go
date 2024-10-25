package api

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/whoamiApolo/api-students/db"
	"net/http"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentDB := db.NewStudentHandler(database)

	return &API{
		Echo: e,
		DB:   studentDB,
	}
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudents)
	api.Echo.GET("/students/:id", api.getStudent)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

// Handler - gerenciar a rota raiz "/"
func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}

	return c.JSON(http.StatusOK, students)
}

func (api *API) createStudents(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error creating student")
	}

	return c.String(http.StatusOK, "Create a new student")
}

func (api *API) getStudent(c echo.Context) error {
	id := c.Param("id")
	getStud := fmt.Sprintf("Get student with ID: %s", id)
	return c.String(http.StatusOK, getStud)
}

func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	updateStud := fmt.Sprintf("Update student with ID: %s", id)
	return c.String(http.StatusOK, updateStud)
}

func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("Delete student with ID: %s", id)
	return c.String(http.StatusOK, deleteStud)
}
