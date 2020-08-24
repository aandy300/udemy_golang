// 182. Hands-on exercise#1 log.Fatalln("ERR:", err)

// sample 因 沒有任何錯誤 所以不會動作 後續練習基本模型

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

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("ERR:", err)
	}
	fmt.Println(string(bs))

}
