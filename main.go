package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/urfave/cli"
	"github.com/urfave/cli/v2"
)

var (
	port    string
	debug   bool
	timeout time.Duration
)

func main() {
	app := &cli.App{
		Name: "barcode",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "port",
				Usage:       "Port for the server to listen on",
				Aliases:     []string{"p"},
				EnvVars:     []string{"PORT"},
				Destination: &port,
				Value:       "3030",
			},
			&cli.BoolFlag{
				Name:        "debug",
				Usage:       "Enable Debug Logs",
				Aliases:     []string{"verbose"},
				EnvVars:     []string{"DEBUG"},
				Destination: &debug,
				Value:       false,
			},
			&cli.DurationFlag{
				Name:        "timeout",
				Usage:       "Request Timeouts",
				Aliases:     []string{"request-timeout"},
				EnvVars:     []string{"REQUEST_TIMEOUT"},
				Destination: &timeout,
				Value:       60 * time.Second,
			},
		},
		Action: func(c *cli.Context) error {
			if debug {
				log.Base().SetLevel("debug")
			}
			r := chi.NewRouter()

			// A good base middleware stack
			r.Use(middleware.RequestID)
			r.Use(middleware.RealIP)
			r.Use(middleware.Logger)
			r.Use(middleware.Recoverer)

			// Set a timeout value on the request context (ctx), that will signal
			// through ctx.Done() that the request has timed out and further
			// processing should be stopped.
			r.Use(middleware.Timeout(60 * time.Second))

			r.Route("/", func(r chi.Router) {
				r.Handle("/metrics", promhttp.Handler())
				r.Mount("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				})

				// r.Mount("/barcode", controller.V1Barcodes{}.Routes())
				// r.Mount("/product", controller.V1Products{}.Routes())
				// r.Mount("/healthcheck", controller.HealthCheck{}.Routes())
				r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
					http.FileServer(http.Dir("./public")).ServeHTTP(w, r)
				})
			})

			log.Debugf("Launching server listening on: %s\n", port)
			http.ListenAndServe(fmt.Sprintf(":%s", port), r)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
