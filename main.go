package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(letterCombinations("222"))
	end := time.Now()
	fmt.Println(end.Sub(start))
}
