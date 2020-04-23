package main

//Append
//append [:] 定點
//append(z[1:], x[3:]...)
//z[1:] = 定點在z[]的最尾端 之後的append會加在這
//z[:3] = 定點在z[]index的2(:3結束值不包括) 之後的append會加在這

import "fmt"

func main() {

	x := []int{119, 229, 339, 449, 559, 669}
	fmt.Println(x)

	for i, d := range x {
		fmt.Printf("x類型:%T\tindex:%d\tvalue:%d\n", x, i, d)
	}

	z := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(z)
	z = append(z[:1], x[3:]...)
	fmt.Println(z)

}
