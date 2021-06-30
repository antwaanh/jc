package main

import (
	"jc/src/controllers"
	"jc/src/services/config"
	"jc/src/services/server-statistics"
	"log"
	"net/http"
)

func main() {
	config.SetEnvFromFile(".env")
	stats.Total = 0
	stats.Avg = 0

	mux := http.NewServeMux()
	mux.HandleFunc("/hash", controllers.PostHash)
	mux.HandleFunc("/hash/", controllers.GetHashById)
	mux.HandleFunc("/stats", controllers.GetStats)
	mux.HandleFunc("/shutdown", controllers.PostShutdown)

	e := http.ListenAndServe(":"+config.GetEnv("PORT"), mux)
	if e != nil {
		log.Fatal(e)
	}
}
