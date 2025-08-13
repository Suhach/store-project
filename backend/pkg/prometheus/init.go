package prometheus

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var requestCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"method", "path", "status"},
)

func InitMetrics() {
	prometheus.MustRegister(requestCounter)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		requestCounter.WithLabelValues(c.Request.Method, c.FullPath(), fmt.Sprintf("%d", c.Writer.Status())).Inc()
	}
}
