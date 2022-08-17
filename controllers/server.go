package controllers

import (
	"flag"
	"fmt"
	"github.com/disco07/blog/config"
	"log"
	"net/http"
	"time"
)

type Server struct {
	config config.Config
}

func Run() {
	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 4001, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "dev", "Application environment (dev|prod)")
	flag.StringVar(&cfg.DB, "dsn", "postgres://moviego:moviego@localhost/moviego?sslmode=disable", "Postgres connection string")
	flag.Parse()

	server := Server{
		config: cfg,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      server.initializeRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
		return
	}

}
