//陣列中陣列 迴圈印出陣列中陣列
package main

import "fmt"

func main() {

	x := []string{"James", "Bond", "Shaken, not stirred"}
	z := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	y := [][]string{x, z}
	fmt.Println(y)

	for i, d := range y {
		fmt.Println("y[] index :", i)
		for i2, d2 := range d {
			// fmt.Println("現在在y[]的index:", i, "i2_index:", i2, "d2_:", d2)
			fmt.Printf("index: %v\t value:%v\n", i2, d2)
		}
	}
}
