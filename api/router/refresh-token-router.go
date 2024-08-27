package router

import (
	"loan-tracking/api/controller"
	"loan-tracking/domain"
	"loan-tracking/utils"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(r *gin.Engine, userUsecase domain.UserUsecaseInterface, refreshTokenUsecase domain.RefreshTokenUsecaseInterface, jwtService utils.JWTService) {
	refreshTokenController := controller.NewRefreshTokenController(userUsecase, refreshTokenUsecase, jwtService)
	r.POST("/refreshtoken", refreshTokenController.RefreshToken)
}
