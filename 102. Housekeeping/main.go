/*102.Housekeeping 概念 概論 OOP

go 官方Q&A
https://golang.org/doc/faq#Is_Go_an_object-oriented_language
go packages 官方文件
https://golang.org/pkg/

Golang 一切都是有關於type , 不能同時有2個type
資料結構,傳統OOP

1. Code 清潔明瞭 優先考可讀性
2. 省資源 ex int int64 flow32 etc...*/

package main

import "fmt"

func main() {
	p1 := struct {
		f_name string
		l_name string
		age    int
	}{
		f_name: "a",
		l_name: "ndy",
		age:    27,
	}
	fmt.Println(p1)

}
