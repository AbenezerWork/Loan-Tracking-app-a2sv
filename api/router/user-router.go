package router

// import (
// 	"loan-tracking/api/controller"
// 	"loan-tracking/usecase"
// 	"loan-tracking/utils"
//
// 	"github.com/gin-gonic/gin"
// )
//
// func NewUserRouter(r *gin.Engine, userUsecase *usecase.UserUsecase, jwtService utils.JWTService) {
// 	userController := controller.NewUserController(userUsecase)
// 	r.POST("/login", userController.Login)
// 	forgotPasswordController := controller.NewForgotPasswordController(userUsecase)
// 	r.POST("/forgotpassword", forgotPasswordController.ForgotPassword)
// 	r.POST("/verfiyforgotpassword", forgotPasswordController.VerifyForgotOTP)
// 	auth := r.Group("/api")
// 	auth.Use(utils.AdminMiddleware(jwtService))
// 	{
// 		// Admin-specific routes
// 		auth.GET("/getallusers", userController.GetAllUsers)
// 		auth.DELETE("/deleteuser/:id", userController.DeleteUser)
// 		// Admin-specific routes
// 		auth.POST("/getallusers", userController.GetAllUsers)
// 		auth.PUT("/deleteusers/:id", userController.DeleteUser)
// 	}
// }
