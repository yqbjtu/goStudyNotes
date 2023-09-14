package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	name, err := os.Hostname()
	if err == nil {
		fmt.Printf("hostname:%s\n", name)
	}

	fmt.Printf("cpu num:%d\n", runtime.GOMAXPROCS(0))
	fmt.Println("逻辑CPU数量: ", runtime.NumCPU())

	// Goexit方法会终止当前goroutine协程
	// 创建一个goroutine
	go func() {
		fmt.Println("子协程开始执行...")
		runtime.Goexit() //终止当前goroutine
		fmt.Println("子协程执行结束...")
	}()

	time.Sleep(3 * time.Second) //主函数休眠3秒，让子协程有充分的时间执行完
	fmt.Println("主函数执行完毕")
}
