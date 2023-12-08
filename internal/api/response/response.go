package response

import "github.com/gin-gonic/gin"

func Data(v any) gin.H {
	return gin.H{"data": v}
}
