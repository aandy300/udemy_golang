// 136. JSON Marshal https://golang.org/pkg/encoding/json/#example_Marshal
// 範例 JSON Marshal ¶ 將資料打包成JSON格式
// 結構體字頭要大寫 JSON印才正常 Name v name x
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {

	p1 := person{
		Name: "ming",
		Age:  29,
	}
	p2 := person{
		Name: "aaa",
		Age:  33,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs), "1111", bs)
}
