package main

import (
	"fmt"
	"os"

	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/yfuruyama/crzerolog"
	"go.opencensus.io/plugin/ochttp"

	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Ctx(r.Context()).Info().Msg("Task Completed")
		fmt.Fprint(w, "Hello")
	})

	rootLogger := zerolog.New(os.Stdout)
	middleware := crzerolog.InjectLogger(&rootLogger)
	handler := middleware(mux)

	httpHandler := &ochttp.Handler{
		Propagation: &propagation.HTTPFormat{},
		Handler:     handler,
	}
	if err := http.ListenAndServe(":8080", httpHandler); err != nil {
		log.Fatal().Msg("Can't start service")
	}
}
