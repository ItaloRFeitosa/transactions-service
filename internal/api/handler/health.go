package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/transactions-service/api/model"
)

type Health struct{}

func (h *Health) Register(r gin.IRouter) {
	r.GET("/health", h.Handle)
}

// swagger:route GET /health health getHealth
//
// Health Endpoint
//
//	Produces:
//	- application/json
//	Responses:
//	  200: HealthResponse
//	  500: ErrorResponse
func (Health) Handle(c *gin.Context) {
	c.JSON(http.StatusOK, model.HealthResponse{Data: model.Health{Success: true}})
}
