ackage main

import (
	"fmt"
	"sync"
)

// 程序交替打印字符和数字

func main() {
	chChan := make(chan struct{}, 1)
	numChan := make(chan struct{}, 1)
	var wg sync.WaitGroup

	// 控制了先打印数字，因为先给numChan 进行<-
	numChan <- struct{}{}
	wg.Add(1)
	go func(chChan, numChan chan struct{}) {
		defer wg.Done()
		for i := 'A'; i <= 'Z'; i++ {
			<-chChan
			fmt.Print(string(i))
			numChan <- struct{}{}
		}
	}(chChan, numChan)

	wg.Add(1)
	go func(chChan, numChan chan struct{}) {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			<-numChan
			fmt.Print(i)
			chChan <- struct{}{}
		}
	}(chChan, numChan)

	wg.Wait()
}

