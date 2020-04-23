//整束縛點數差異

//平台環境確認
//golang 編譯各平台 中文解說
//https://www.echochio.nctu.me/2018/06/golang_multiple_platforms/
package main

import (
	"fmt"
	"runtime"
)

func main() {

	x := 42
	y := 42.222
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
