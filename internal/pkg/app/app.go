package app

import (
	"fmt"
	"github.com/Kirnata/middleware/internal/app/endpoint"
	"github.com/Kirnata/middleware/internal/app/mw"
	"github.com/Kirnata/middleware/internal/app/service"
	"github.com/labstack/echo"
	"log"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Server running ")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal()
	}
	return nil
}
