package main

import (
	"context"
	"log"
	"net/http"

	"github.com/olezhek28/link-shortener/internal/app"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ctx := context.Background()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		// nolint
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Fatal(err)
		}
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
