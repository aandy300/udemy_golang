// 181. error sample
// 01
// Println格式使用其操作數的默認格式，並寫入標準輸出。

package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened:", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
}

// Println formats using the default formats for its operands and writes to standard output.
