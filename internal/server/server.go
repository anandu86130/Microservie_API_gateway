package server

import "github.com/gin-gonic/gin"

type ServerStruct struct {
	R *gin.Engine
}

func (s *ServerStruct) StartServer() {
	s.R.Run(":8081")
}

func Server() *ServerStruct {
	engine := gin.Default()
	return &ServerStruct{
		R: engine,
	}
}
