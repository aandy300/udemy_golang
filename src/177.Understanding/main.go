// 177.Understanding err 解說 圖

// https://godoc.org/errors

// https://blog.golang.org/error-handling-and-go

// type error interface {
//     Error() string
// }

// https://godoc.org/errors#New

// // New returns an error that formats as the given text.
// // Each call to New returns a distinct error value even if the text is identical.
// func New(text string) error {
// 	return &errorString{text}
// }

package main

func main() {
	// f, err := os.Open("filename.ext")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// do something with the open *File f
}
