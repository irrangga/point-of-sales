package router

import (
	"point-of-sales/config/db"
	"point-of-sales/controller"
	"point-of-sales/repository"
	"point-of-sales/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	r := gin.Default()

	DB := db.DB

	newRoute := r.Group("")

	// REPOSITORY DEPENDENCY
	userRepo := repository.NewUserRepoImpl(DB)

	// USE CASE DEPENDENCY
	userUseCase := usecase.NewUserUsecaseImpl(userRepo)

	// CONTROLLER DEPENDENCY
	controller.NewUserControllerImpl(newRoute, userUseCase)

	return r
}
