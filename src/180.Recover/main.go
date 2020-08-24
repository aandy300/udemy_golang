// 180.Recover 復原 停止當前 往上尋找 返回調用方

// 圖

// Panic是一個內置函數，可停止常規控制流並開始恐慌。
// 當函數F調用恐慌時，F的執行停止，F中任何延遲的函數都將正常執行，然後F返回其調用方。
// 對於呼叫者，F然後表現得像是發生了恐慌。
// 該過程將繼續執行堆棧，直到返回當前goroutine中的所有函數為止，此時程序崩潰。緊急情況可以通過直接調用緊急情況來啟動。
// 它們也可能是由運行時錯誤引起的，例如越界數組訪問。

// Recover是一個內置函數，可以重新獲得對緊急恐慌例程的控制。恢復僅在延遲函數內部有用。
// 在正常執行期間，恢復調用將返回nil並且沒有其他效果。
// 如果當前goroutine處於恐慌狀態，則調用recovery會捕獲提供給panic的值並恢復正常執行。

// https://blog.golang.org/defer-panic-and-recover

package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		// 我們進入調用堆棧，首先執行任何延遲的語句。
		panic(fmt.Sprintf("333%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// 周圍的函數返回後，將按照後進先出的順序執行延遲的函數調用。
// 該函數顯示“ 3210”：

// func b() {
//     for i := 0; i < 4; i++ {
//         defer fmt.Print(i)
//     }
// }

// 在此示例中，延遲函數 在周圍函數返回後遞增返回值i 。因此，此函數返回2：

// func c() (i int) {
//     defer func() { i++ }()
//     return 1
// }
