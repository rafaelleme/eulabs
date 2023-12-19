package main

import (
	"api/app/controller"
	"api/app/persistence"
	"api/app/repository"
	"api/app/service"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dbConn, err := persistence.InitDB()
	if err != nil {
		panic("Cannot connect to db")
	}

	productRepo := repository.NewProductRepository(dbConn)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	e := echo.New()

	e.Use(middleware.Recover())

	productController.RegisterRoutes(e)

	port := 8080
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
