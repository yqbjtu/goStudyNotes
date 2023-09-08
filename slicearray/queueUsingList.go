package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("mock Queue Using List")
	// queue FIFO 
	q2 := list.New()
	for i := 0; i < 5; i++ {
		fmt.Println("Enqueuing ", q2.PushBack(i).Value)
	}
	for q2.Len() > 0 {
		front := q2.Front()
		q2.Remove(front)
		fmt.Println("Dequeuing ", front.Value)
	}
}
