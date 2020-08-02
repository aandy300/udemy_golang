// 150.Documentation 圖 文件概念解說?
// 不是告知 物體 要分享出去
// 而是告知本身帶有分享的情報?
// https://golang.org/ref/spec#Go_statements

// <- = channl 後面會講

package main

import (
	"fmt"
)

func dosomething(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- dosomething(5)
	}()
	fmt.Println(<-ch)
}

// go list.Sort()  // run list.Sort concurrently; don't wait for it.
// A function literal can be handy in a goroutine invocation.

// func Announce(message string, delay time.Duration) {
//     go func() {
//         time.Sleep(delay)
//         fmt.Println(message)
//     }()  // Note the parentheses - must call the function. 注意括號-必須調用該函數。
// }

// Go statements 轉到語句
// “ go”語句作為一個獨立的並發控制線程（或goroutine）在同一地址空間內開始執行函數調用。

// GoStmt =“ go” 表達式。
// 表達式必須是函數或方法調用；不能用括號括起來。內置函數的調用與表達式語句一樣受到限制 。

// 函數值和參數是 在調用goroutine中照常計算的，但是與常規調用不同，程序執行不等待調用的函數完成。相反，該函數開始在新的goroutine中獨立執行。當函數終止時，其goroutine也終止。如果函數具有任何返回值，則在函數完成時將其丟棄。

// go Server（）
// go func（ch chan <-bool）{for {sleep（10）; ch <-true}}（c）
