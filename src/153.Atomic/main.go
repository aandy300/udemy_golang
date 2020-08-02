// 153.Atomic  原子包 跟 mutex類似?
// cmd > go run -race main.go
// https://golang.org/pkg/sync/atomic/

// cmd跑 有時會有 搶發 的出現 || 這就像事物以不同的速度並行運行並發射它們

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
			runtime.Gosched()
			fmt.Println("記數\t", atomic.LoadInt64(&counter)) // 指向 counter 讀取
			wg.Done()
		}()
		fmt.Println("序", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutine:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
