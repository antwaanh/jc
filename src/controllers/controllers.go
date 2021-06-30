package controllers

import (
	"encoding/json"
	"jc/src/services/dao"
	"jc/src/services/server-statistics"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
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
		Id:    len(dao.Instance) + 1,
		Value: pw,
	}

	go dao.HashAndUpdatePassword(entry.Id, entry.Value)

	stats.RequestTime = time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go stats.UpdateStats(&wg)
	wg.Wait()

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

	if len(dao.Instance) < hashId {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("Hash not found"))
		return
	}

	password := dao.Instance[hashId]

	if password.Value == "" {
		res.WriteHeader(http.StatusAccepted)
		res.Write([]byte("Password is being processed"))
	}

	res.Write([]byte(password.Value))
}

func GetStats(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/stats/" && req.Method != "GET" {
		http.NotFound(res, req)
		return
	}

	j, e := json.Marshal(stats.GetStats())

	if e != nil {
		log.Println(e)
	}

	res.Write(j)
}

func PostShutdown(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/shutdown/" && req.Method != "POST" {
		http.NotFound(res, req)
		return
	}

	res.Write([]byte("{message: 'shutdown signal sent'}"))
}
