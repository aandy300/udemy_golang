// 184.Hands-on exercise#3 error interface 自訂 Error() 接口

// e.(customErr).s 這個動作是 Assertion 斷言 || 底部解說?

// 流程看圖解

package main

import "fmt"

type customErr struct {
	s string
}

func (cErr customErr) Error() string {

	// 如果使用:
	// return fmt.Println("tex")
	// 錯誤:
	// too many arguments to return
	// have (int, error)
	// want (string)
	// 因為: func Printf(format string, a ...interface{}) (n int, err error)
	// return int error 2個 但 Error() 只收 string 一個
	// 解: 使用 func Sprintf(format string, a ...interface{}) string
	return fmt.Sprintf("\nhere at Error() 錯誤: %v", cErr.s)

}

func main() {
	p1 := customErr{
		s: "hallo",
	}
	foo(p1)
}

func foo(e error) {
	// fmt.Printf("%v\n %t\n", e, e)
	fmt.Println("\nhere at foo():", e, "\n", e.(customErr).s)
	// 如果使用:
	// fmt.Println(e.s)
	// 錯誤:
	// e.s undefined (type error has no field or method s)
	// e.s未定義（類型錯誤沒有字段或方法）
	// 因為:
	// 在這裡 e 是 type error 而
	// type error {
	// 	Error() string
	// }
	// 只是個接口 沒有存 string 的容器?
	// 解:
	// 追加有 string 容器的 type
	// fmt.Println("\n", e.(customErr).s)

	// e.(customErr).s 這個動作是 Assertion 斷言
	// 就像您有實現接口的東西，而您需要斷言它實際上是您知道實現該接口的另一種類型。
	// 白話 追加 type?
	// 跟型別轉換不同 型別轉換 是 換 type
}
