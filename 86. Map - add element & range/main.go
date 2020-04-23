//map 更改 新增 檢查 字典 for map slice運用
package main

import "fmt"

func main() {
	m := map[string]int{
		"A": 12222,
		"B": 2,
	}
	fmt.Println("字典原本:", m)

	v, k := m["A"] //檢測是否在字典 1 TRUE / 0 FALSE
	fmt.Println("檢查是否在字典裡:", v, k)

	m["B"] = 9 //更改資料
	fmt.Println(`m["B"]更改資料後:`, m["B"])

	m["C"] = 3 //新增資料
	fmt.Println("新增資料後:", m)

	//for map運用
	fmt.Println("for map---")
	for i, d := range m {
		fmt.Println(i, d)
	}
	fmt.Println("for map---")

	//for slice運用
	fmt.Println("for clip---")
	cli := []int{4, 5, 6, 7, 8, 9}
	for i, z := range cli {
		fmt.Println(i, z)
	}
	fmt.Println("for clip---")

	// i = d 的key(值)
	if i, d := m["A"]; d {
		fmt.Println(i, d)
	}
}
