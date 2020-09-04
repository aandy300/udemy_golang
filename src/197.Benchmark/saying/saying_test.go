// saying test PKG

package saying

import (
	"fmt"
	"testing"
)

// func Greet(){...} 註解
func TestGreet(t *testing.T) {
	s := Greet("ming")
	if s != "welcom here was saying PKG:ming" {
		t.Error("應該為:", "welcom here was saying PKG:ming", "現在為", s)
	}
}

// godoc Exsample Go文件範例
func ExampleGreet() {
	fmt.Println(Greet("ming"))
	// Output:
	// welcom here was saying PKG:ming
}

// 代碼效能測試
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("ming")
	}
}
