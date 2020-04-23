//111 Methods 方法?
// func (r receiver) identifier(parameters) (return(s)) {code...}
//讓 有特定type的 元件 可以call 有設定該type的 function
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
type newtype struct {
	nt string
}

func (x person) spec() {
	fmt.Println("hahahahaha")
}

func main() {
	p1 := scagent{
		person: person{
			fname: "aaaa",
			lname: "bbbb",
		},
		ltk: true,
	}
	fmt.Printf("%T\n", p1)
	p1.spec()
	// p2 := newtype{
	// 	nt: "ming",
	// }
	// p2.spec()
	p3 := person{
		fname: "ming",
		lname: "andy",
	}
	p3.spec()

}
