package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(mergeTwoLists(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}, &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}))
	end := time.Now()
	fmt.Println(end.Sub(start))
}
