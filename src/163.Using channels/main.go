// 163. Using channels
// 分工  <-放入數字 <-放入整個channel

// 建立一般channels
// 在funcs中，您可以指定
// * 接收频道 receive channel
// 		您可以从通道接收值
// 		接收通道参数
// 		在功能中，您只能从通道中提取值
// 		您无法关闭接收频道
// * 发送频道 send channel
// 		您可以将值推送到通道
// 		您无法从频道接收/拉取/读取
// 		您只能将值推送到通道

package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)
	//  如果 bar 也 go bar(c) 需要用waitgroup 和 加序 才可以印出? go 最多使用一個(不使用WAIt的話?)
	bar(c)

	fmt.Println("out")
}

func foo(c chan<- int) {
	c <- 3
}
func bar(c <-chan int) {
	fmt.Println(<-c)
}
