// 154.Hands-on exercise#1
// 增加序 使用 Waitgroup 等待程序 防止程序直接結束

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

	wg.Add(2) // WaitGroup裡添加goroutine
	go foo()  // go xxx() = 增加使用一條序? Go statements?
	go haa()
	bar()

	wg.Wait() // 如果沒有 wait 程序會直接結束 接收 到 wg.Done 才繼續執行下去
	fmt.Println("---------end")
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 2; i++ {
		fmt.Println("foo:", i)
	}
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Done()
}
func haa() {
	for i := 0; i < 2; i++ {
		fmt.Println("haa:", i)
	}
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Done()
}
func bar() {
	for i := 0; i < 2; i++ {
		fmt.Println("bar:", i)
	}
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}
