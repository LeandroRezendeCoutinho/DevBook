package middlewares

import (
	"api/src/auth"
	"fmt"

	"github.com/labstack/echo/v4"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := auth.ValidateToken(c); err != nil {
			return c.JSON(401, map[string]string{"error": fmt.Sprintf("Unauthorized: %s", err.Error())})
		}
		return next(c)
	}
}
