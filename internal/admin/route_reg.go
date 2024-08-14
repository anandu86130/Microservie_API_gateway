package admin

import (
	"github.com/Microservice_API_gateway/internal/admin/handler"
	adminpb "github.com/Microservice_API_gateway/internal/admin/pb"
	"github.com/gin-gonic/gin"
)

type AdminRoutes struct {
	client adminpb.AdminServiceClient
}

func (a *AdminRoutes) Login(c *gin.Context) {
	handler.AdminLoginHandler(c, a.client)
}

func (a *AdminRoutes) CreateProduct(c *gin.Context) {
	handler.CreateProductHandler(c, a.client)
}

func (a *AdminRoutes) FindProduct(c *gin.Context) {
	handler.FindProductHandler(c, a.client)
}

func (a *AdminRoutes) FindAllProduct(c *gin.Context) {
	handler.FetchAllProductHandler(c, a.client)
}
