// 137. JSON unmarshal JSON 格式 還原 tag

// JSON檔(str) > []byte() 轉換為uint8 > 宣 []type 切片儲存 > json.Unmarshal 解碼 > for range 印出

// js to go https://mholt.github.io/json-to-go/
//			解包        資料 結構類型用address連結
// err := json.Unmarshal(bs, &people)

// 									pointed
// func Unmarshal(data []byte, v interface{}) error

// https://golang.org/pkg/encoding/json/#example_Unmarshal
// // Field appears in JSON as key "myName".
// Field int `json:"myName"`

// ---------------------------
// // Field appears in JSON as key "myName" and
// // the field is omitted from the object if its value is empty,
// // as defined above.
// Field int `json:"myName,omitempty"`

// // Field appears in JSON as key "Field" (the default), but
// // the field is skipped if empty.
// // Note the leading comma.
// Field int `json:",omitempty"`

// // Field is ignored by this package.
// Field int `json:"-"`

// // Field appears in JSON as key "-".
// Field int `json:"-,"`
// ---------------------------

// ---------------------------
// 字段在JSON中顯示為鍵“ myName”。
// 字段int`json：“ myName”`

// //字段在JSON中顯示為鍵“ myName”，並且
// //如果該字段的值為空，則從該對像中省略該字段，
// //如上定義。
// 字段int`json：“ myName，omitempty”`

// //字段在JSON中顯示為鍵“字段”（默認），但是
// //如果為空，則跳過該字段。
// //注意前導逗號。
// 字段int`json：“，omitempty”`

// //此包將忽略字段。
// 字段int`json：“-”`

// //字段在JSON中顯示為鍵“-”。
// 字段int`json：“-，”
// ---------------------------

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	//			tag
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {

	s := `[{"Name":"ming","Age":29},{"Name":"aaa","Age":33}]`

	// 將s轉換為uint8
	// 									pointed
	// func Unmarshal(data []byte, v interface{}) error
	// type byte = uint8W
	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//宣告 people有person的type
	//兩者都可以 people := []person{} or var people []person
	var people []person
	//			解包     資料 結構類型用address連結
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all data", people)

	for i, v := range people {
		fmt.Println("\nperson number", i)
		fmt.Println(v.Name, v.Age)
	}

}
