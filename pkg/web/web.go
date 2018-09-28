package web

import (
	"github.com/erei/coemmunity/pkg/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer(static, views string, config *config.Config) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static
	e.Static("/img", static+"/images")
	e.Static("/js", static+"/js")
	e.Static("/stylesheets", static+"/stylesheets")

	// Pages
	e.File("/", views+"/index.html")

	return e
}
