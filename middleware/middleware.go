package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"restApi/logger"
	"restApi/metrics"
	"time"
)

// Logging оборачивает запрос, позволяя добавить логирование
func Logging(f http.Handler, logger *logger.Log, metrics *metrics.Metrics) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Запрос: " + r.URL.String())

		start := time.Now()
		f.ServeHTTP(w, r)
		elapsed := time.Since(start)

		metrics.RequestDuration.Observe(float64(elapsed) / float64(time.Second))
		metrics.RequestDurationQuantile.Observe(float64(elapsed) / float64(time.Second))

		metrics.RequestCount.With(prometheus.Labels{"method": r.Method, "url": r.RequestURI}).Add(1)
		metrics.TestCounter.Inc()
	})
}
