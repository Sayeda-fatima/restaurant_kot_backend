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
	// restaurant table
	restaurantTableValidator := validator.NewRestaurantTableValidator()
	restaurantTableRepository := repository.NewRestaurantTableRepository(db)
	restaurantTableUsecase := usecase.NewRestaurantTableUsecase(restaurantTableRepository, restaurantTableValidator)
	restaurantTableController := controller.NewRestaurantTableController(restaurantTableUsecase)

	// current inventory
	currentInventoryRepository := repository.NewCurrentInventoryRepository(db)
	// product
	productValidator := validator.NewProductValidator()
	productRepository := repository.NewProductRepository(db)
	// inventory transaction
	inventoryTransactionRepository := repository.NewInventoryTransactionRepository(db)
	
	productUsecase := usecase.NewProductUsecase(productRepository, productValidator, inventoryTransactionRepository)
	productController := controller.NewProductController(productUsecase)

	inventoryTransactionUsecase := usecase.NewInventoryTransactionUsecase(inventoryTransactionRepository, productRepository, currentInventoryRepository)
	inventoryTransactionController := controller.NewInventoryTransactionController(inventoryTransactionUsecase)

	// recipe
	recipeValidator := validator.NewRecipeValidator()
	recipeRepository := repository.NewRecipeRepository(db)
	recipeUsecase := usecase.NewRecipeUsecase(recipeRepository, recipeValidator)
	recipeController := controller.NewRecipeController(recipeUsecase)

	// recipe product 
	recipeProductValidator := validator.NewRecipeProductValidator()
	recipeProductRepository := repository.NewRecipeProductRepository(db)
	recipeProductUsecase := usecase.NewRecipeProductUsecase(recipeProductRepository, recipeProductValidator)
	recipeProductController := controller.NewRecipeProductController(recipeProductUsecase)

	// menu
	menuValidator := validator.NewMenuValidator()
	menuRepository := repository.NewMenuRepository(db)
	menuUsecase := usecase.NewMenuUsecase(menuRepository, menuValidator)
	menuController := controller.NewMenuController(menuUsecase)

	// menu item
	menuItemValidator := validator.NewMenuItemValidator()
	menuItemRepository := repository.NewMenuItemRepository(db)
	menuItemUsecase := usecase.NewMenuItemUsecase(menuItemRepository, menuItemValidator)
	menuItemController := controller.NewMenuItemController(menuItemUsecase)

	// customer
	customerValidator := validator.NewCustomerValidator()
	customerRepository := repository.NewCustomerRepository(db)
	customerUsecase := usecase.NewCustomerUsecase(customerRepository, customerValidator)
	customerController := controller.NewCustomerController(customerUsecase)

	// staff 
	staffValidator := validator.NewStaffValidator()
	staffRepository := repository.NewStaffRepository(db)
	staffUsecase := usecase.NewStaffUsecase(staffRepository, staffValidator)
	staffController := controller.NewStaffController(staffUsecase)

	// staff weekly schedule
	weeklyStaffScheduleValidator := validator.NewWeeklyStaffScheduleValidator()
	weeklyStaffScheduleRepository := repository.NewWeeklyStaffScheduleRepository(db)
	weeklyStaffScheduleUsecase := usecase.NewWeeklyStaffScheduleUsecase(weeklyStaffScheduleRepository, weeklyStaffScheduleValidator)
	weeklyStaffScheduleController := controller.NewWeeklyStaffSchedule(weeklyStaffScheduleUsecase)

	// allergen
	allergenValidator := validator.NewAllergenValidator()
	allergenRepository := repository.NewAllergenRepository(db)
	allergenUsecase := usecase.NewAllergenUsecase(allergenRepository, allergenValidator)
	allergenController := controller.NewAllergenController(allergenUsecase)

	// menu allergen
	menuAllergenValidator := validator.NewMenuAllergenValidator()
	menuAllergenRepository := repository.NewMenuAllergenRepository(db)
	menuAllergenUsecase := usecase.NewMenuAllergenUsecase(menuAllergenRepository, menuAllergenValidator)
	menuAllergenController := controller.NewMenuAllergenController(menuAllergenUsecase)

	// cart
	cartValidator := validator.NewCartValidator()
	cartRepository := repository.NewCartRepository(db)
	// cart item
	cartItemValidator := validator.NewCartItemValidator()
	cartItemRepository := repository.NewCartItemRepository(db)
	cartItemUsecase := usecase.NewCartItemUsecase(cartItemRepository, cartItemValidator, db)
	cartItemController := controller.NewCartItemController(cartItemUsecase)

	cartUsecase := usecase.NewCartUsecase(cartRepository, cartValidator, db)
	cartController := controller.NewCartController(cartUsecase)

	// order
	orderValidator := validator.NewOrderValidator()
	orderRepository := repository.NewOrderRepository(db)
	orderUsecase := usecase.NewOrderUsecase(orderRepository, orderValidator, db, cartUsecase)
	orderController := controller.NewOrderController(orderUsecase)

	// order item
	orderItemValidator := validator.NewOrderItemValidator()
	orderItemRepository := repository.NewOrderItemRepository(db)
	orderItemUsecase := usecase.NewOrderItemUsecase(orderItemRepository, orderItemValidator)
	orderItemController := controller.NewOrderItemController(orderItemUsecase)


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
	routes.RestaurantTableRoutes(e, restaurantTableController)
	routes.ProductRoutes(e, productController)
	routes.InventoryTransactionRoutes(e, inventoryTransactionController)
	routes.CustomerRoutes(e, customerController)
	routes.StaffRoutes(e, staffController, weeklyStaffScheduleController)
	routes.RecipeRoutes(e, recipeController, recipeProductController)
	routes.AllergenRoutes(e, allergenController)
	routes.MenuRoutes(e, menuController, menuItemController, menuAllergenController)
	routes.CartRoutes(e, cartController, cartItemController)
	routes.OrderRoutes(e, orderController, orderItemController)

	e.Logger.Fatal(e.Start(":8000"))
}