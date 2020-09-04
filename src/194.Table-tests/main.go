// 194. Table tests 系列測試? 圖
// 重複 資料夾不能 空格 or 特殊字元

package main

import "fmt"

func main() {
	fmt.Println("2+3=", mySum(2, 3))
	fmt.Println("4+7=", mySum(4, 7))
	fmt.Println("5+9=", mySum(5, 9))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum + 1
}
