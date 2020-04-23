//105 struct 結構體練習 結構體中放入結構體 分類 寫入資料
package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type car_v1 struct {
	value_l bool
	vehicle
}
type car_v2 struct {
	value_h bool
	vehicle
}

func main() {

	x := car_v1{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		value_l: false,
	}
	y := car_v2{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		value_h: true,
	}

	fmt.Println(x.doors, x.color, x.value_l)
	fmt.Println(y.doors, y.color, y.value_h)
}
