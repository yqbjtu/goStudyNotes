package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Go 1.19引入的atomic.Int32
func main() {
	var n atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			n.Add(1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(n.Load()) // 1000
}
