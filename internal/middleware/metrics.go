package middleware

import (
	"net/http"
	"strconv"

	"github.com/olezhek28/link-shortener/internal/metrics"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func AddMetrics(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &StatusRecorder{ResponseWriter: w, Status: http.StatusOK}

		h.ServeHTTP(recorder, r)

		metrics.IncRequestTotal(r.Method, r.RequestURI, strconv.FormatInt(int64(recorder.Status), 10))
	})
}
