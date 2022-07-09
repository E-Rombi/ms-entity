package main

import (
	"ms-entity/internal/app/repository"
	productService "ms-entity/internal/app/service/product"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	port = ":8080"
)

func main() {
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	productRepository := repository.NewProductRepository()

	productService.NewRegisterProductService(productRepository).RegisterServer(server)
	productService.NewFindProductByIdService(productRepository).RegisterServer(server)

	server.Logger.Fatal(server.Start(port))
}
