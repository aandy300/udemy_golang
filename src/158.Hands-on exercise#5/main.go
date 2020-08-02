// 158.Hands-on exercise#5 atomic 原子包 race
// https://godoc.org/sync/atomic
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("開始,cpus:", runtime.NumCPU())
	fmt.Println("開始,goroutines", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1) // 指向 counter +1
			// runtime.Gosched()
			fmt.Println("記數\t", atomic.LoadInt64(&counter)) // 讀 counter
			wg.Done()
		}()
		fmt.Println("序", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutine:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
