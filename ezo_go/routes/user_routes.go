package routes

import (
	"github.com/NazishAhsan/easy_busy_book_go/controller"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, uc controller.UserController) {
	// e.POST("/signup", uc.SignUp)
	// e.POST("/login", uc.Login)

	u := e.Group("/api")

	u.POST("/signup", uc.SignUp)
	u.POST("/login", uc.Login)
	u.POST("/logout", uc.Logout)
	u.GET("/csrf-token", uc.CsrfToken)
}