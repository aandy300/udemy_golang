// 138. Writer interface 各種 寫與印 的方法(func)

// Encorder 144. main 和 看圖

// 扯到file 都會友 Write的type
//https://golang.org/pkg/os/#NewFile

// func WriteFile ¶ https://golang.org/pkg/io/ioutil/#WriteFile
// func (*File) Write ¶ https://golang.org/pkg/os/#File.Write
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello playword")
	fmt.Fprint(os.Stdout, "hello playword")
	io.WriteString(os.Stdout, "\nhello playword")
}
