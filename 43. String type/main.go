package main

import "fmt"

func main() {
	x := "hello word haaaaaa"
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = "heloo tw"

	// 將 x byte10進位 存入到 bs[]
	// 印出 bs[]
	// 印出格式
	bs := []byte(x)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// 新增容器裝計數器 x helloword的長度 計數器+1
	for i := 0; i < len(x); i++ {
		// 把x字串變成陣列 將第一個字 的 %#U UTF-8 編碼 印出
		// 印出utf-8代碼 目標為 x[i]
		fmt.Printf("%#U", x[i])
	}
	//清除格式? 換行?
	fmt.Println("")

	//for loop 跑 x[]數量 v = x陣列的數量
	//%d = 陣列的編號 %#x = 16進
	for i, v := range x {
		fmt.Printf("index postion %d 十六進制 %#x\n", i, v)
	}
}
