package userhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/tonet-me/tonet-core/delivery/http/middleware"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	r := e.Group("/users")

	r.POST("/login-or-register", h.loginOriRegister)
	r.GET("/profile", h.getUserInfo, middleware.Authentication(h.authSvc, h.authConfig))
}
