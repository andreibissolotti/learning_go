package controller

import (
	"net/http"
	"strconv"

	"github.com/andreibissolotti/learning_go/internal/model"
	"github.com/andreibissolotti/learning_go/internal/usecase"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{Message: "Id do produto não pode ser nulo"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{Message: "Id do produto deve ser um numero inteiro"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	foundedProduct, err := p.productUseCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if foundedProduct == nil {
		response := model.Response{Message: "Produto não encontrado"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, foundedProduct)
}

func (p *productController) DelProductById(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{Message: "Id do produto não pode ser nulo"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{Message: "Id do produto deve ser um numero inteiro"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.productUseCase.DelProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, "")
}