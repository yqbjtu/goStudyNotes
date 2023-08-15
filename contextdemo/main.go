package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	// 通过chan 停止goroutine
	go func() {
		for {
			select {
			case <-stop: // 收到了停滞信号
				fmt.Println("监控退出，chan停止了...")
				return
			default:
				fmt.Println("goroutine chan监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	// 通过context 停止goroutine
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，ctx停止了...")
				return
			default:
				fmt.Println("goroutine ctx监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("通过 chan <- 通知监控停止")
	stop <- true

	fmt.Println("通过 ctx cancel，通知监控停止")
	cancel()

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(4 * time.Second)
	fmt.Println("main done")
}

//运行结果
//goroutine chan监控中...
//goroutine ctx监控中...
//goroutine chan监控中...
//goroutine ctx监控中...
//goroutine chan监控中...
//goroutine ctx监控中...
//通过 chan <- 通知监控停止
//监控退出，chan停止了...
//goroutine ctx监控中...
//通过 ctx cancel，通知监控停止
//监控退出，ctx停止了...
//main done
