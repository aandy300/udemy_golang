// 194. Table tests 系列測試?
// math_test.go

package main

import "testing"

func TestMath(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{1, 2, 3}, 6},
		test{[]int{30, 22, 4}, 56},
		test{[]int{22, 19, 21}, 62},
	}
	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("錯誤不等於", v.answer, "現在為:", x)
		}
	}
}
