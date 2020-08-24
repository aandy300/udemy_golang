// 186. Hands-on exercise#5 預先練習 TEST pkg

// CI/CD 解說中文
// https://ithelp.ithome.com.tw/articles/10219083

//  https://godoc.org/testing
//  http://www.golang-book.com/books/intro/12
package math

import "testing"

func main() {
	TestAbs(0)
}

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
