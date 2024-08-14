package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	userpb "github.com/Microservice_API_gateway/internal/user/pb"
	"github.com/gin-gonic/gin"
)

func FetchAllProductHandler(c *gin.Context, client userpb.UserServicesClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*30000)
	defer cancel()

	response, err := client.UserProductList(ctx, &userpb.RNoParam{})
	if err != nil{
		log.Printf("error when finding all productlist err:%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":http.StatusBadRequest,
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":http.StatusAccepted,
		"message":"fetched all products successfully",
		"data": response,
	})
}

func FindProductHandler(c *gin.Context,client userpb.UserServicesClient){
	ctx, cancel := context.WithTimeout(c,time.Second*3000)
	defer cancel()

	id := c.Query("id")
	name := c.Query("name")

	response := &userpb.ProductDetails{}

	var err error
	if id == "" && name == ""{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":http.StatusBadRequest,
			"error":"invalid query",
		})
		return
	}else if id != ""{
		id, err := strconv.Atoi(id)
		if err != nil{
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":http.StatusBadRequest,
				"error": "invalid id",
			})
			return
		}
		response, err = client.UserProductByID(ctx, &userpb.ProductID{Id: uint32(id)})
		if err != nil{
			log.Printf("error when finding product by id err: %v",err)
			c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
				"status":http.StatusBadRequest,
				"error":err.Error(),
			})
			return
		}
	}else if name != ""{
		response, err = client.UserProductByName(ctx, &userpb.ProductByName{Name: name})
		if err != nil{
			log.Printf("error when finding product by name err: %v",err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":http.StatusBadRequest,
				"error":err.Error(),	
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusAccepted,
		"message": "fetched all products successfully",
		"data": response,
	})
}
