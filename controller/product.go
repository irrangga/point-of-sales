package controller

import (
	"net/http"
	"point-of-sales/middleware"
	"point-of-sales/models"
	"point-of-sales/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	prod usecase.ProductUsecaseInterface
}

func NewProductControllerImpl(r *gin.RouterGroup, prod usecase.ProductUsecaseInterface) {
	handler := ProductController{prod}

	auth := r.Group("")
	auth.Use(middleware.AuthMiddleware().MiddlewareFunc())
	{
		auth.POST("/product/insert", handler.AddProductController)
		auth.GET("/list_products", handler.GetAllProductsController)
		auth.GET("/product/:id", handler.GetProductController)
		auth.PUT("/product/:id", handler.UpdateProductController)
		auth.DELETE("/product/:id", handler.DeleteProductController)
	}
}

func (pc ProductController) AddProductController(c *gin.Context) {
	var prod models.Product

	if err := c.ShouldBindJSON(&prod); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	productData, err := pc.prod.AddProductUsecase(prod)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    productData,
	}
	c.JSON(http.StatusOK, response)
}

func (pc ProductController) GetAllProductsController(c *gin.Context) {
	id := c.Query("merchant_id")

	productData, err := pc.prod.GetAllProductsUsecase(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    productData,
	}

	c.JSON(http.StatusOK, response)
}

func (pc ProductController) GetProductController(c *gin.Context) {
	var prod models.Product
	id := c.Param("id")

	productData, err := pc.prod.GetProductUsecase(prod, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    productData,
	}
	c.JSON(http.StatusOK, response)
}

func (pc ProductController) UpdateProductController(c *gin.Context) {
	var prod models.Product
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	err := c.ShouldBindJSON(&prod)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	productData, err := pc.prod.UpdateProductUsecase(prod, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	productData.ID = idInt

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    productData,
	}
	c.JSON(http.StatusOK, response)
}

func (pc ProductController) DeleteProductController(c *gin.Context) {
	var prod models.Product
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	productData, err := pc.prod.DeleteProductUsecase(prod, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	productData.ID = idInt

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    productData,
	}
	c.JSON(http.StatusOK, response)
}
