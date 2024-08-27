package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"loan-tracking/repository"
	"loan-tracking/usecase"
	"loan-tracking/utils"
)

func InitRoutes(r *gin.Engine, client *mongo.Client) {

	r.MaxMultipartMemory = 8 << 20

	jwtService := utils.NewJWTService(os.Getenv("JWT_SECRET"), "Kal", os.Getenv("JWT_REFRESH_SECRET"))

	userCollection := client.Database("Loan-Tracker").Collection("Users")
	tokenCollection := client.Database("Loan-Tracker").Collection("Tokens")

	userMockCollection := repository.NewMongoCollection(userCollection)
	tokenMockCollection := repository.NewMongoCollection(tokenCollection)

	userRepo := repository.NewUserRepository(userMockCollection)
	tokenRepo := repository.NewTokenRepository(tokenMockCollection)

	userUsecase := usecase.NewUserUsecase(jwtService, userRepo, tokenRepo)

	NewSignupRouter(r, userUsecase)

}
