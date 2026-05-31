package routes

import (
	"os"

	"github.com/Sayeda-fatima/restaurant_kot_backend/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RecipeRoutes(e *echo.Echo, rc controller.RecipeController, rp controller.RecipeProductController){

	r := e.Group("/api/recipe")
	r.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	r.GET("", rc.GetRecipeList)
	r.POST("", rc.CreateRecipe)
	r.GET("/:id", rc.GetRecipe)
	r.GET("/:id/cost", rc.GetRecipeCost)
	r.PUT("/:id", rc.UpdateRecipe)
	r.DELETE("/:id", rc.DeleteRecipe)
	
	// recipe product routes
	r.GET("/:recipe/product", rp.GetProductInRecipe)
	r.POST("/:recipe/product", rp.CreateRecipeProduct)
	r.PUT("/:recipe/product/:id", rp.UpdateRecipeProduct)
	r.DELETE("/:recipe/product/:id", rp.DeleteRecipeProduct)
}