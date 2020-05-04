//112 interfaces & polymorphism
// func (容器 類型) 接收器
// type 名稱 interface {內容...}

package main

import "fmt"

type person struct {
	fname string
	lname string
}
type scagent struct {
	person
	ltk bool
}
type human interface { //如果你可以執行 speak() 你就有我 human 的 type // 增加type
	speak() // 在speak()的func 接收器裡 增加 可執行speak()的type
}

func (s scagent) speak() {
	fmt.Println("程序speak() 接收器條件為(s scagent):", s.fname, s.lname)
}

//有 此type的人才可以使用
func (p person) speak() {
	fmt.Println("程序speak() 接收器條件為(p person):", p.fname, p.lname)
}

//bar 只能放有 human type的 // 運用方面配合 interface可以多種type
func bar(h human) {
	//type 細分過濾 human 有 scagent 和 person
	switch h.(type) {
	case person:
		fmt.Printf("程序 bar(h human)_caseing_caseh.(類型): %T , 值: %v%v\n", h, h.(person).fname, h.(person).lname)
	case scagent:
		fmt.Printf("程序 bar(h human)_caseing_caseh.(類型): %T , 值: %v%v\n", h, h.(scagent).fname, h.(scagent).lname)
	}
	// fmt.Println("func bar(h hyman)default no switch:h", h)
}

func main() {
	sa1 := scagent{
		person: person{
			fname: "eous(scagent)",
			lname: "gyou(scagent)",
		},
		ltk: true,
	}
	sa2 := person{
		fname: "ming(person)",
		lname: "andy(person)",
	}
	p3 := person{
		fname: "The(person)",
		lname: "FOX(person)",
	}

	sa1.speak()
	sa2.speak()
	p3.speak()
	fmt.Println(p3)
	fmt.Printf("%T\n", sa1)

	bar(sa1)
	bar(sa2)
	bar(p3)

	sa1.speak()

	// 型別轉換 ex sample
	type hotdog int
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	// 型別轉換 x(hotdog) -> y(int)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
