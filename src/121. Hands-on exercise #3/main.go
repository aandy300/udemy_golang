//121.#3 defer練習 延遲置程序最尾端 // 1 2 3 0
package main

import "fmt"

func main() {
	defer fmt.Println("0")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}
