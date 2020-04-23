//106 隱匿 struct 結構體 宣在main 裡面
//結構體內宣 map 和 array
package main

import "fmt"

func main() {

	x := struct {
		number int
		//struct use map
		friend map[string]int
		//struct use array
		name []string
	}{
		number: 1,

		//結構體內map用法 寫入資料
		friend: map[string]int{
			"111": 2,
			"222": 3,
		},
		//結構體內array用法 寫入資料
		name: []string{
			"andy",
			"ming",
		},
	}

	fmt.Println(x.friend)
	fmt.Println(x.name)
	fmt.Println(x.number)

}
