package mw

import (
	"github.com/labstack/echo"
	"log"
	"strings"
)

const roleAdmin string = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		val := ctx.Request().Header.Get("User-Role")
		if strings.EqualFold(val, roleAdmin) {
			log.Println("Hey, this is Santa! 🎅")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
