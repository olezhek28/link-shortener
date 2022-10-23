package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/olezhek28/link-shortener/internal/app/pkg/app"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {
	ctx := context.Background()

	recordMetrics()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Can't create app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Can't run app: %s", err.Error())
	}
}
