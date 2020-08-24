// 178.Checking erros err sample || panic log 錯誤回報 輸出
// 中文解說
// http://legendtkl.com/2016/03/11/go-log/

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("-錯誤 log.Println_nofile.txt-", err)
	}
	defer f2.Close()

	log.SetOutput(f)

	// if err != nil {
	// 	// 回報類型
	// 	// fmt.Println("-錯誤 fmt.Println-", err)
	// 	log.Println("-錯誤 log.Println-", err)
	// 	// log.Fatalln("-錯誤 log.Fatalln-",err)
	// 	// panic(err)
	// }
}
