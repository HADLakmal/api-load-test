package middleware

import (
	"github.com/HADLakmal/api-load-test/internal/util/container"
	"github.com/pickme-go/metrics"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// MetricsMiddleware alrers the request.
type MetricsMiddleware struct {
	httpReqDuration metrics.Observer
}

// NewMetricsMiddleware creates a new instance of MetricMiddleware
func NewMetricsMiddleware(ctr *container.Container) *MetricsMiddleware {

	httpReqDuration := ctr.Adapters.MetricsReporter.Observer(metrics.MetricConf{Path: "http_request_duration_microseconds",
		Labels: []string{
			"code", "method", "url",
		},
	},
	)

	return &MetricsMiddleware{
		httpReqDuration: httpReqDuration,
	}
}

// Middleware executes middleware rules of MetricsMiddleware.
func (rtm *MetricsMiddleware) Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		startTime := time.Now()
		lrw := newLoggingResponseWriter(w)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(lrw, r)

		duration := float64(time.Since(startTime).Nanoseconds() / 1e3)

		rtm.httpReqDuration.Observe(duration, map[string]string{
			"code":   strconv.Itoa(lrw.statusCode),
			"method": r.Method,
			"url":    generalizePath(r.URL.Path),
		})
	})
}

// The loggingResponseWriter is created embedding http.ResponseWriter
// https://golang.org/doc/effective_go.html#embedding
// https://ndersson.me/post/capturing_status_code_in_net_http/
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
func generalizePath(path string) string {

	routeParts := strings.Split(path, "/")

	for i, routePart := range routeParts {

		_, errInt := strconv.ParseInt(routePart, 10, 64)
		if errInt == nil {
			routeParts[i] = "<id>"
			continue
		}

		_, errFloat := strconv.ParseFloat(routePart, 64)
		if errFloat == nil {
			routeParts[i] = "<val>"
			continue
		}

		if strings.Contains(routePart, "{") {
			routeParts[i] = "<path_var>"
			continue
		}
	}

	return strings.Join(routeParts, "/")
}
