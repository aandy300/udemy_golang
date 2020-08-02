// 149.  Method sets revisited　圖
// https://golang.org/ref/spec#Method_sets
// 接收器是 *pointer   值就是 *pointer  (at pointer)
// “类型的方法集确定该类型实现的接口...。” 類型要相同?
// goroutine有一个简单的模型：它是一个与其他goroutine在同一地址空间中同时执行的函数。

// “类型的方法集确定该类型实现的接口.....”〜golang spec

// Receivers       Values
// -----------------------------------------------
// （t T）    	T和*T
// （t *T）		*T

package main

import (
	"fmt"
)

func main() {
	fmt.Println("講解")

}
