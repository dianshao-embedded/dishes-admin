package middleware

import (
	"dishes-admin-mod/internal/app/contextx"
	"dishes-admin-mod/pkg/logger"
	"dishes-admin-mod/pkg/util/trace"

	"github.com/gin-gonic/gin"
)

// Get or set trace_id in request context
func TraceMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		traceID := c.GetHeader("X-Request-Id")
		if traceID == "" {
			traceID = trace.NewTraceID()
		}

		ctx := contextx.NewTraceID(c.Request.Context(), traceID)
		ctx = logger.NewTraceIDContext(ctx, traceID)
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set("X-Trace-Id", traceID)

		c.Next()
	}
}
