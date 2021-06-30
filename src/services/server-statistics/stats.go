package stats

import (
	"sync"
	"time"
)

var Total int64
var Avg int64
var RequestTime time.Time
var resource sync.Mutex

type ServerStats struct {
	Total int64
	Average int64
}

func UpdateStats(wg *sync.WaitGroup) {
	resource.Lock()
	Total += 1

	Avg = Avg + time.Since(RequestTime).Microseconds()
	resource.Unlock()
	wg.Done()
}

func GetStats() ServerStats {
	return ServerStats{Total, Avg}
}




