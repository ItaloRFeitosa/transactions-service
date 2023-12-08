package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/italorfeitosa/transactions-service/pkg/requestid"
)

func RequestID(c *gin.Context) {
	requestID := c.Request.Header.Get("X-Request-ID")
	if requestID == "" {
		requestID = uuid.NewString()
	}

	c.Request = c.Request.WithContext(requestid.NewContext(c.Request.Context(), requestID))
	c.Header("X-Request-ID", requestID)
	c.Next()
}
