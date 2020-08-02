// 143. 解碼
// https://mholt.github.io/json-to-go/ JSON-to-Go
// https://mholt.github.io/json-to-go/ 解碼格式 type user struct{...} > JSON檔(str) > []byte() 轉換為uint8 > 宣 []type 切片儲存 > json.Unmarshal 解碼 > for range 印出

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var users []user

	err := json.Unmarshal(bs, &users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("data", users)

	for i, v := range users {
		fmt.Println("number", i)
		fmt.Println(v.Age, v.First)
		for _, d := range v.Sayings {
			fmt.Println("\t", d)
		}
	}
}
