package di

import (
	"github.com/Microservice_API_gateway/config"
	"github.com/Microservice_API_gateway/internal/admin"
	"github.com/Microservice_API_gateway/internal/server"
	"github.com/Microservice_API_gateway/internal/user"
)

func Init() {
	server := server.Server()
	config.LoadConfig()
	user.NewUserRoutes(server.R)
	admin.NewAdminRoute(server.R)
	server.StartServer()
}
