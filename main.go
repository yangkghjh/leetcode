package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, 1))
	end := time.Now()
	fmt.Println(end.Sub(start))
}
