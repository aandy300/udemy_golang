// 183.Hands-on exercise#2
// fmt.Errorf() & errors.New()

// 圖?

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	// Fatalln() 內建 exit 所以不用 return
	if err != nil {
		log.Fatalln("err")
	}
	// 需要 return 返回 main 結束程式
	// if err != nil {
	// 	log.Println("err")
	// 	return
	// }
	fmt.Println(string(bs))
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("err:%v", err)
		// 下面 等於 上面的 背後在跑的內容 || 上面為簡化版
		// return bs, errors.New(fmt.Sprintf("ERR:%v", err))
	}
	return bs, nil
}
