package controller

import (
	"net/http"
	"point-of-sales/models"
	"point-of-sales/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usr usecase.UserUsecaseInterface
}

func NewUserControllerImpl(r *gin.RouterGroup, usr usecase.UserUsecaseInterface) {
	handler := UserController{usr}

	r.POST("/user/create", handler.AddUserController)
	r.GET("/list_users", handler.GetAllUsersController)
	r.GET("/user/:id", handler.GetUserController)
	r.PUT("/user/:id", handler.UpdateUserController)
	r.DELETE("/user/:id", handler.DeleteUserController)
}

func (uc UserController) AddUserController(c *gin.Context) {
	var usr models.Merchant

	if err := c.ShouldBindJSON(&usr); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userData, err := uc.usr.AddUserUsecase(usr)

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

func (uc UserController) GetAllUsersController(c *gin.Context) {
	userData, err := uc.usr.GetAllUsersUsecase()

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

func (uc UserController) GetUserController(c *gin.Context) {
	var usr models.Merchant
	id := c.Param("id")

	userData, err := uc.usr.GetUserUsecase(usr, id)

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

func (uc UserController) UpdateUserController(c *gin.Context) {
	var usr models.Merchant
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	err := c.ShouldBindJSON(&usr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userData, err := uc.usr.UpdateUserUsecase(usr, id)
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

func (uc UserController) DeleteUserController(c *gin.Context) {
	var usr models.Merchant
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	userData, err := uc.usr.DeleteUserUsecase(usr, id)
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
