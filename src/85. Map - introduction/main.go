//map 字典 key
package main

import "fmt"

func main() {
	m := map[string]int{
		"an":   11,
		"aaan": 33,
	}
	fmt.Println(m)
	fmt.Println(m["an"])

	v, k := m["an"]
	fmt.Println(v, k)

	m["an"] = 55
	fmt.Println(m["an"])
}
