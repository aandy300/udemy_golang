//111 Methods 方法?
//      接收器 接收條件 func名稱  參數         返回      內容
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
	fmt.Println("have type person , can call me")
}

func main() {
	//scagent 包含 person 所以也可以
	p1 := scagent{
		person: person{
			fname: "aaaa",
			lname: "bbbb",
		},
		ltk: true,
	}
	fmt.Printf("%T\n", p1)
	p1.spec()

	// p2 type不一樣 call不了
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
