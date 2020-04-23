//運算子 bool || &&
package main

import "fmt"

var x string = "Ming"

func main() {
	fmt.Println(true && true)  //true
	fmt.Println(true && false) //false
	fmt.Println(true || true)  //true
	fmt.Println(true || false) //true
	fmt.Println(!true)         //false
}
