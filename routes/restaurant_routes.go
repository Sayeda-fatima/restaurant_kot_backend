package routes

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/controller"
	"github.com/labstack/echo/v4"
)

func RestaurantRoutes(e *echo.Echo, rc controller.RestaurantController){

	re := e.Group("/api/restaurant")

	re.GET("/:id", rc.GetRestaurantList)
	re.POST("", rc.CreateRestaurant)
	re.PUT("/:id", rc.UpdateRestaurant)
	re.DELETE("/:id", rc.DeleteRestaurant)
}