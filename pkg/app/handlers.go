package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
