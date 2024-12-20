package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InventoryTransactionRoutes(e *echo.Echo, ic controller.InventoryTransactionController){

	i := e.Group("/api/inventory")
	i.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	i.GET("", ic.GetInventoryTransactionList)
	i.POST("/:product/stock", ic.AddStock)
	i.POST("/:product/stock/adjust", ic.AdjustStock)
	i.POST("/:product/stock/waste", ic.RecordWaste)

	i.POST("/current", ic.CreateCurrentInventoryValue)
	i.GET("/cogs", ic.GetCostOfGoodsSold)
	i.GET("/waste", ic.GetWasteDuringTimePeriod)
	i.GET("/daily-consumption", ic.GetDailyConsumption)
}