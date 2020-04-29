//122. #4 自訂 type 參數 struct 練習 (容器 允許的type)func 過濾選擇器
package main

import "fmt"

type person struct {
	f_name string
	l_name string
	age    int
}

type worker struct {
	person
	tool string
}

func main() {
	p := person{
		f_name: "A",
		l_name: "ndy",
		age:    27,
	}

	w1 := worker{
		person: person{
			f_name: "Ming",
			l_name: "lee",
			age:    30,
		},
		tool: "computer",
	}

	fmt.Println(p, w1)
	p.speak()
	w1.speak()
}

func (p person) speak() {
	fmt.Println("hi", "i am", p.f_name, p.l_name, "my age was", p.age)
}
