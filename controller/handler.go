package controller

import (
	"github.com/labstack/echo/v4"
	"go-api/domain"
	"log"
	"net/http"
)

func save(c echo.Context) error {
	var p domain.Person
	if err := c.Bind(&p); err != nil {
		return err
	}

	err := controller.Save(p)
	if err != nil {
		log.Printf("Error happened commons.Save. Err: %s", err)
	}

	return c.JSON(http.StatusCreated, p)
}

func get(c echo.Context) error {
	resp := controller.Get()
	return c.JSON(http.StatusOK, resp)
}
