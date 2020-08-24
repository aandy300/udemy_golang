// 185. Hands-on exercise#4
//  Error 與 181>07>05 類似 看圖
package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// e := fmt.Errorf("%v", f)
		// or
		e := errors.New("一二三")
		return 1, sqrtError{"0", "1", e}
	}
	return 42, nil
}
