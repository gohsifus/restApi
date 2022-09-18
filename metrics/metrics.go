package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	RequestCount            *prometheus.CounterVec
	TestCounter             prometheus.Counter
	RequestDuration         prometheus.Histogram
	RequestDurationQuantile prometheus.Summary
}

// NewMetrics Вернет экземпляр структуры с метриками
func NewMetrics() *Metrics {
	reqCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "restapi_requests_total",
			Help: "Количество запросов к сервису",
		},
		[]string{"method", "url"},
	)

	testCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "i_am_test_counter",
		Help: "Тестовая метрика",
	})

	reqDur := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "restapi_request_duration",
		Help: "Длительность запросов",
	})

	reqDurQ := prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "restapi_request_duration_quantile",
		Help: "Длительность запросов по квантилям",
	})

	prometheus.MustRegister(reqCounter, testCounter, reqDur, reqDurQ)

	return &Metrics{
		RequestCount:            reqCounter,
		TestCounter:             testCounter,
		RequestDuration:         reqDur,
		RequestDurationQuantile: reqDurQ,
	}
}
