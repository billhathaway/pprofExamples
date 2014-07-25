package main

import (
	"github.com/billhathaway/pprofExamples/badLRU"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	f, _ := os.Create("badLRU.profile")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	size := 1000
	cache := badLRU.New(size)
	for i := 0; i < size; i++ {
		cache.Put(i, i)
	}
	for i := 0; i < size*1000; i++ {
		cache.Get(rand.Intn(size * 2))
		cache.Get(rand.Intn(size * 2))
		cache.Get(rand.Intn(size * 2))
		cache.Get(rand.Intn(size * 2))
		cache.Get(rand.Intn(size * 2))
		cache.Get(rand.Intn(size * 2))
		cache.Put(i, i)
	}
}
