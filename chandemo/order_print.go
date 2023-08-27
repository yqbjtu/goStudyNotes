
package main


import (
	"fmt"
	"sync"
)
 
/**
题目 指定各种不同顺序执行 First(), Second(), Third() 三个 goroutine，但三者都必须以不变顺序印出字串，印出顺序不受顺序执行影响。

goroutine 若不刻意控制，将无法保证执行的先后顺序，因此本题就是要考核对 goroutine 顺序控制的能力
*/

func printFirst(wg *sync.WaitGroup, notifySecond chan int) {
	defer wg.Done()

	fmt.Println("first")
	// 只有first执行完毕，通过 chan notifySecond 通知printSecond中可以开始执行了
	close(notifySecond)
}

func printSecond(wg *sync.WaitGroup, notifySecond chan int, notifyThird chan int) {
	defer wg.Done()

	select {
	// printFirst 没有调用close(notifySecond)时，一直阻塞中
		case <-notifySecond:
			fmt.Println("second")
	}
	
	close(notifyThird)
}
 
func printThird(wg *sync.WaitGroup, notifyThird chan int) {
	defer wg.Done()

	select {
		// 在printSecond 没有调用close(notifyThird)时，一直阻塞中
		case <-notifyThird:
			fmt.Println("third")
	}

}
 
func main() {
	chanForSecond := make(chan int)
	chanForThird := make(chan int)

	wg := new(sync.WaitGroup)
    wg.Add(3)

	go printThird(wg, chanForThird)
	go printFirst(wg, chanForSecond)
	//go printThird(ch2)
	go printSecond(wg, chanForSecond, chanForThird)
 

	wg.Wait()
	fmt.Println("不管go routine调用顺序如何， 按first second third顺序执行完毕")
}
