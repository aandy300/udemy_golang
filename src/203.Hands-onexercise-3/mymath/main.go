package mymath

import (
	"fmt"
	"sort"
)

// CenteredAvg ...
// sort 排序更改 > 拿中間存入xi > 加總數字存入 n > n / xi長度(len)
// CenteredAvg computes the average of a list of numbers
// after removing the largest and smallest values.
func CenteredAvg(xi []int) float64 {
	fmt.Println("xi1:", xi)
	sort.Ints(xi)
	fmt.Println("xi2:", xi)
	xi = xi[1:(len(xi) - 1)]
	fmt.Println("xi3:", xi)
	n := 0
	for _, v := range xi {
		n += v
	}
	fmt.Println("N:", n)
	fmt.Println("float64xi:", float64(len(xi)))
	f := float64(n) / float64(len(xi))
	return f
}
