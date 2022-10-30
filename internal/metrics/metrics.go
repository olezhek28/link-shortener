package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const appPrefix = "link_shortener_"

type Metrics struct {
	RequestTotal *prometheus.CounterVec
}

var metrics *Metrics

func Init(_ context.Context) error {
	metrics = &Metrics{
		RequestTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: appPrefix + "requests_total",
				Help: "Количество запросов к серверу",
			},
			[]string{"method", "handler", "code"},
		),
	}

	prometheus.MustRegister(metrics.RequestTotal)

	return nil
}

func IncRequestTotal(method, handler, code string) {
	metrics.RequestTotal.WithLabelValues(method, handler, code).Inc()
}
