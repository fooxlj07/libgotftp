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
		Subsystem: "tftp",
		Name:      "tftp_request_total",
		Help:      "Total number of tftp request.",
	})
	requestDuration := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "tftp_request_duration",
		Help:    "The time to get a response",
		Buckets: prometheus.LinearBuckets(20, 5, 5),
	})

	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
	return &MetricCollector{
		RequestCount:    requestCount,
		RequestDuration: requestDuration,
	}
}
