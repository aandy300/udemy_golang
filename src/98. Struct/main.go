//Struct 結構體 像 C# Object 塞一堆參數 其他人也可以與取用
//ex 用資料夾分類 一層一層...
//Struct 中文參考 https://michaelchen.tech/golang-programming/struct/
package main

import "fmt"

//自製類型 沿用類型struct
//設定有哪些 參數 和 參數類型
type person struct {
	f_name string
	l_name string
}

func main() {

	//寫入資料
	x := person{
		f_name: "A",
		l_name: "ndy",
	}
	fmt.Println(x.f_name, x.l_name)
}
