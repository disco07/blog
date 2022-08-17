package controllers

import (
	"flag"
	"fmt"
	"github.com/disco07/blog/config"
	"github.com/disco07/blog/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

type Server struct {
	config config.Config
	repo   repository.Repository
}

func Run() {
	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 4001, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "dev", "Application environment (dev|prod)")
	flag.StringVar(&cfg.DB, "dsn", "postgres://blog:blog@localhost/blog?sslmode=disable", "Postgres connection string")
	flag.Parse()

	db, err := open(cfg)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	server := Server{
		config: cfg,
		repo:   repository.NewRepository(db),
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

func open(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DB)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
