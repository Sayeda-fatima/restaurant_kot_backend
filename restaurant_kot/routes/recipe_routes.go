package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RecipeRoutes(e *echo.Echo, rc controller.RecipeController){

	r := e.Group("/api/recipe")
	r.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	r.GET("", rc.GetRecipeList)
	r.POST("", rc.CreateRecipe)
	r.PUT("/:id", rc.UpdateRecipe)
	r.DELETE("/:id", rc.DeleteRecipe)
	
}