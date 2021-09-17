package tftp

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MetricCollector struct {
	RequestCount    prometheus.Counter
	RequestDuration prometheus.Histogram
}

func NewMetricCollector() *MetricCollector {
	requestCount := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tftp_request_total",
		Help: "Total number of tftp request.",
	})
	requestDuration := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "tftp_request_duration",
		Help:    "The time to get a response",
		Buckets: prometheus.LinearBuckets(0.01, 0.01, 10),
	})

	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
	return &MetricCollector{
		RequestCount:    requestCount,
		RequestDuration: requestDuration,
	}
}
