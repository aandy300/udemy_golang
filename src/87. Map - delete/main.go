//map 更改 新增 檢查 字典
package main

import "fmt"

func main() {
	m := map[string]int{
		"A": 111,
		"B": 222,
		"C": 333,
	}
	fmt.Println(m)

	delete(m, "A")
	fmt.Println("刪除A之後:", m)
	fmt.Println("如果刪除或不存在,印出得職會為0:", m["A"])

	if v, ok := m["B"]; ok {
		fmt.Println("印出 值,存在還是不存在:", v, ok)
		delete(m, "B")
		delete(m, "C")
		fmt.Println("刪除所有map後:", m)
	}
}
