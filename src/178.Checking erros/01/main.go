// 178.Checking erros

package main

import (
	"fmt"
)

func main() {
	// 因 println 所以是6 +1
	n, err := fmt.Println("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
