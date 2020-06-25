//134. Hands-on exercucise #2
//生出個值 印出 更改該位置的值

// & = gives you the address 記憶體 " 位置 "
// * = 代表顯示該記憶體位置的值
package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "andy",
	}
	fmt.Println(p1)
	changename(&p1)
	fmt.Println(p1)
}

func changename(p *person) {
	p.name = "ming"
	// (*p).name = "ming"
}
