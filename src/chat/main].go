package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	//println("Memory usage before function execution:")
	//printMemUsage()

	var wg = &sync.WaitGroup{}
	defer func() {
		wg.Wait()
	}()

	wg.Add(1)
	go func() {
		println("Huan Hello, World!")
		wg.Done()
	}()
	//wg.Add(1)
	go func() {
		println("Phuong Hello, World!")
		time.Sleep(1 * time.Second)
		println("done Phuong")
		wg.Done()
	}()

	//println("Memory usage after function execution:")
	//printMemUsage()
}
