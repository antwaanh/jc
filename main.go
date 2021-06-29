package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type PasswordRow struct {
	id        int
	value     string
	createdAt int64
}

type Stats struct {
	total   int
	average int
}

var queue = list.New()
var repo = map[int]string{} // lock and write

// watch queue
func queueWatcher(q *list.List) {
	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Println("Watching...", e)
	}
}

func main() {
	// import environment vars
	setEnvFromFile(".env")

	go queueWatcher(queue)

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

	pw := req.FormValue("password") // queue to storage

	// TODO: move to service
	pr := PasswordRow{
		id: queue.Len() + 1, value: pw, createdAt: time.Now().Unix(),
	}

	queue.PushBack(pr)
	repo[pr.id] = "..." // store empty string until hash is complete

	fmt.Println(repo)

	res.Write([]byte(strconv.Itoa(pr.id)))
}

func GetHashById(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	test, _ := regexp.MatchString("/hash/[0-9]/", path)

	if !test && req.Method != "GET" {
		http.NotFound(res, req)
		return
	}

	hashId, _ := strconv.Atoi(strings.Split(path, "/")[2])

	res.Write([]byte(repo[hashId]))
}

func GetStats(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/stats/" && req.Method != "GET" {
		http.NotFound(res, req)
		return
	}

	json, e := json.Marshal(Stats{total: 20, average: 2})

	if e != nil {
		log.Fatal(e)
	}

	res.Write([]byte(json))
}

func PostShutdown(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/shutdown/" && req.Method != "POST" {
		http.NotFound(res, req)
		return
	}

	res.Write([]byte("{message: 'shutdown signal sent'}"))
}
