package main

import (
	"log"
	"runtime"
)

// map中删除一个key，key在map中所占内存不会立即释放。
var intMap map[int][]int
var cnt = 10
var s = make([]int, 5, 10)

func main() {
	printMemStats()

	initMap()
	runtime.GC()
	printMemStats()

	log.Println(len(intMap))
	for i := 0; i < cnt; i++ {
		delete(intMap, i)
	}
	log.Println(len(intMap))

	runtime.GC()
	printMemStats()

	intMap = nil
	runtime.GC()
	printMemStats()
}

func initMap() {
	intMap = make(map[int][]int, cnt)
	s = []int{1, 3, 5, 7, 9}
	for i := 0; i < cnt; i++ {
		intMap[i] = s
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}
