package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Start(port string)
}

type HttpServer struct {
	engine *gin.Engine
}

func NewHttpServer(router *gin.Engine) Server {
	return &HttpServer{engine: router}
}

func (s *HttpServer) Start(port string) {
	fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
	if err := s.engine.Run(":" + port); err != nil {
		log.Fatalf("Error on server start: %v", err)
	}
}
