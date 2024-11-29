package routes

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	"github.com/labstack/echo/v4"
)

func RestaurantRoutes(e *echo.Echo, rc controller.RestaurantController){

	r := e.Group("/api/restaurant")

	r.GET("/:id", rc.GetRestaurantList)
	r.POST("", rc.CreateRestaurant)
	r.PUT("/:id", rc.UpdateRestaurant)
	r.DELETE("/:id", rc.DeleteRestaurant)
}