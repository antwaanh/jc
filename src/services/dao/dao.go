package dao

import (
	"crypto/sha512"
	"encoding/base64"
	"jc/src/services/config"
	"sync"
	"time"
)

var Instance = map[int]PasswordEntry{}

type PasswordEntry struct {
	Id        int
	Value     string
	CreatedAt int64
}

func StorePassword(id int, pw string, resource *sync.Mutex) {
	resource.Lock()
	Instance[id] = PasswordEntry{Value: pw, CreatedAt: time.Now().Unix()}
	resource.Unlock()
}

func UpdatePassword(id int, pw string, resource *sync.Mutex) {
	resource.Lock()
	Instance[id] = PasswordEntry{Value: pw}
	resource.Unlock()
}

func HashAndUpdatePassword(id int, pw string, resource *sync.Mutex) {
	hashInterval, _ := time.ParseDuration(config.GetEnv("HASH_INTERVAL"))
	StorePassword(id, "", resource)

	timer := time.NewTimer(hashInterval * time.Second)
	<-timer.C
	hash := sha512.Sum512([]byte(pw))

	UpdatePassword(id, base64.URLEncoding.EncodeToString(hash[:]), resource)
}
