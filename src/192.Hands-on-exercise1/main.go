// 191. Writing documentation 打包PKG 註解位置

// import PKG路徑 不能特殊字元和空格  x *#$&#

package main

import (
	"fmt"
	"udemy_golang/src/192.Hands-on-exercise1/m_Yearchange"
)

func main() {
	fmt.Println("1")
	m_Yearchange.Yearchange(2)
}
