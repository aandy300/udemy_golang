//101.其他種 STRUCT
//如果只需要使用單次 宣在程式碼內
//如果多次使用 不建議這樣使用
//多次使用還是乖乖 宣好宣完整程式碼比較乾淨

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
