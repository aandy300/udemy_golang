// 159.Hands-on exercise#6
// cmd go run
// go build  > 生EXE
// go install > 生到 go\bin

// https://dotblogs.com.tw/hochile/2020/02/04/135624 解說

//GOARCH是正在運行的程序的體系結構目標：386，amd64，arm，s390x等之一。
// 這個常數可以識別在哪麼CPU架構上
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// windows amd64 顯示CPU類型?
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
