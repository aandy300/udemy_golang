// 155.Hands-on exercise#2 此练习将加强我们对方法集的理解 method

// A繼續往下走因為沒通行證human所以不能過
// B回源頭指向human 因為有碰到所以判定可以過(interface有type OR 能用func***()就算)?

// receiver = 接收器

package main

import (
	"fmt"
)

type Person struct {
	First string
}

type Human interface {
	say()
}

func main() {
	p1 := Person{
		First: "p1a",
	}

	//因為往回指向& 有碰到human所以 算有type(interface) 所以給過?
	saysomething(&p1)

	// 不能用因為沒有type
	// saysomething(p1)

	p1.say()
}

func (p *Person) say() {
	fmt.Println("hi")
}

func saysomething(h Human) {
	h.say()
}
