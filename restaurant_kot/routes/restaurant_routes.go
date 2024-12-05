package routes

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	"github.com/labstack/echo/v4"
)

func RestaurantRoutes(e *echo.Echo, rc controller.RestaurantController){

	re := e.Group("/api/restaurant")

	re.GET("/:id", rc.GetRestaurantList)
	re.POST("", rc.CreateRestaurant)
	re.PUT("/:id", rc.UpdateRestaurant)
	re.DELETE("/:id", rc.DeleteRestaurant)
}