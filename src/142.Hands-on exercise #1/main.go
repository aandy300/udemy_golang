// 142. 打包練習 //注意 struct大寫
// https://gobyexample.com/json
// map[string]interface{}將保存字符串到任意數據類型的映射。
// string(bs) 因為 string有對應的映射所以可以解碼JSON?

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println("users容器存的資料", users)

	// your code goes here

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("打包前?", string(bs))
	fmt.Println("打包後(json.Marshal後)", bs)

}
