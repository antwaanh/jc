package dao

import (
	"crypto/sha512"
	"encoding/base64"
	"time"
)

var Instance = map[int]PasswordEntry{}

type PasswordEntry struct {
	Id        int
	Value     string
	CreatedAt int64
}

func StorePassword(id int, pw string) {
	Instance[id] = PasswordEntry{Value: pw, CreatedAt: time.Now().Unix()}
}

func UpdatePassword(id int, pw string) {
	Instance[id] = PasswordEntry{Value: pw}
}

func HashAndUpdatePassword(id int, pw string) {
	StorePassword(id, "")

	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	hash := sha512.Sum512([]byte(pw))

	UpdatePassword(id, base64.URLEncoding.EncodeToString(hash[:]))
}

