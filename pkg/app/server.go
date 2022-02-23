package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"weight-tracker/pkg/api"
)

type Server struct {
	router        *gin.Engine
	userService   api.UserService
	weightService api.WeightService
}

func NewServer(router *gin.Engine, userService api.UserService, weightService api.WeightService) *Server {
	return &Server{
		router:        router,
		userService:   userService,
		weightService: weightService,
	}
}

func (s *Server) Run() error {
	r := s.Routes()

	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}
	return nil
}
