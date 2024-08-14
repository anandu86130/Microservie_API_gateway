package admin

import (
	"log"

	"github.com/Microservice_API_gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewAdminRoute(r *gin.Engine) {
	client, err := ClientDial()
	if err != nil {
		log.Fatalf("error Not connected with gRPC server: %v", err.Error())
	}

	adminHandler := &AdminRoutes{
		client: client,
	}

	apiAdmin := r.Group("/api/admin")

	{
		apiAdmin.POST("/login", adminHandler.Login)
	}

	apiAuthAdmin := r.Group("/api/admin/auth")
	apiAuthAdmin.Use(middleware.Authorization("admin"))

	{
		apiAuthAdmin.POST("/product", adminHandler.CreateProduct)
		apiAuthAdmin.GET("/products", adminHandler.FindAllProduct)
		apiAuthAdmin.GET("/product", adminHandler.FindProduct)
	}
}
