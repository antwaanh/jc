package queue

import (
	"container/list"
	"fmt"
)

var Instance *list.List

// watch queue
func Manager() {
	for e := Instance.Front(); e != nil; e = e.Next() {
		fmt.Println("Watching...", e)
	}
}
