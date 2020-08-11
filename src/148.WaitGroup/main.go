// 148. WaitGroup 等待  分序列執行 防止其他序未結束就結束程序
// wg.Add(1) // WaitGroup裡添加1個goroutine
// wg.Wait() // 如果沒有 wait 程序會直接結束 接收 到 wg.Done 才繼續執行下去
// wg.Done()
// go foo()  // go xxx() = 增加使用一條序?
// go ***() 優先順序為2 沒 "go" ***()得先跑 ex bar() ? or go 先跑?  165. go func()..先跑

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)                 // os windows or linux
	fmt.Println("ARCH\t\t", runtime.GOARCH)             // cpu type
	fmt.Println("CPUs\t\t", runtime.NumCPU())           // cpu有幾序
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) // 使用的序列

	wg.Add(1) // WaitGroup裡添加1個goroutine
	go foo()  // go xxx() = 增加使用一條序? Go statements?
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait() // 如果沒有 wait 程序會直接結束 接收 到 wg.Done 才繼續執行下去
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
