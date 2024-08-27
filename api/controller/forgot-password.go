package controller

import (
	"fmt"
	"loan-tracking/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordController struct {
	userUsecase domain.UserUsecaseInterface
}

func NewForgotPasswordController(uu domain.UserUsecaseInterface) *ForgotPasswordController {
	return &ForgotPasswordController{
		userUsecase: uu,
	}
}

func (fp *ForgotPasswordController) ForgotPassword(c *gin.Context) {
	email := domain.VerificationClaims{}
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("email: ", email)

	// panic(email)
	err := fp.userUsecase.ForgotPassword(&email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "err": "controller"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent to email"})
}

func (fp *ForgotPasswordController) VerifyForgotToken(c *gin.Context) {
	token := c.Query("token")
	claims, err := fp.userUsecase.VerifyForgotPassword(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fp.userUsecase.UpdateUser(claims.Name, claims.Password)

	c.JSON(http.StatusOK, gin.H{})
}
