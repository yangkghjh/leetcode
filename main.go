package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(generateParenthesis(3))
	end := time.Now()
	fmt.Println(end.Sub(start))
}
