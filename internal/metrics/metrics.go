package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const appPrefix = "link_shortener_"

type Metrics struct {
	requestTotal *prometheus.CounterVec
	cacheHit     *prometheus.CounterVec
	cacheMiss    *prometheus.CounterVec
}

var metrics *Metrics

func Init(_ context.Context) error {
	metrics = &Metrics{
		requestTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: appPrefix + "requests_total",
				Help: "Количество запросов к серверу",
			},
			[]string{"method", "handler", "code"},
		),
		cacheHit: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: appPrefix + "cache_hit_total",
				Help: "Количество запросов к кешу",
			},
			[]string{"handler"},
		),
		cacheMiss: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: appPrefix + "cache_miss_total",
				Help: "Количество запросов к кешу",
			},
			[]string{"handler"},
		),
	}

	prometheus.MustRegister(metrics.requestTotal)

	return nil
}

func IncRequestTotal(method, handler, code string) {
	metrics.requestTotal.WithLabelValues(method, handler, code).Inc()
}

func IncCacheHit(handler string) {
	metrics.cacheHit.WithLabelValues(handler).Inc()
}

func IncCacheMiss(handler string) {
	metrics.cacheMiss.WithLabelValues(handler).Inc()
}
