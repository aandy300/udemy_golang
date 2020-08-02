// 156.Hands-on exercise#3  序 賽跑 並 更改 值 ++

// build 生 exe || go build -race

// cmd > ./ 这意味着在Mac上的该目录端口中，文件执行它。

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs) // WaitGroup裡添加goroutine

	for i := 0; i < gs; i++ {
		fmt.Println("star||count:", counter)
		go func() {
			v := counter // 記數用
			// time.Sleep(time.Second)
			runtime.Gosched() // 产生处理器，从而允许其他goroutine运行。 它不会暫停当前的goroutine，因此执行会自动恢复。
			v++
			counter = v
			wg.Done() // goroutine -1
			fmt.Println("done||count:", counter)
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait() // 等待到 goroutine = 0 or 1? 執行以下
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
