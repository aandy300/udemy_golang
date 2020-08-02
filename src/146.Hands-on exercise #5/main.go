// 146. sort 格式化 排序 自訂 同  影片 145內建sort(int,string...) 146sort(自訂)
// type ByName []user
// 做 interface type  目標格式類型 > 下條件 len,swap,less(條件)

// sort後就寫入了直接更改記憶體位置資料 > sort.Strings(內建)改文字順序(say) > 照名稱(自訂type)排序 > 照年紀(自訂tpye)排序
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByName []user

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong and here's the error:", err)
	}

	fmt.Println("---------------by sort string")
	for i, v := range users {
		sort.Strings(v.Sayings)
		fmt.Println("num", i)
		fmt.Println(v.First, v.Last, v.Age)
		for _, s := range v.Sayings {
			fmt.Println("\t", s)
		}
	}

	sort.Sort(ByName(users))
	fmt.Println("---------------by sort name")
	for i, v := range users {
		fmt.Println("num", i, v.First, v.Last, v.Age)
		for _, s := range v.Sayings {
			fmt.Println("\t", s)
		}
	}

	sort.Sort(ByAge(users))
	fmt.Println("---------------by sort age")
	for i, v := range users {
		fmt.Println("num", i, v.First, v.Last, v.Age)
		for _, s := range v.Sayings {
			fmt.Println("\t", s)
		}
	}

}
