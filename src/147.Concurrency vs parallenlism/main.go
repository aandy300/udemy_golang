// 147.  Concurrency Is Not Parallelism 並行 非 併發 解說?

// 解說
// https://medium.com/mr-efacani-teatime/concurrency%E8%88%87parallelism%E7%9A%84%E4%B8%8D%E5%90%8C%E4%B9%8B%E8%99%95-1b212a020e30

// Concurrency：相同的工作集合，一起完成同一份工作，互相合作，做團稽
// Parallelism：不同的工作集合，各自完成自己的工作，不互相干擾，有各自的考績
package main

import "fmt"

func main() {
	fmt.Println(".")
}
