//break continue
package main

import "fmt"

func main() {
	x := 1
	for {
		x++
		fmt.Println("break之上", x)
		if x > 30 {
			fmt.Println("break裡面", x)
			break
		}
		fmt.Println("continue之上")
		if x%2 != 0 {
			fmt.Println("continue裡面", x)
			continue
		}
		fmt.Println("IF之下", x)
	}
	fmt.Println("done.")

J:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 6 {
				break J //现在终止的是j 循环，而不是i的那个
			}
			fmt.Println(i)
		}
	}
}
