// 191. Writing documentation 打包PKG 註解位置
package m_Yearchange

import (
	"fmt"
)

// Yearchange() 備註
func Yearchange(xi int) int {
	sum := xi * 7
	fmt.Println("here is in -Yearchange pkg- u send:", xi)
	fmt.Println("human year:", xi, "dog year:", sum)
	return sum
}
