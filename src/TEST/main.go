package main

import (
	"fmt"
	"runtime"
)

func say(s string) {

	for i := 0; i < 2; i++ {

		runtime.Gosched()

		fmt.Println("text printout:", s)

	}

}

func main() {

	go say("world")

	say("hello")

	fmt.Println("test")

}
