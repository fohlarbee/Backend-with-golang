package main

import (
	"log"
	"net/http"
	"time"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}
func (app * application) mount() *chi.Mux{
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/v1", func(r chi.Router){
		r.Get("/health", app.healthCheckHandler)
	})

	// http.ListenAndServe(":3000", r)

	return r;
}

func (app *application) run(mux *chi.Mux) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		 ReadTimeout: time.Second * 10,
		 IdleTimeout: time.Minute,
	}

	log.Printf("Starting server on %s", app.config.addr)
	return srv.ListenAndServe()
}   