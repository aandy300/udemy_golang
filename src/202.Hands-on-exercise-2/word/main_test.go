// word
package word

import (
	"fmt"
	"testing"
	"udemy_golang/src/202.Hands-on-exercise-2/quote"
)

// strings.Fields(s) + mpa[v]+++ sample
// https://play.golang.org/p/u0rAH4zd4Ff

func TestUseCount(t *testing.T) {
	// switch 看遇到的 key 對照路線  m 已經在 useCount(...) 裡面 轉成 m[v]++ 了
	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", "1")
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", "1")
			}
		case "three":
			if v != 3 {
				t.Error("got", v, "want", "3")
			}
		}
	}
}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("got", n, "want", 3)
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
