package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Microservice_API_gateway/internal/model"
	userpb "github.com/Microservice_API_gateway/internal/user/pb"
	"github.com/gin-gonic/gin"
)

func UserSignupHandler(c *gin.Context, client userpb.UserServicesClient){
	ctx, cancel := context.WithTimeout(c, time.Second*3000)
	defer cancel()

	var user model.UserModel
	if err := c.ShouldBindJSON(&user); err != nil{
		log.Printf("error when binding json:%v",err)
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"status":http.StatusBadRequest,
			"error":err.Error(),
		})
		return
	}

	response, err := client.Signup(ctx, &userpb.SignupRequest{
		Email: user.Email,
	})

	if err!=nil{
		log.Printf("error when signing up user %v err: %v", user.Email,err)
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"status":http.StatusBadRequest,
			"error":err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusAccepted,
		"message": fmt.Sprintf("%v otp generated successfully", user.Email),
		"data": response,
	})
}

func UserVerifyOTPHandler(c *gin.Context, client userpb.UserServicesClient){
	ctx, cancel := context.WithTimeout(c,time.Second*3000)
	defer cancel()
	
	var request model.VerifyOTP
	if err := c.ShouldBindJSON(&request); err != nil{
		log.Printf("error when binding json :%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"status":http.StatusBadRequest,
			"error":err.Error(),
		})
		return
	}

	response, err := client.VerifyOTP(ctx, &userpb.VerifyOTPRequest{
		Email: request.Email,
		Otp: request.Otp,
	})

	if err != nil{
		log.Printf("error when verifying OTP for user %v err: %v", request.Email, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":http.StatusAccepted,
		"message":"OTP verified successfully",
		"token": response.Token,
	})
}

func UserLoginHandler(c *gin.Context, client userpb.UserServicesClient){
	ctx, cancel := context.WithTimeout(c,time.Second*3000)
	defer cancel()

	var user model.UserModel
	if err := c.ShouldBindJSON(&user); err != nil{
		log.Printf("error when binding json: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"status":http.StatusBadRequest, 
			"error": err.Error(),
		})
		return
	}

	response, err:= client.Login(ctx, &userpb.LoginRequest{
		Email: user.Email,
	})

	if err != nil{
		log.Printf("error when signing up user %v err: %v", user.Email,err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":http.StatusBadRequest,
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":http.StatusAccepted,
		"message":fmt.Sprintf("%v login successfully", user.Email),
		"data":response,
		"token":response.Token,
	})
}

