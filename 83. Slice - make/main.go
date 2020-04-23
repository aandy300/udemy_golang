//Slice make 自行製作容器 非引用防止使用資源過大
package main

import "fmt"

func main() {

	//make([]類型,現在的長度,總容量)
	x := make([]int, 5, 100)

	//len() = 使用到的長度
	//cap() = 實際上的容量
	fmt.Println("現在的長度", len(x))
	fmt.Println("總容量", cap(x))
	fmt.Println(x)

	x = append(x, 1)
	fmt.Println("現在的長度", len(x))
	fmt.Println("總容量", cap(x))
	fmt.Println(x)
}
