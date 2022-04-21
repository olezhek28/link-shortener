package main

import (
	"context"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"github.com/olezhek28/link-shortener/internal/app/pkg/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Can't create app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Can't run app: %s", err.Error())
	}
}
