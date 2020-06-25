//132. method sets 指標 印值ver dc圖

// 直接印值ver https://play.golang.org/
// 指標ver https://play.golang.org/p/glWZmm0gY6

// & = gives you the address 記憶體 " 位置 "
// * = 代表顯示該記憶體位置的值
package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	fmt.Println(math.Pi)
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	info(&c)
}
