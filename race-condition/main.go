package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Cpu: ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	wg.Add(gs)

	//creating mutex to prevent race
	//var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			//lock to prevent access on the variable
			//mu.Lock()
			//Atomic prevent race without using mutex
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter: ", atomic.LoadInt64(&counter))
			//time.Sleep(time.Second)
			runtime.Gosched()
			//unlock to free the variable
			//mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count: ", counter)
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
}
