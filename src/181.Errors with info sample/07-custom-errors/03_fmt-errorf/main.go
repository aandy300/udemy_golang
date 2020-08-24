// 181. Errors with info
// 03 印出錯誤的值
// Errorf根據格式說明符進行格式化，並以
//滿足錯誤的值。
//
//如果格式說明符包含帶有錯誤操作數的％w動詞，
//返回的錯誤將實現Unwrap方法，返回操作數。 它是
//包含多個％w動詞或為其提供操作數無效
//沒有實現錯誤接口的代碼。 ％w動詞不然
//％v的同義詞。

// https://golang.org/pkg/errors/#example_New
// https://golang.org/src/errors/errors.go
package main

import (
	"fmt"
	"log"
)

func main() {

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("再次進行數學運算：負數的平方根:%v", f)
	}
	return 42, nil
}
