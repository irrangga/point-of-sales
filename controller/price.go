package controller

import (
	"net/http"
	"point-of-sales/middleware"
	"point-of-sales/models"
	"point-of-sales/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PriceController struct {
	prc usecase.PriceUsecaseInterface
}

func NewPriceControllerImpl(r *gin.RouterGroup, prc usecase.PriceUsecaseInterface) {
	handler := PriceController{prc}

	auth := r.Group("")
	auth.Use(middleware.AuthMiddleware().MiddlewareFunc())
	{
		auth.POST("/price/insert", handler.AddPriceController)
		auth.GET("/list_prices", handler.GetAllPricesController)
		auth.GET("/price/:id", handler.GetPriceController)
		auth.PUT("/price/:id", handler.UpdatePriceController)
		auth.DELETE("/price/:id", handler.DeletePriceController)
	}
}

func (pc PriceController) AddPriceController(c *gin.Context) {
	var prc models.Price

	if err := c.ShouldBindJSON(&prc); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	priceData, err := pc.prc.AddPriceUsecase(prc)

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
		Data:    priceData,
	}
	c.JSON(http.StatusOK, response)
}

func (pc PriceController) GetAllPricesController(c *gin.Context) {
	id := c.Query("product_id")

	priceData, err := pc.prc.GetAllPricesUsecase(id)
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
		Data:    priceData,
	}

	c.JSON(http.StatusOK, response)
}

func (pc PriceController) GetPriceController(c *gin.Context) {
	var prc models.Price
	id := c.Param("id")

	priceData, err := pc.prc.GetPriceUsecase(prc, id)

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
		Data:    priceData,
	}
	c.JSON(http.StatusOK, response)
}

func (pc PriceController) UpdatePriceController(c *gin.Context) {
	var prc models.Price
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	err := c.ShouldBindJSON(&prc)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	priceData, err := pc.prc.UpdatePriceUsecase(prc, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	priceData.ID = idInt

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    priceData,
	}
	c.JSON(http.StatusOK, response)
}

func (pc PriceController) DeletePriceController(c *gin.Context) {
	var prc models.Price
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	priceData, err := pc.prc.DeletePriceUsecase(prc, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	priceData.ID = idInt

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    priceData,
	}
	c.JSON(http.StatusOK, response)
}
