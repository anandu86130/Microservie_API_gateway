package user

import (
	"github.com/Microservice_API_gateway/internal/user/handler"
	userpb "github.com/Microservice_API_gateway/internal/user/pb"
	"github.com/gin-gonic/gin"
)

type User struct {
	client userpb.UserServicesClient
}

func (u *User) Signup(c *gin.Context) {
	handler.UserSignupHandler(c, u.client)
}

func (u *User) Login(c *gin.Context) {
	handler.UserLoginHandler(c, u.client)
}

func (u *User) VerifyOTP(c *gin.Context) {
	handler.UserVerifyOTPHandler(c, u.client)
}

func (u *User) FindProduct(c *gin.Context) {
	handler.FindProductHandler(c, u.client)
}

func (u *User) FindAllProduct(c *gin.Context){
	handler.FetchAllProductHandler(c, u.client)
}