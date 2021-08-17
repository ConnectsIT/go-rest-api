package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

type AppStatus struct {
	Status       string
	Environnment string
	Version      string
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

	fmt.Println("running")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := AppStatus{
			Status:       "Available",
			Environnment: cfg.env,
			Version:      "1.0.0",
		}

		js, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Println(err)
	}
}
