package controller

import (
	"api/app/entity"
	"api/app/service"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{service}
}

func (c *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/products", c.GetProducts)
	e.GET("/products/:id", c.GetProductByID)
	e.POST("/products", c.CreateProduct)
	e.PUT("/products/:id", c.UpdateProduct)
	e.DELETE("/products/:id", c.DeleteProduct)
}

func (c *ProductController) GetProducts(ctx echo.Context) error {
	filter := ctx.QueryParam("filter")

	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil {
		limit = 10
	}

	products, err := c.service.GetProducts(filter, page, limit)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) GetProductByID(ctx echo.Context) error {
	id := ctx.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product id"})
	}

	product, err := c.service.GetProductByID(productID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Product not found"})
	}

	return ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) CreateProduct(ctx echo.Context) error {
	var product entity.Product

	if err := ctx.Bind(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Err on decode json"})
	}

	if err := c.service.CreateProduct(product.Name, product.Price); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}

func (c *ProductController) UpdateProduct(ctx echo.Context) error {
	id := ctx.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product id"})
	}

	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			map[string]string{"error": "Cannot get body"})
	}

	var updatedProduct entity.Product
	if err := json.Unmarshal(body, &updatedProduct); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Err on decode json"})
	}

	if err := c.service.UpdateProduct(productID, updatedProduct); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *ProductController) DeleteProduct(ctx echo.Context) error {
	id := ctx.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product id"})
	}

	if err := c.service.DeleteProduct(productID); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
