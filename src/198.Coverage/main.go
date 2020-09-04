// 198.Coverage 涵蓋量? 這裡只有解說 c.out 在 197../saying

// go help testflag

// go test -cover
// -cover
// 啟用覆蓋率分析。
//              請注意，因為覆蓋範圍通過註釋源起作用
//              編譯，編譯和測試失敗之前的代碼
//              啟用的覆蓋範圍可能報告的行號不對應
//              原始來源。

// go -coverprofile c.out
// -coverprofile cover.out
//              所有測試通過後，將coverage配置文件寫入文件輸出 c.out。
//              設置-cover。

// go tool cover -html=c.out
// 開預覽器 看覆蓋率 (方便閱讀)
// 原理 > go tool > go tool cover > -html=c.out > 生成coverage配置文件的HTML表示

package main

import (
	"fmt"
)

func main() {
	fmt.Println("main.go")
}
