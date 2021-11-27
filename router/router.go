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
	outletRepo := repository.NewOutletRepoImpl(DB)
	productRepo := repository.NewProductRepoImpl(DB)

	// USE CASE DEPENDENCY
	userUsecase := usecase.NewUserUsecaseImpl(userRepo)
	outletUsecase := usecase.NewOutletUsecaseImpl(outletRepo)
	productUsecase := usecase.NewProductUsecaseImpl(productRepo)

	// CONTROLLER DEPENDENCY
	controller.NewUserControllerImpl(newRoute, userUsecase)
	controller.NewOutletControllerImpl(newRoute, outletUsecase)
	controller.NewProductControllerImpl(newRoute, productUsecase)

	return r
}
