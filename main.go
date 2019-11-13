package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	x := []int{}
	fmt.Println(removeDuplicates(x))
	fmt.Println(x)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
