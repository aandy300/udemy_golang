//error
package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

func isEnable(enable bool) (bool, error) {
	if enable {
		return false, errors.New("You can't enable this setting")
	}

	return true, nil
}

var y int

func main() {

	//_ = 空的儲存空間 void 不打算再用得容器

	_, err := sprt(-10)
	// _, err :=
	if err != nil {
		fmt.Println(err)
	}

}

func sprt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("error herer")
	}
	return 42, nil

}

func foo(x int) {
	// sync.Locker
	fmt.Println(x)
}
