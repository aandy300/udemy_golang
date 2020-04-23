//104 struct 結構體練習 加入map for loop 印出資料
//map 在 for loop i = key
//map 在 注意key不要同名 跑圈會錯誤

package main

import "fmt"

type person struct {
	fname string
	lname string
	favo  []string
}

func main() {

	p1 := person{
		fname: "andy",
		lname: "lee",
		favo:  []string{"chaco", "chaco2", "chaco3"},
	}
	p2 := person{
		fname: "ming",
		lname: "leee",
		favo:  []string{"mint1", "mint2", "mint3"},
	}

	//寫入map
	m := map[string]person{
		p1.lname: p1,
		p2.lname: p2,
	}

	//印出map
	//i on map = key / i2 on array = index
	for i, v := range m {
		fmt.Println(i, v, "i am key")
		for i2, v2 := range v.favo {
			fmt.Println(i2, "v2", v2)
		}
	}
}
