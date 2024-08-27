package controller

import (
	"loan-tracking/domain"
	"loan-tracking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	userUsecase domain.UserUsecaseInterface
}

func NewSignupController(usecase domain.UserUsecaseInterface) *SignupController {
	return &SignupController{usecase}
}

func (sc *SignupController) Signup(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if email already exists
	existingUser, _ := sc.userUsecase.GetUserByEmail(&user.Email)
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Check if username already exists
	existingUser, _ = sc.userUsecase.GetUserByUsername(user.Name)
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err, "error": "Failed to hash password"})
		return

	}

	// Generate and send OTP
	veruser := &domain.VerificationClaims{
		Name:     user.Name,
		Email:    user.Email,
		Role:     "user",
		Password: hashed,
	}
	err = sc.userUsecase.GenerateVerificationToken(veruser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err, "error": "Failed to send OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}
func (sc *SignupController) VerifyEmail(c *gin.Context) {
	token := c.Param("token")
	err := sc.userUsecase.VerifyUser(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusCreated)
}
