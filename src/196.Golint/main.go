// 196.Golint 檢查代碼工具 回報不良代碼風格

// CMD > golint ./...

// github https://github.com/golang/lint
// install go get -u golang.org/x/lint/golint

// CSDN Golint介紹
// https://blog.csdn.net/cgl1079743846/article/details/90665161

// gofmt
// formats go code 格式化code
// go vet
// reports suspicious constructs 回報可疑結構
// golint
// reports poor coding style 回報不良代碼風格

package main

import (
	"fmt"
)

func main() {
	fmt.Println("main.go")
}
