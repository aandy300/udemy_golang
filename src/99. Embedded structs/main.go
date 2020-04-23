//Struct 結構體 中 再放結構體
//ex 用資料夾分類 一層一層...
//Struct 中文參考 https://michaelchen.tech/golang-programming/struct/
package main

import "fmt"

//自製類型 沿用類型struct
//設定有哪些 參數 和 參數類型
type scagent struct {
	person
	ltk bool
}

type person struct {
	f_name string
	l_name string
}

func main() {

	//寫入資料
	//寫入其他struct時要進去
	agent1 := scagent{
		person: person{
			f_name: "A",
			l_name: "ndy",
		},
		ltk: true,
	}

	p2 := person{
		f_name: "m",
		l_name: "ing",
	}

	fmt.Println(agent1)
	//讀取其他struct時不用進去可以直接取用
	//要進去也可以高岡而已
	//ex 沒進去
	fmt.Println(agent1.f_name, agent1.l_name)
	//ex 進去
	fmt.Println(agent1.person.f_name, agent1.person.l_name)

	fmt.Println(p2.f_name, p2.l_name)
}
