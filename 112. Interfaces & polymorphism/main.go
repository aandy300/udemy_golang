//112 interfaces & polymorphism
// func (* type) 接口
// interface {允許目標}
// bar switch csse...
// bar (h human) -> interface允許的範圍?

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
type human interface { //如果你可以執行 speak() 你就有我 human 的 type // 篩選?
	speak() //允許條件 包含? 使用到他的都可以允許?
}
type ntype struct { //測試用type 可刪
	ltlook bool
}

func (s scagent) speak() {
	fmt.Println("func speak() i get type scagent", s.fname, s.lname)
}

//有 此type的人才可以使用
func (p person) speak() {
	fmt.Println("func speak() i get type person", p.fname, p.lname)
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Printf("func bar(h human)_caseing_caseh.(type): %T , value was: %v%v\n", h, h.(person).fname, h.(person).lname)
	case scagent:
		fmt.Printf("func bar(h human)_caseing_caseh.(type): %T , value was: %v%v\n", h, h.(scagent).fname, h.(scagent).lname)
	}
	//fmt.Println("func bar(h hyman)default no switch", h)
}

//可以直接使用
func testfunc() { //測試用 可刪
	fmt.Println("i am test func")
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
	p4 := ntype{}
	fmt.Println(p4)

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
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
