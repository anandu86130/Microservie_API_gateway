package user

import (
	"log"

	"github.com/Microservice_API_gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewUserRoutes(r *gin.Engine) {
	client, err := ClientDial()
	if err != nil {
		log.Fatalf("error not connected to gRPC server, %v", err.Error())
	}

	userHandler := &User{
		client: client,
	}

	apiUser := r.Group("api/user")
	{
		apiUser.POST("/signup", userHandler.Signup)
		apiUser.POST("/login", userHandler.Login)
		apiUser.POST("/verifyotp", userHandler.VerifyOTP)
	}

	apiUserAuth := r.Group("/api/user/auth")
	apiUserAuth.Use(middleware.Authorization("user"))
	{
		apiUserAuth.GET("/getproducts", userHandler.FindAllProduct)
		apiUserAuth.GET("/getproduct", userHandler.FindProduct)
	}
}
