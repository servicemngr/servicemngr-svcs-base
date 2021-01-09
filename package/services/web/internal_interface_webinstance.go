package web

import (
	"github.com/labstack/echo/v4"
)

type WebInstance interface {
	WebRegisterRoutes(group *echo.Group)
}
