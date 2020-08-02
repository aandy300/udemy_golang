// 140. Sort custom 自訂分類
// https://golang.org/pkg/sort/#example_
//自訂type > type interface interface{..} > sort.Sort(ByAge(people))
//         > sort.Sort 呼叫 data.len .less .swap >執行完後更改記憶體位置資料
// 箱中箱 byage > people > person
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

//https://golang.org/pkg/sort/#Interface 下面三行func官方文件
//len長度 return len(a) swap更換位置x to y | y to x less誰先誰後算式
func (a ByAge) Len() int           { return len(a) }              // 長度
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }    // 更換位置
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age } // 誰先誰後 條件

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	//https://golang.org/pkg/sort/#Sort
	//func Sort(data Interface)
	//sort.Sort = call data.len less swap
	sort.Sort(ByAge(people))
	fmt.Println(people)

}
