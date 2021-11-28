package controller

import (
	"net/http"
	"point-of-sales/middleware"
	"point-of-sales/models"
	"point-of-sales/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OutletController struct {
	out usecase.OutletUsecaseInterface
}

func NewOutletControllerImpl(r *gin.RouterGroup, out usecase.OutletUsecaseInterface) {
	handler := OutletController{out}

	auth := r.Group("")
	auth.Use(middleware.AuthMiddleware().MiddlewareFunc())
	{
		auth.POST("/outlet/create", handler.AddOutletController)
		auth.GET("/list_outlets", handler.GetAllOutletsController)
		auth.GET("/outlet/:id", handler.GetOutletController)
		auth.PUT("/outlet/:id", handler.UpdateOutletController)
		auth.DELETE("/outlet/:id", handler.DeleteOutletController)
	}
}

func (oc OutletController) AddOutletController(c *gin.Context) {
	var out models.Outlet

	if err := c.ShouldBindJSON(&out); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	outletData, err := oc.out.AddOutletUsecase(out)

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
		Data:    outletData,
	}
	c.JSON(http.StatusOK, response)
}

func (oc OutletController) GetAllOutletsController(c *gin.Context) {
	id := c.Query("merchant_id")

	outletData, err := oc.out.GetAllOutletsUsecase(id)
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
		Data:    outletData,
	}

	c.JSON(http.StatusOK, response)
}

func (oc OutletController) GetOutletController(c *gin.Context) {
	var out models.Outlet
	id := c.Param("id")

	outletData, err := oc.out.GetOutletUsecase(out, id)

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
		Data:    outletData,
	}
	c.JSON(http.StatusOK, response)
}

func (oc OutletController) UpdateOutletController(c *gin.Context) {
	var out models.Outlet
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	err := c.ShouldBindJSON(&out)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	outletData, err := oc.out.UpdateOutletUsecase(out, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	outletData.ID = idInt

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    outletData,
	}
	c.JSON(http.StatusOK, response)
}

func (oc OutletController) DeleteOutletController(c *gin.Context) {
	var out models.Outlet
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	outletData, err := oc.out.DeleteOutletUsecase(out, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	outletData.ID = idInt

	response := models.ResponseCustom{
		Status:  http.StatusOK,
		Message: "success",
		Data:    outletData,
	}
	c.JSON(http.StatusOK, response)
}
