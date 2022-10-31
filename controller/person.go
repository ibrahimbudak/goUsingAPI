package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-api/domain"
	"go-api/service"
	"net/http"
)

var controller personController

type PersonController interface {
	Get() []domain.Person
	Save(m domain.Person) error
}

type personController struct {
	s service.PersonService
}

func NewPersonController(s service.PersonService) PersonController {
	controller = personController{s: s}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{"*"},
	}))

	e.GET("/list", get)
	e.POST("/kayit", save)

	e.Logger.Fatal(e.Start(":8080"))

	return controller
}

func (c personController) Get() []domain.Person {
	return c.s.Get()
}

func (c personController) Save(p domain.Person) error {
	return c.s.Save(p)
}
