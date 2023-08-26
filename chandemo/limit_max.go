package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 通过channel控制最大并发数
const MAX_CONCURRENT_JOBS = 2

func main() {

	waitChan := make(chan struct{}, MAX_CONCURRENT_JOBS)
	count := 0
	for {
        // 带缓冲区的chan，可以暂存数据，如果缓冲区满了，发生阻塞
		count++
        fmt.Println(count,  " will be sent to chan")
		waitChan <- struct{}{}
		go func(count int) {
			job(count)
			<-waitChan
		}(count)
	}
}

func job(index int) {
	fmt.Println(index, "begin doing something")
	time.Sleep(time.Duration(rand.Intn(10) * int(time.Second)))
	fmt.Println(index, "done")
}

//结果
//1  will be sent to chan
//2  will be sent to chan
//3  will be sent to chan   这里阻塞了
//1 begin doing something
//2 begin doing something
//1 done   等待1完成后， chan可以再次写数据
//4  will be sent to chan
//3 begin doing something
//2 done
//5  will be sent to chan
//4 begin doing something
//3 done
//6  will be sent to chan
//5 begin doing something
//5 done
//7  will be sent to chan
//6 begin doing something
//4 done
//8  will be sent to chan
//7 begin doing something
//exit status 2
