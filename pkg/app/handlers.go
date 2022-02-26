package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"weight-tracker/pkg/api"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		response := map[string]string{
			"status": "success",
			"data":   "Weight tracker API running smoothly",
		}
		c.JSON(http.StatusOK, response)
	}
}

func BuildResponse(status, data string) map[string]string {
	response := map[string]string{
		"status": status,
		"data":   data,
	}
	return response
}

func (s *Server) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newUser api.NewUserRequest

		err := c.ShouldBindJSON(&newUser)

		var response map[string]string

		if err != nil {
			log.Printf("handler error: %v", err)
			response = BuildResponse("badRequest", err.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		}

		err = s.userService.New(newUser)

		if err != nil {
			log.Printf("service error: %v", err)
			response = BuildResponse("internalServerError", err.Error())

			c.JSON(http.StatusInternalServerError, err)
			return
		}
		response = BuildResponse("success", "new user created")

		c.JSON(http.StatusOK, response)
	}
}
