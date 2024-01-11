package main

import (
	"github.com/Agungsusilo2/Hackfest/model/app"
	"github.com/Agungsusilo2/Hackfest/model/controller"
	"github.com/Agungsusilo2/Hackfest/model/repository"
	"github.com/Agungsusilo2/Hackfest/model/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db := app.GetConnection()
	validate := validator.New()
	applicationRepository := repository.NewCategoryRepository()
	applicationService := service.NewApplicationService(applicationRepository, db, validate)
	categoryController := controller.NewCategoryController(applicationService)

	router := httprouter.New()

	router.GET("/api/applicants", categoryController.FindAll)
	router.POST("/api/applicants", categoryController.Create)
	router.GET("/api/applicants/:applicantId", categoryController.FindById)
	router.PUT("/api/applicants/:applicantId", categoryController.Update)
	router.DELETE("/api/applicants/:applicantId", categoryController.Delete)

	server := http.Server{Addr: "localhost:3000",
		Handler: router}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
