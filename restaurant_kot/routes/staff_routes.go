package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func StaffRoutes(e *echo.Echo, sc controller.StaffController){

	s := e.Group("/api/staff")
	s.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	s.GET("", sc.GetStaffListByOrganization)
	s.GET("/restaurant", sc.GetStaffListByRestaurant)
	s.POST("", sc.CreateStaff)
	s.PUT("/:id", sc.UpdateStaff)
	s.DELETE("/:id", sc.DeleteStaff)
}