package main

import (
	"jc/src/controllers"
	"jc/src/services/config"
	"log"
	"net/http"
)

func main() {
	config.SetEnvFromFile(".env")

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
