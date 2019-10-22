package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100))
	end := time.Now()
	fmt.Println(end.Sub(start))
}
