package endpoint

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type Service interface {
	DayLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}
func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.s.DayLeft()

	s := fmt.Sprintf("%d minutes until New Year!", d)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}
