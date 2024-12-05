package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func MenuAllergenRoutes(e *echo.Echo, ma controller.MenuAllergenController){

	m := e.Group("/api/menu")
	m.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	m.GET("/:menuID/allergen", ma.GetMenuAllergenList)
	m.POST("/:menuID/allergen", ma.CreateMenuAllergen)
	m.PUT("/:menuID/allergen/:id", ma.UpdateMenuAllergen)
	m.DELETE("/:menuID/allergen/:id", ma.DeleteMenuAllergen)
}