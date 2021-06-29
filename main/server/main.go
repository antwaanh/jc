package main

import (
	"container/list"
	"jc/src/controllers"
	"jc/src/services/config"
	"jc/src/services/dao"
	"jc/src/services/queue"
	"log"
	"net/http"
)

func main() {
	// import environment vars
	config.SetEnvFromFile(".env")

	queue.Instance = list.New()
	queue.Manager()

	dao.Instance = map[int]string{}

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
