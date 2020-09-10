// 202.Hands-on-exercise#2 同樣練習 條件限制 複習 拆文字統計數字
// test
// benchmarks
// coverage
// coverage net
// exsample

package main

import (
	"fmt"
	"udemy_golang/src/202.Hands-on-exercise-2/quote"
	"udemy_golang/src/202.Hands-on-exercise-2/word"
)

func main() {
	// 印出所有文字 word.Count(quote.SunAlso)  = 全文本 const name SunAlso
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		// 印出 回傳來的 m = map[string]int || key value
		fmt.Println("key", k, "value", v)
	}
}
