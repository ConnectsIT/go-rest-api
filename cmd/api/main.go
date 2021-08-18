package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type AppStatus struct {
	Status       string `json:"status"`
	Environnment string `json:"environment"`
	Version      string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
}

type config struct {
	port int
	env  string
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "server port to listen on")
	flag.StringVar(&cfg.env, "env", "developement", "Applicatioon environment(develooppement/production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("starting server on port ", cfg.port)

	err := srv.ListenAndServe()

	if err != nil {
		log.Println(err)
	}
}
