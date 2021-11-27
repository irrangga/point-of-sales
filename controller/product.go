package controller

import (
	"net/http"
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

	r.POST("/product/insert", handler.AddProductController)
	r.GET("/list_products", handler.GetAllProductsController)
	r.GET("/product/:id", handler.GetProductController)
	r.PUT("/product/:id", handler.UpdateProductController)
	r.DELETE("/product/:id", handler.DeleteProductController)
}

func (uc ProductController) AddProductController(c *gin.Context) {
	var prod models.Product

	if err := c.ShouldBindJSON(&prod); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userData, err := uc.prod.AddProductUsecase(prod)

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
		Data:    userData,
	}
	c.JSON(http.StatusOK, response)
}

func (uc ProductController) GetAllProductsController(c *gin.Context) {
	userData, err := uc.prod.GetAllProductsUsecase()

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
		Data:    userData,
	}

	c.JSON(http.StatusOK, response)
}

func (uc ProductController) GetProductController(c *gin.Context) {
	var prod models.Product
	id := c.Param("id")

	userData, err := uc.prod.GetProductUsecase(prod, id)

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
		Data:    userData,
	}
	c.JSON(http.StatusOK, response)
}

func (uc ProductController) UpdateProductController(c *gin.Context) {
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

	userData, err := uc.prod.UpdateProductUsecase(prod, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userData.ID = idInt

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    userData,
	}
	c.JSON(http.StatusOK, response)
}

func (uc ProductController) DeleteProductController(c *gin.Context) {
	var prod models.Product
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	userData, err := uc.prod.DeleteProductUsecase(prod, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userData.ID = idInt

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    userData,
	}
	c.JSON(http.StatusOK, response)
}