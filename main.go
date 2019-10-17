package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	end := time.Now()
	fmt.Println(end.Sub(start))
}
