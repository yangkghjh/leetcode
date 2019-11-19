package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(divide(1, 2))
	end := time.Now()
	fmt.Println(end.Sub(start))
}
