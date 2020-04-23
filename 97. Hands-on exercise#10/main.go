//map 字典 for loop 印出所有資料 加入資料 刪除資料
package main

import "fmt"

func main() {

	// m := map[string][]string{}
	m := map[string][]string{
		`A`: []string{`111`, `222`, `777`},
		`B`: []string{`333`, `444`, `888`},
		`C`: []string{`555`, `666`, `999`},
	}
	fmt.Println(m)

	m[`D`] = []string{`000`, `111`, `000`}

	delete(m, `A`)

	for i, v := range m {
		fmt.Println(i, v)
		for i2, v2 := range v {
			// fmt.Printf()
			// fmt.Printf("這是m map裡的: %v\t 裡面的第 %v 筆資料\t 值為 %v\n", i, i2, v2)
			// fmt.Println()
			fmt.Println("這是m裡的:", i, "裡的第", i2, "資料", "值為", v2)
		}
	}
}
