package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	age := Years(10)
	if age != 70 {
		t.Error("應該為:", 70, "現在為:", age)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}
