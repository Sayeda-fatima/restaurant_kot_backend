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

	// product
	productValidator := validator.NewProductValidator()
	productRepository := repository.NewProductRepository(db)
	productUseCase := usecase.NewProductUsecase(productRepository, productValidator)
	productController := controller.NewProductController(productUseCase)

	// product category
	productCategoryValidator := validator.NewProductCategoryValidator()
	productCategoryRepository := repository.NewProductCategoryRepository(db)
	productCategoryUsecase := usecase.NewProductCategoryUsecase(productCategoryRepository, productCategoryValidator)
	productCategoryController := controller.NewProductCategoryController(productCategoryUsecase)

	// product stock
	productStockValidator := validator.NewProductStockValidator()
	productStockRepository := repository.NewProductStockRepository(db)
	productStockUseCase := usecase.NewProductStockUsecase(productStockRepository, productStockValidator)
	productStockController := controller.NewProductStockController(productStockUseCase)

	// product image
	productImageValidator := validator.NewProductImageValidator()
	productImageUpload := common.NewImageUpload()
	productImageRepository := repository.NewProductImageRepository(db)
	productImageUseCase := usecase.NewProductImageUsecase(productImageRepository, productImageValidator, productImageUpload)
	productImageController := controller.NewProductImageController(productImageUseCase)

	// supplier
	supplierValidator := validator.NewSupplierValidator()
	supplierRepository := repository.NewSupplierRepository(db)
	supplierUseCase := usecase.NewSupplierUsecase(supplierRepository, supplierValidator)
	supplierController := controller.NewSupplierController(supplierUseCase)

	// expense
	expenseValidator := validator.NewExpenseValidator()
	expenseRepository := repository.NewExpenseRepository(db)
	expenseUseCase := usecase.NewExpenseUsecase(expenseRepository, expenseValidator)
	expenseController := controller.NewExpenseController(expenseUseCase)

	// cart
	cartValidator := validator.NewCartValidator()
	cartRepository := repository.NewCartRepository(db)
	cartUseCase := usecase.NewCartUsecase(cartRepository, cartValidator)
	cartController := controller.NewCartController(cartUseCase)
	// cart item
	cartItemValidator := validator.NewCartItemValidator()
	cartItemRepository := repository.NewCartItemRepository(db)
	cartItemUseCase := usecase.NewCartItemUsecase(cartItemRepository, cartItemValidator)
	cartItemController := controller.NewCartItemController(cartItemUseCase)

	// order
	orderValidator := validator.NewOrderValidator()
	orderRepository := repository.NewOrderRepository(db)
	orderUsecase := usecase.NewOrderUsecase(orderRepository, orderValidator)
	orderController := controller.NewOrderController(orderUsecase)
	// order item
	orderItemValidator := validator.NewOrderItemValidator()
	orderItemRepository := repository.NewOrderItemRepository(db)
	orderItemUsecase := usecase.NewOrderItemUsecase(orderItemRepository, orderItemValidator)
	orderItemController := controller.NewOrderItemController(orderItemUsecase)
	
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
	routes.ProductRoutes(e, productController, productStockController, productImageController)
	routes.ProductCategoryRoutes(e, productCategoryController)
	routes.SupplierRoutes(e, supplierController)
	routes.ExpenseRoutes(e, expenseController)
	routes.CartRoutes(e, cartController, cartItemController)
	routes.OrderRoutes(e, orderController, orderItemController)
	e.Logger.Fatal(e.Start(":8000"))
}