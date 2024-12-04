package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func AllergenRoutes(e *echo.Echo, ac controller.AllergenController){

	a := e.Group("/api/allergen")
	a.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	a.GET("", ac.GetAllergenList)
	a.POST("", ac.CreateAllergen)
	a.PUT("/:id", ac.UpdateAllergen)
	a.DELETE("/:id", ac.DeleteAllergen)
}