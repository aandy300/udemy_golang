// 178.Checking erros err sample || panic log 錯誤回報

// https://golang.org/pkg/builtin/#panic

// log.Println()
// log.Fatalln() // 印出錯誤後 結束 exit
// 	os.Exit()
// log.Panicln() // 可以回復
// 	deferred functions run
// 	can use “recover”

// panic() CMD 跑才 看的到錯誤訊息
// panic() // 在 "當前goroutine" 增加deferr 最後執行 nil 0 改成 1 程式繼續
// panic 中文解說 底部小節 https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-panic-recover/

package main

import (
	"os"
)

func main() {
	// f, err := os.Create("log.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer f.Close()

	// log.SetOutput(f)

	// f2, err := os.Open("no-file.txt")
	// if err != nil {
	// 	// log.Println("-錯誤 log.Println_nofile.txt-", err)
	// 	panic(err)
	// }
	// defer f2.Close()

	_, err := os.Open("no-file.txt")
	if err != nil {
		// 回報類型
		// fmt.Println("-錯誤 fmt.Println-", err)
		// log.Println("-錯誤 log.Println-", err)
		// -----------------
		// log.Fatalln("-錯誤 log.Fatalln-", err)
		// Fatalln()
		// Fatalln()這個function就是兩行code的組合，這個function是包在"log"中的。
		// 其實就是執行這兩行，Print出錯誤內容後結束程式
		// defer 不會跑
		// fmt.Println(err)
		// os.Exit(1)
		// -----------------

		panic(err)
	}
}
