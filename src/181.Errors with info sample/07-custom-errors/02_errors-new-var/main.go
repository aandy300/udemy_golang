// 181. Errors with info
// 02 +01 多加一個容器儲存 方便使用

// https://golang.org/pkg/errors/#example_New
// https://golang.org/src/errors/errors.go

// 查看錯誤的使用。標準庫中的新增功能：
// http://golang.org/src/pkg/bufio/bufio.go
// http://golang.org/src/pkg/io/io.go
package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorgateMath = errors.New("錯誤訊息:")

func main() {
	//  *errors.errorString
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return 42, nil
}
