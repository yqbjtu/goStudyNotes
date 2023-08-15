package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	// 通过chan 停止goroutine
	go watch(stop, "routine1")
	go watch(stop, "routine2")
	go watch(stop, "routine3")

	time.Sleep(5 * time.Second)
	fmt.Println("通过 chan <- 通知监控停止")
	stop <- true

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(2 * time.Second)
	fmt.Println("main done")
}

func watch(stop chan bool, name string) {
	for {
		select {
		case <-stop: // 收到了停滞信号
			fmt.Println(name, "监控退出，chan停止了...")
			return
		default:
			fmt.Println(name, "goroutine chan监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

//运行结果
//routine1 goroutine chan监控中...
//routine3 goroutine chan监控中...
//routine2 goroutine chan监控中...
//routine1 goroutine chan监控中...
//routine3 goroutine chan监控中...
//routine2 goroutine chan监控中...
//routine2 goroutine chan监控中...
//routine1 goroutine chan监控中...
//routine3 goroutine chan监控中...
//通过 chan <- 通知监控停止
//routine3 监控退出，chan停止了... (只有一个goroutine退出了)
//routine1 goroutine chan监控中...
//routine2 goroutine chan监控中...
//main done
//routine2 goroutine chan监控中...
