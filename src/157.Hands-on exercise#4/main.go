// 157.Hands-on exercise#4 Mutex 使用互斥锁修复您在上一个练习中创建的竞赛条件

// 鎖住 會用到的值 跑完再給其他程序使用該值 另類取代 runtime.Gosched() ?

// cmd go run -race main.go

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(gs) // WaitGroup裡添加goroutine

	for i := 0; i < gs; i++ {
		// fmt.Println("star||count:", counter) //會殘留 go race ? 註解掉就不會有 Found 1 data race(s) 影響 data race 沒跑完?
		go func() {
			mu.Lock()
			v := counter // 記數用
			// time.Sleep(time.Second)
			v++
			counter = v
			fmt.Println("count:", counter)
			mu.Unlock()
			wg.Done() // goroutine -1
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait() // 等待到 goroutine = 0 or 1? 執行以下
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
