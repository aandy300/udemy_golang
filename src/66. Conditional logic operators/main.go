//bool 運算子 條件
package main

import (
	"fmt"
)

func main() {
	//&&：且 (and)
	//||：或 (or)
	//!：非 (not)
	fmt.Println(true && true)  //true
	fmt.Println(true && false) //false
	fmt.Println(true || true)  //true
	fmt.Println(true || false) //true
	fmt.Println(!true)         //false
}
