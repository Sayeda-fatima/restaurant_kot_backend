package main

import (

	"github.com/NazishAhsan/easy_busy_book_go/database"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/usecase"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
	"github.com/NazishAhsan/easy_busy_book_go/routes"
	"github.com/NazishAhsan/easy_busy_book_go/controller"
)

func main(){

	db := database.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUseCase)
	e:= routes.NewRoute(userController)
	e.Logger.Fatal(e.Start(":8000"))
}