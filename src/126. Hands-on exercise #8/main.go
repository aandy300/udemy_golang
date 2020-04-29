//#126 .8 func return fun()
package main

import "fmt"

func main() {

	f := foo()()
	fmt.Println(f)

}

func foo() func() int {
	fmt.Println("hi")
	return func() int {
		fmt.Println("hi2")
		return 8
	}
}
