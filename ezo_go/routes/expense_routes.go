package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_go/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func ExpenseRoutes(e *echo.Echo, ec controller.ExpenseController){

	c := e.Group("/api/expense")
	c.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		TokenLookup: "header:x-auth",
	}))

	c.GET("", ec.GetExpenseList)
	c.POST("", ec.CreateExpense)
	c.PUT("/:id", ec.UpdateExpense)
	c.PUT("/:id/soft-delete", ec.DeleteExpense)
}