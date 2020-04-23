//107 function syntax 句法

// func (r 接收器) 標示符identifier(參數s)  (返回sreturn(s)) {...}
// func (r receivere) identifier(parameters) (return(s)) {...}

// 回傳 會自動回傳到觸發 function的容器
// 如果沒指定類型 回傳時可以改(?)

package main

import "fmt"

func main() {
	// 簡易版 function
	hi()
	// 帶入參數 function
	get_value_in_function("value")
	// 帶入參數 並 回傳 function 需要容器
	value_st := string("0000")
	fmt.Println(value_st)
	////////////////////////////////////////////////
	st := value_in_and_return_new_value_out(value_st)
	fmt.Println(st)
	//  雙value 帶入帶出
	x, y := two_value("A", "B")
	fmt.Println(x)
	fmt.Println(y)
}

// 簡易版 function 單純召喚並運作
func hi() {
	fmt.Println("hello word")
}

// 帶入參數 function 若要使用帶入的參數 需要容器 ex: x
func get_value_in_function(x string) {
	fmt.Println("i am function i get:", x)
}

// 帶入參數 並 回傳 function 需要容器
func value_in_and_return_new_value_out(st string) string {
	return fmt.Sprint("回傳到st", st)
}

// 雙參數帶入帶出 回傳時更改類型 or 別的東西
func two_value(x1 string, y2 string) (string, bool) {
	a := fmt.Sprint(x1, y2)
	b := false
	return a, b
}
