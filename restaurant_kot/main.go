package main

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/common"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/database"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/middlewares"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/routes"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/usecase"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	common.Newlogger()
	db := database.NewDB()
	email := common.NewEmailService()
	e := echo.New()

	// user
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository, userValidator, email)
	userController := controller.NewUserController(userUseCase)

	// organization
	organizationValidator := validator.NewOrganizationValidator()
	organizationRepository := repository.NewOrganizationRepository(db)
	organizationUsecase := usecase.NewOrganizationUsecase(organizationRepository, organizationValidator)
	organizationController := controller.NewOrganizationController(organizationUsecase)

	// restaurant
	restaurantValidator := validator.NewRestaurantValidator()
	restaurantRepository := repository.NewRestaurantRepository(db)
	restaurantUsecase := usecase.NewRestaurantUsecase(restaurantRepository, restaurantValidator)
	restaurantController := controller.NewRestaurantController(restaurantUsecase)

	// staff 
	staffValidator := validator.NewStaffValidator()
	staffRepository := repository.NewStaffRepository(db)
	staffUsecase := usecase.NewStaffUsecase(staffRepository, staffValidator)
	staffController := controller.NewStaffController(staffUsecase)


	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000", os.Getenv("APP_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middlewares.LoggingMiddleWare)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// routes
	routes.AuthRoutes(e, userController)
	routes.OrganizationRoutes(e, organizationController)
	routes.RestaurantRoutes(e, restaurantController)
	routes.StaffRoutes(e, staffController)

	e.Logger.Fatal(e.Start(":8000"))
}