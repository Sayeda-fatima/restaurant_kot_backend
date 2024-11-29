package routes

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	"github.com/labstack/echo/v4"
)

func OrganizationRoutes(e *echo.Echo, oc controller.OrganizationController){

	o := e.Group("/api/organization")

	o.GET("", oc.GetOrganizationList)
	o.POST("", oc.CreateOrganization)
	o.PUT("/:id", oc.UpdateOrganization)
	o.DELETE("/:id", oc.DeleteOrganization)
}