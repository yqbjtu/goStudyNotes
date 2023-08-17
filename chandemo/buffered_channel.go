package main

import (
	"fmt"
)

var data = []string{"Hello", "Go", "Linux"}

func sendData(ch chan string) {
	for _, elem := range data {
		fmt.Printf("向chan string放入%v \n", elem)
		ch <- elem
	}
	close(ch)
}

func main() {

	// creates capacity of 2
	ch := make(chan string, 2)
	sendData(ch)
	//go sendData(ch)
	for v := range ch {
		fmt.Printf("从chan string读取%v \n", v)
	}
}

/*运行结果， 不使用go sendData(ch)， chan的缓存长度是2 ，但是放入了3
go run buffered_channel.go
向chan string放入Hello
向chan string放入Go
向chan string放入Linux
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.sendData(0x60?)
/Users/ericyang/GolandProjects/typecast/buffered_channel.go:12 +0xb0
main.main()
/Users/ericyang/GolandProjects/typecast/buffered_channel.go:21 +0x30
exit status 2

行结果， go sendData(ch)
g向chan string放入Hello
向chan string放入Go
向chan string放入Linux
从chan string读取Hello
从chan string读取Go
从chan string读取Linux
*/
