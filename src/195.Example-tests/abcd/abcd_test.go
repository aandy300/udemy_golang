// output:要在下面一行才作用

// PKG abcd 註解
package abcd

import (
	"fmt"
)

// func Sum(){...} 註解
func ExampleSum() {
	fmt.Println(Sum(2, 3))
	// Output:
	// 5
}
