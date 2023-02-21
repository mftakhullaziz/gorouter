package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"gohttp/apirest"
	"gohttp/config"
	"gohttp/repository"
	"gohttp/service"
	"gohttp/utils/exception"
	"gohttp/utils/helper"
	"net/http"
)

func main() {
	db := config.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryService := service.NewCategoryServiceImpl(categoryRepository, db, validate)
	categoryRest := apirest.NewCategoryRestImpl(categoryService)

	router := httprouter.New()

	router.GET("/internal/api/categories", categoryRest.FindAll)
	router.POST("/internal/api/categories", categoryRest.Create)
	router.GET("/internal/api/categories/:catId", categoryRest.FindById)
	router.PUT("/internal/api/categories/:catId", categoryRest.Update)
	router.DELETE("/internal/api/categories/:catId", categoryRest.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
