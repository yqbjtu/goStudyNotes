package main

import (
	"fmt"
	"time"
)

func main() {
	// select 随机挑选一个可通信的case执行， 向chan2和chan4发送数据是可以执行
	chanCanRead()
	fmt.Println("----")

	chanTimeout()
	fmt.Println("----")

	chanIsFull()
	fmt.Println("----")

	var ch1, ch2, ch3 chan int
	var i1, i2 int

	// select 随机挑选一个可通信的case执行，如果所有case都没有数据到达，则执行default， 如果没有default，则阻塞
	select {
	case i1 = <-ch1:
		fmt.Println("receive data from ch1:", i1)
	case ch2 <- i2:
		fmt.Println("send data to ch2:", i2)

	case i3, ok := <-ch3:
		if ok {
			fmt.Println("receive data from ch3:", i3)
		} else {
			fmt.Println("ch3 is closed")
		}
		/*如果注释default，并且上的三个channel都没有数据到达， 就会出现如下情况
		fatal error: all goroutines are asleep - deadlock!
		goroutine 1 [select]:
		*/
	default:
		fmt.Println("no chan， execute default")
	}

}

func chanCanRead() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	ch4 := make(chan int, 1)

	var i1, i2 int

	select {
	case i1 = <-ch1:
		fmt.Println("receive data from ch1:", i1)
	case ch2 <- i2:
		fmt.Println("send data to ch2:", i2)
	case i3, ok := <-ch3:
		if ok {
			fmt.Println("receive data from ch3:", i3)
		} else {
			fmt.Println("ch3 is closed")
		}
	case ch4 <- 271828:
		fmt.Println("send data to ch2:", 271828)
	default:
		fmt.Println("no chan， execute default")
	}
}

func chanTimeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- "some value is written"

	}()

	select {
	case rs := <-ch:
		fmt.Println("result:", rs)

	case <-time.After(time.Second * 1):
		fmt.Println("timeout!")
	}
}

func chanIsFull() {
	ch := make(chan int, 1)

	ch <- 2

	select {
	case ch <- 3:
		fmt.Println("value in chan:", <-ch)
		fmt.Println("channel vaule is:", <-ch)

	default:
		//如果ch := make(chan int, 2)， 向chan写入就不会阻塞，因为ch <- 2只写入一个
		fmt.Println(" to send to chan is blocked, chan is full")
	}
}
/*运行结果
send data to ch2: 0
----
timeout!
----
 to send to chan is blocked, chan is full
----
no chan，  execute default
*/
