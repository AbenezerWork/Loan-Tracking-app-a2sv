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
	loanCollection := client.Database("Loan-Tracker").Collection("Loans")

	userMockCollection := repository.NewMongoCollection(userCollection)
	tokenMockCollection := repository.NewMongoCollection(tokenCollection)
	loanMockCollection := repository.NewMongoCollection(loanCollection)

	userRepo := repository.NewUserRepository(userMockCollection)
	tokenRepo := repository.NewTokenRepository(tokenMockCollection)
	loanRepo := repository.NewLoanRepository(loanMockCollection)

	userUsecase := usecase.NewUserUsecase(jwtService, userRepo, tokenRepo)
	loanUsecase := usecase.NewLoanUseCase(loanRepo)
	tokenUsecase := usecase.NewTokenUsecase(tokenRepo, *jwtService)

	NewSignupRouter(r, userUsecase)
	NewUserRouter(r, userUsecase, *jwtService)
	NewLoanRouter(r, loanUsecase, *jwtService)
	NewRefreshTokenRouter(r, userUsecase, tokenUsecase, *jwtService)

}
