package main

import (
	"go-api/internal/controller"
	"go-api/internal/db"
	"go-api/internal/repository"
	"go-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductReository(dbConection)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.DELETE("/product/:productId", ProductController.DelProductById)

	server.Run(":8000")
}
