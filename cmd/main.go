package main

import (
	"flag"
	"log"
	"os"
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

}
