// 181. Errors with info
// 04 +03 追加容器儲存起來用

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
		ErrNorgateMath := fmt.Errorf("再次進行數學運算：負數的平方根:%v", f)
		return 0, ErrNorgateMath
	}
	return 42, nil
}
