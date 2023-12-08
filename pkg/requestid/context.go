package requestid

import (
	"context"

	"github.com/google/uuid"
)

type key string

var contextKey key = "request_id"

func NewContext(ctx context.Context, args ...string) context.Context {
	var requestID string
	if len(args) == 0 {
		requestID = uuid.NewString()
	} else {
		requestID = args[0]
	}

	return context.WithValue(ctx, contextKey, requestID)
}

func FromContext(ctx context.Context) string {
	var requestID string
	v := ctx.Value(contextKey)

	if v == nil {
		return requestID
	}

	requestID, ok := v.(string)
	if !ok {
		return requestID
	}

	return requestID
}
