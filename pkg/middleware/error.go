package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/transactions-service/pkg/errs"
)

func Error(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	originalError := c.Errors[0].Err
	httpError := errs.ToHttpError(originalError)
	if httpError.StatusCode == http.StatusInternalServerError {
		slog.Error(originalError.Error(), "error", originalError)
	}
	c.AbortWithStatusJSON(httpError.StatusCode, httpError)
}
