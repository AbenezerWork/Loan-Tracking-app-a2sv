package router

import (
	"loan-tracking/api/controller"
	"loan-tracking/domain"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(r *gin.Engine, userUsecase domain.UserUsecaseInterface) {
	// Initialize controllers
	signupController := controller.NewSignupController(userUsecase)
	// Public routes
	r.POST("users/signup", signupController.Signup)
	r.GET("users/verify/:token", signupController.VerifyEmail)
}
