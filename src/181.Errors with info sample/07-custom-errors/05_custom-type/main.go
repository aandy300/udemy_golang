// 181. Errors with info  !!!圖 與 練習#3 類似
// 05 02+03+04 儲存起來發用 自訂錯誤訊息? type

// 中文解說 https://michaelchen.tech/golang-programming/error-handling/

// https://golang.org/pkg/errors/#example_New
// https://golang.org/src/errors/errors.go
package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("\n錯誤訊息:%v,%v,%v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		fmt.Println("123")
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("\t錯誤的值為：%v", f)
		return 0, norgateMathError{"\n\t  N", " W", nme}
	}
	return 42, nil
}

// 在標準庫中查看具有錯誤類型的結構的使用：
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
