package main

import (
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup
	var counter int32 = 0

	wg.Add(3)

	go func() {
		atomic.AddInt32(&counter, 1)
		wg.Done()
	}()
	go func() {
		v := counter
		v++
		counter = v
		wg.Done()
	}()
	go func() {
		v := counter
		v++
		counter = v
		wg.Done()
	}()

	wg.Wait()
}
