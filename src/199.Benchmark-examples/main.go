// 199.Benchmark examples 效能比較
// CMD > go test -bench .

// Benchmark Sample

// func BenchmarkJoin(b *testing.B) {
// 	xs = strings.Split(s, " ")
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		Join(xs)
// 	}
// }

package main

import (
	"fmt"
	"strings"
	"udemy_golang/src/199.Benchmark-examples/mystr"
)

const s = "We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't feel insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us. It's not just in some of us; it's in everyone. And as we let our own light shine, we unconsciously give other people permission to do the same. As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson"

func main() {

	xs := strings.Split(s, " ")

	for _, v := range xs[1:] {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))
}
