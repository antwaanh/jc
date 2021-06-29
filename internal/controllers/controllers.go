package controllers

import (
	"encoding/json"
	"fmt"
	"jc/internal/services/dao"
	"jc/internal/services/queue"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Stats struct {
	Total   int `json:"total"`
	Average int `json:"average"`
}

func PostHash(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hash/" && req.Method != "POST" {
		http.NotFound(res, req)
		return
	}

	pw := req.FormValue("password")

	entry := dao.PasswordEntry{
		Id:        queue.Instance.Len() + 1,
		Value:     pw,
		CreatedAt: time.Now().Unix(),
	}

	dao.PersistPassword(entry.Id, entry.Value)
	queue.Instance.PushBack(entry)

	fmt.Println(dao.Instance)

	res.Write([]byte(strconv.Itoa(entry.Id)))
}

func GetHashById(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	test, _ := regexp.MatchString("/hash/[0-9]/", path)

	if !test && req.Method != "GET" {
		http.NotFound(res, req)
		return
	}

	hashId, _ := strconv.Atoi(strings.Split(path, "/")[2])
	hashedPw := dao.Instance[hashId]

	res.Write([]byte(hashedPw))
}

func GetStats(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/stats/" && req.Method != "GET" {
		http.NotFound(res, req)
		return
	}

	total := len(dao.Instance)
	avg := total / 5
	s := Stats{Total: total, Average: avg}
	json, e := json.Marshal(s)

	if e != nil {
		log.Println(e)
	}

	res.Write(json)
}

func PostShutdown(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/shutdown/" && req.Method != "POST" {
		http.NotFound(res, req)
		return
	}

	res.Write([]byte("{message: 'shutdown signal sent'}"))
}
