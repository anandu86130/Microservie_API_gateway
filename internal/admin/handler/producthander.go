package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	adminpb "github.com/Microservice_API_gateway/internal/admin/pb"
	"github.com/Microservice_API_gateway/internal/model"
	"github.com/gin-gonic/gin"
)

func CreateProductHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*3000)
	defer cancel()

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Printf("error when binding json: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	response, err := client.CreateProduct(ctx, &adminpb.AProductDetails{
		Category:    product.Category,
		Name:        product.Name,
		Price:       product.Price,
		Imagepath:   product.Imagepath,
		Description: product.Description,
		Size:        product.Size,
		Quantity:    product.Quantity,
	})
	if err != nil {
		log.Fatalf("error when creating productdetails: %v err: %v", product.Name, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v created successfully", product.Name),
		"data":    response,
	})
}

func FetchAllProductHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*30000)
	defer cancel()

	response, err := client.FetchProducts(ctx, &adminpb.AdminNoParam{})
	if err != nil {
		log.Printf("error finding all productlist err: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "fetched all products successfully",
		"data":    response,
	})
}

func FindProductHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*3000)
	defer cancel()

	id := c.Query("id")
	name := c.Query("name")

	response := &adminpb.AProductDetails{}

	var err error
	if id == "" && name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "invalid query",
		})
		return
	} else if id != "" {
		id, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "invalid id",
			})
			return
		}
		response, err = client.FetchByProductID(ctx, &adminpb.AProductByID{Id: uint32(id)})
		if err != nil {
			log.Printf("error when finding product by id err: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		}
	} else if name != "" {
		response, err = client.FetchByName(ctx, &adminpb.AProductByName{Name: name})
		if err != nil {
			log.Printf("error when finding product by name err:%v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "fetched all Productdetails successfully",
		"data":    response,
	})
}
