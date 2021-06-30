package server

import (
	"jc/src/services/config"
	"log"
	"net/http"
	"time"
)

var Instance http.Server

func StartUp() {
	config.SetEnvFromFile(".env")
	host := config.GetEnv("HOST")
	port := config.GetEnv("PORT")
	Instance = http.Server{Addr: ":"+port}
	log.Printf("Starting server at %v:%v", host, port)
}

func Shutdown() {
	time.Sleep(2 * time.Second)
	e := Instance.Shutdown(nil)
	log.Println(e)
}