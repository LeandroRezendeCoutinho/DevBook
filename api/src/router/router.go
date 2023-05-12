package router

import "github.com/labstack/echo/v4"

func Generate() *echo.Echo {
	return echo.New()
}
