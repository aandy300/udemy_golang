package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1) //Add buffer cause not deadlock. http://stackoverflow.com/questions/14050673/why-are-my-channels-deadlocking
	needRoutine := false

	if needRoutine {
		go func() {
			time.Sleep(time.Second * 3)
			c <- 3
		}()

	} else {
		c <- 0
	}

	totalResult := <-c*3 + 5

	fmt.Printf("result is %d", totalResult)
}
