package routes

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo, uc controller.UserController) {

	u := e.Group("/api")

	u.POST("/signup", uc.SignUp)
	u.POST("/login", uc.Login)
	u.POST("/logout", uc.Logout)
	u.GET("/csrf-token", uc.CsrfToken)
	u.POST("/refresh-token", uc.RefreshToken)
	u.POST("/forgot-password", uc.ForgotPassword)
	u.GET("/reset-password/{token}", uc.CsrfToken)
	u.POST("/reset-password", uc.ResetPassword)
}
