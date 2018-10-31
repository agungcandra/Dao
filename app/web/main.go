package main

import (
	"log"
	"net/http"

	"github.com/agungcandra/dao"
	"github.com/agungcandra/dao/bootstrap/database"
	"github.com/agungcandra/dao/bootstrap/environment"
	"github.com/agungcandra/dao/handler"
	"github.com/bukalapak/packen/middleware"
	"github.com/julienschmidt/httprouter"
)

func main() {
	environment.Load()

	router := httprouter.New()
	database := database.GetConnection()

	dao := dao.NewDao(database)

	userHandler := handler.NewUserHandler(dao)

	log.Println("Running on " + environment.ENVIRONMENT + " mode")
	router.GET("/users", middleware.HTTP(middleware.ApplyDecorators(userHandler.CreateUser, middleware.WithStandardContext())))
	http.ListenAndServe(":1234", router)
}
