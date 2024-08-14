package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	adminpb "github.com/Microservice_API_gateway/internal/admin/pb"
	"github.com/Microservice_API_gateway/internal/model"
	"github.com/gin-gonic/gin"
)

func AdminLoginHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	var admin model.AdminModel
	if err := c.ShouldBindJSON(&admin); err != nil {
		log.Printf("error when binding JSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx := context.Background()
	response, err := client.AdminLogin(ctx, &adminpb.AdminRequest{
		Username: admin.Username,
		Password: admin.Password,
	})
	if err != nil {
		log.Printf("error when logging in admin %v err: %v", admin.Username, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusAccepted,
		"message": fmt.Sprintf("%v logged in successfully", admin.Username),
		"data": response,
	})
}
