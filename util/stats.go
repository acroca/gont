package util

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type stats struct {
	Frames int
	Steps  int
	Mutex  sync.Mutex
}

var (
	// Stats stores the app stats
	Stats = stats{}
	m     runtime.MemStats
)

func (stats *stats) Start() {
	for {
		time.Sleep(1 * time.Second)
		stats.Print()
	}
}
func (stats *stats) Print() {
	stats.Mutex.Lock()
	defer stats.Mutex.Unlock()
	runtime.ReadMemStats(&m)
	fmt.Printf("FPS: %d\tSPS: %d\tMem: %d\n", stats.Frames, stats.Steps, m.Alloc)
	stats.Frames = 0
	stats.Steps = 0
}
