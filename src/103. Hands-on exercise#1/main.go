//103 struct 結構體練習 for loop印出資料
package main

import "fmt"

type person struct {
	fname string
	lname string
	favo  []string
}

func main() {
	p1 := person{
		fname: "a",
		lname: "ndy",
		favo:  []string{"chaco", "chaco2", "chaco3"},
	}
	p2 := person{
		fname: "m",
		lname: "ing",
		favo:  []string{"mint1", "mint2", "mint3"},
	}
	fmt.Println(p1.fname, p1.lname, p1.favo)
	fmt.Println(p2.fname, p2.lname, p2.favo)

	for i, v := range p1.favo {
		fmt.Println("p1喜歡的口味順序", i, v)
	}
	for i, v := range p2.favo {
		fmt.Println("p2喜歡的口味順序", i, v)
	}
}
