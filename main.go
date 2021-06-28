package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	// import environment vars
	setEnvFromFile(".env")

	mux := http.NewServeMux()
	mux.HandleFunc("/hash", PostHash)
	mux.HandleFunc("/hash/", GetHashById)
	mux.HandleFunc("/stats", GetStats)
	mux.HandleFunc("/shutdown", PostShutdown)

	http.ListenAndServe(":"+getEnv("PORT"), mux)
}

func PostHash(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hash/" && req.Method != "POST" {
		http.NotFound(res, req)
		return
	}

	res.Write([]byte("/hash"))
}

func GetHashById(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	test, _ := regexp.MatchString("/hash/[0-9]/", path)

	if !test && req.Method != "GET" {
		http.NotFound(res, req)
		return
	}

	hashId := strings.Split(path, "/")[2]

	fmt.Println(hashId)
	res.Write([]byte("/hash/" + hashId))
}

func GetStats(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/stats/" && req.Method != "GET" {
		http.NotFound(res, req)
		return
	}

	res.Write([]byte("/stats"))
}

func PostShutdown(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/shutdown/" && req.Method != "POST" {
		http.NotFound(res, req)
		return
	}

	res.Write([]byte("{message: 'shutdown signal sent'}"))
}
