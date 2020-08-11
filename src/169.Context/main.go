// 169. Context 上層出現錯無時 將訊號同步給下層 圖
// 在Go服務器中，每個傳入請求都在其自己的goroutine中進行處理。 請求處理程序通常會啟動其他goroutine來訪問後端，例如數據庫和RPC服務。 處理請求的goroutine集合通常需要訪問特定於請求的值，例如最終用戶的身份，授權令牌和請求的截止日期。 當一個請求被取消或超時時，處理該請求的所有goroutine應該迅速退出，以便系統可以回收他們正在使用的任何資源。 在Google，我們開發了一個上下文包，可以輕鬆地將跨API邊界的請求範圍的值，取消信號和截止日期傳遞給處理請求所涉及的所有goroutine。 該軟件包可作為上下文公開使用。 本文介紹瞭如何使用該程序包，並提供了一個完整的工作示例。

// 中文解說
// https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/

// https://blog.golang.org/context
// https://golang.org/pkg/context/
// 取消 愈時請求

// 簡單版 exsample https://play.golang.org/p/FPRcTKoLrb_u

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num gortins 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num gortins 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	// 取消 關閉序?
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num gortins 3:", runtime.NumGoroutine())
}
