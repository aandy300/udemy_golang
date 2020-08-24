// 181. error sample
// 04

/*
...致命功能在寫入日誌消息後調用os.Exit（1）...
*/

// Fatalln等效於Println（），然後調用os.Exit（1）。

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		log.Fatalln(err)
		//		panic(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}

/*
... the Fatal functions call os.Exit(1) after writing the log message ...
*/

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
