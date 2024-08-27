package router

import (
	"loan-tracking/api/controller"
	"loan-tracking/domain"
	"loan-tracking/utils"

	"github.com/gin-gonic/gin"
)

func NewLoanRouter(router *gin.Engine, loanUsecase domain.LoanUseCaseInterface, jwtService utils.JWTService) {
	loanController := controller.NewLoanController(loanUsecase)
	loanRoutes := router.Group("/loans")
	loanRoutes.Use(utils.AuthMiddleware(jwtService))
	{
		loanRoutes.POST("/", loanController.CreateLoan)
		loanRoutes.GET("/:id", loanController.GetLoan)

	}
	admin := router.Group("/admin")
	admin.Use(utils.AdminMiddleware(jwtService))
	{
		admin.PUT("/:id/decision", loanController.UpdateLoanStatus)
		admin.DELETE("/:id", loanController.DeleteLoan)
		admin.GET("/", loanController.ListLoans)

	}
}
