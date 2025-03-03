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
	address := fmt.Sprintf(":%s", port)
	fmt.Printf("🚀 Server running at http://localhost%s\n", address)

	if err := s.engine.Run(address); err != nil {
		log.Fatalf("❌ Error starting server: %v", err)
	} else {
		log.Println("✅ Server started successfully!")
	}
}
