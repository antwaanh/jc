package main

import (
	"jc/src/controllers"
	"jc/src/services/server"
	stats "jc/src/services/server-statistics"
	"log"
	"net/http"
)

func main() {
	stats.Total = 0
	stats.Avg = 0
	stats.ShutdownSig = false

	server.StartUp()
	http.HandleFunc("/hash", controllers.PostHash)
	http.HandleFunc("/hash/", controllers.GetHashById)
	http.HandleFunc("/stats", controllers.GetStats)
	http.HandleFunc("/shutdown", controllers.PostShutdown)

	e := server.Instance.ListenAndServe()
	if e != nil {
		log.Fatal(e)
	}

}
