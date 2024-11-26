package routes

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo, uc controller.UserController) {
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	e.POST("/logout", uc.Logout)
	e.GET("/csrf-token", uc.CsrfToken)
}
