// 181. error sample
// 02

/*
程序包日誌實現了一個簡單的日誌記錄程序包...寫入標準錯誤並打印每個記錄的消息的日期和時間...
*/

// log.Println調用Output打印到標準記錄器。 以fmt.Println的方式處理參數。

package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
}

/*
Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ...
*/

// log.Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
