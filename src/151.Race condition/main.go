// 151. Race  condition  cpu 序 分頭執行 賽跑 搶輸出?

// 在cmd執行時 使用 go run -race main.go

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("cpus:", runtime.NumCPU())
	fmt.Println("goroutines", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs) // 增加序 goroutines

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second) // 非網頁在 IDE要sleep 不然會跑太快  其他的99還沒done就 結束了 因而counter數量不會是1 因為還有其他的在跑
			// or 先跑其他的?
			runtime.Gosched()
			v++
			counter = v
			wg.Done() // = -1 goroutines 序?
		}() // Note the parentheses - must call the function. 注意括號-必須調用該函數。
		fmt.Println("goroutines", runtime.NumGoroutine())
	}
	fmt.Println("goroutines", runtime.NumGoroutine())
	wg.Wait() // 等到追加的序為0 goroutines
	fmt.Println("counter", counter)
}
