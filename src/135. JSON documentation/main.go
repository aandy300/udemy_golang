// 135. JSON documentation https://golang.org/pkg/encoding/json/#example_Marshal
// 範例 JSON Marshal ¶ // Unmarshal ¶
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

// Unmarshal ¶

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	var jsonBlob = []byte(`[
// 	{"Name": "Platypus", "Order": "Monotremata"},
// 	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
// ]`)
// 	type Animal struct {
// 		Name  string
// 		Order string
// 	}
// 	var animals []Animal
// 	err := json.Unmarshal(jsonBlob, &animals)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	fmt.Printf("%+v", animals)
// }
