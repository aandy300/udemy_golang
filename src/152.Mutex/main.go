// 152. Mutex 互斥鎖 lock 防止兩條執行緒同時對同一公共資源進行讀寫的 防止搶輸出
// https://golang.org/pkg/sync/#Mutex
// https://golang.org/pkg/sync/#RWMutex 讀/寫

// Gosched（）圖
// https://studygolang.com/articles/3028

// 鎖住直到編輯結束 其他人才能用
// 在cmd執行時 使用 go run -race main.go
// ide上 goroutines: 3 還是會跑出 cmd 執行不會 Why?

//func Gosched（）
//Gosched产生处理器，从而允许其他goroutine运行。 它不会暫停当前的goroutine，因此执行会自动恢复。

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
	var mu sync.Mutex
	wg.Add(gs) // 增加序 goroutines

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// time.Sleep(time.Second) // 非網頁在 IDE要sleep 不然會跑太快  其他的99還沒done就 結束了 因而counter數量不會是1 因為還有其他的在跑
			// or 先跑其他的?
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done() // = -1 goroutines 序?
		}() // Note the parentheses - must call the function. 注意括號-必須調用該函數。

		fmt.Println("*goroutines:", runtime.NumGoroutine())
	}
	wg.Wait() // 等到追加的序為0 goroutines
	fmt.Println("goroutines", runtime.NumGoroutine())
	fmt.Println("counter", counter)
}
