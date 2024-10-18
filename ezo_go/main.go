package main

import (
	"os"
	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/controller"
	"github.com/NazishAhsan/easy_busy_book_go/database"
	"github.com/NazishAhsan/easy_busy_book_go/middlewares"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/routes"
	"github.com/NazishAhsan/easy_busy_book_go/usecase"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){

	common.Newlogger()
	db := database.NewDB()
	e := echo.New()

	// organization
	organizationValidator := validator.NewOrganizationValidator()
	organizationRepository := repository.NewOrganizationRepository(db)
	organizationUseCase := usecase.NewOrganizationUsecase(organizationRepository, organizationValidator)
	organizationController := controller.NewOrganizationController(organizationUseCase)

	// user
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUseCase)

	// customer
	customerValidator := validator.NewCustomerValidator()
	customerRepository := repository.NewCustomerRepository(db)
	customerUseCase := usecase.NewCustomerUsecase(customerRepository, customerValidator)
	customerController := controller.NewCustomerController(customerUseCase)
	
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", os.Getenv("APP_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middlewares.LoggingMiddleWare)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//routes 
	routes.UserRoutes(e, userController)
	routes.CustomerRoutes(e, customerController)
	routes.OrganizationRoutes(e, organizationController)
	e.Logger.Fatal(e.Start(":8000"))
}