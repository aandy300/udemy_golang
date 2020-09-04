// 197.Benchmark 代碼效率測試工具 圖

// main.go

// Benchmark 中文介紹 代碼效率測試工具
// https://blog.wu-boy.com/2018/06/how-to-write-benchmark-in-go/

// CMD > go help testflag 測試文件

// CMD > go test -bench . 全體測試
// CMD > go test -bench Grreet 單體測試

// -benchmem
// -bench正則表達式
//              僅運行與正則表達式匹配的那些基準。
//              默認情況下，不運行任何基準測試。
//              要運行所有基準，請使用'-bench。'或'-bench =。'。
//              正則表達式由方括號（/）分隔
//              字符分成一系列正則表達式，每個
//              基準標識符的一部分必須與相應的標識符匹配
//              序列中的元素（如果有）。
//              與b.N = 1一起運行以標識子基準。例如，
//              給定-bench = X / Y，運行與X匹配的頂級基準
//              使用b.N = 1來查找與Y匹配的任何子基準
//              然後完全運行。
package main

import (
	"fmt"
	"udemy_golang/src/197.Benchmark/saying"
)

func main() {
	fmt.Println("main.go")
	saying.Greet("ming")
}
