//loop + if 印出100裡面可以2除後 商0的數字
package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
