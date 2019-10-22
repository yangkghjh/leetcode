package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(isValid("["))
	end := time.Now()
	fmt.Println(end.Sub(start))
}
