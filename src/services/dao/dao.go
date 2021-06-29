package dao

var Instance = map[int]string{}

type PasswordEntry struct {
	Id        int
	Value     string
	CreatedAt int64
}

func PersistPassword(id int, pw string) {
	Instance[id] = pw
}
