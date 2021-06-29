package main

import (
	"container/list"
	"jc/internal/controllers"
	"jc/internal/envars"
	"jc/internal/services/dao"
	"jc/internal/services/queue"
	"log"
	"net/http"
)

func main() {
	// import environment vars
	envars.SetFromFile(".env")

	queue.Instance = list.New()
	queue.Manager()

	dao.Instance = map[int]string{}

	mux := http.NewServeMux()
	mux.HandleFunc("/hash", controllers.PostHash)
	mux.HandleFunc("/hash/", controllers.GetHashById)
	mux.HandleFunc("/stats", controllers.GetStats)
	mux.HandleFunc("/shutdown", controllers.PostShutdown)

	e := http.ListenAndServe(":"+envars.GetKey("PORT"), mux)
	if e != nil {
		log.Fatal(e)
	}
}
