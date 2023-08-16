package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watchWithContext(ctx, "routine1")
	go watchWithContext(ctx, "routine2")
	go watchWithContext(ctx, "routine3")

	time.Sleep(6 * time.Second)
	fmt.Println("通过 ctx cancel，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(3 * time.Second)
}

func watchWithContext(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，ctx cancel了...")
			return
		default:
			fmt.Println(name, "goroutine ctx监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

// 结论 使用context cancel 函数通知取消时，3 个 goroutine 都会被结束。
// Context 的控制能力，它就像一个控制器一样，按下开关后，所有基于这个 Context 或者衍生的子 Context 都会收到通知
//运行结果
//routine1 goroutine ctx监控中...
//routine3 goroutine ctx监控中...
//routine2 goroutine ctx监控中...
//routine2 goroutine ctx监控中...
//routine1 goroutine ctx监控中...
//routine3 goroutine ctx监控中...
//routine2 goroutine ctx监控中...
//routine3 goroutine ctx监控中...
//routine1 goroutine ctx监控中...
//通过 ctx cancel，通知监控停止
//routine3 监控退出，ctx cancel了...
//routine2 监控退出，ctx cancel了...
//routine1 监控退出，ctx cancel了...
