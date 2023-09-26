package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sigChan := make(chan os.Signal, 1)

	// https://man7.org/linux/man-pages/man1/killall.1.html
	// SIGTERM gracefully kills the process whereas SIGKILL kills the process immediately. SIGTERM signal can be handled,
	// ignored and blocked but SIGKILL cannot be handled or blocked
	// https://en.wikipedia.org/wiki/Signal_(IPC)#SIGKILL
	/*SIGTERM
	"Signal terminate"
	The SIGTERM signal is sent to a process to request its termination. Unlike the SIGKILL signal, it can be caught and interpreted or ignored by the process. This allows the process to perform nice termination releasing resources and saving state if appropriate. SIGINT is nearly identical to SIGTERM.
	*/
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	/*
		SIGKILL
		"Signal kill"
		The SIGKILL signal is sent to a process to cause it to terminate immediately (kill). In contrast to SIGTERM and SIGINT, this signal cannot be caught or ignored, and the receiving process cannot perform any clean-up upon receiving this signal. The following exceptions apply:
		Zombie processes cannot be killed since they are already dead and waiting for their parent processes to reap them.
		Processes that are in the blocked state will not die until they wake up again.
		The init process is special: It does not get signals that it does not want to handle, and thus it can ignore SIGKILL.[13] An exception from this rule is while init is ptraced on Linux.[14][15]
		An uninterruptibly sleeping process may not terminate (and free its resources) even when sent SIGKILL. This is one of the few cases in which a UNIX system may have to be rebooted to solve a temporary software problem.
		SIGKILL is used as a last resort when terminating processes in most system shutdown procedures if it does not voluntarily exit in response to SIGTERM. To speed the computer shutdown procedure, Mac OS X 10.6, aka Snow Leopard, will send SIGKILL to applications that have marked themselves "clean" resulting in faster shutdown times with, presumably, no ill effects.[16] The command killall -9 has a similar, while dangerous effect, when executed e.g. in Linux; it does not let programs save unsaved data. It has other options, and with none, uses the safer SIGTERM signal.
	*/
	done := make(chan bool, 1)

	go func() {
		select {
		case value := <-sigChan:
			//如果是ctrl+c， 不中断程序，只有kill -15 program 才退出， kill -9捕获不到直接退出
			if value == syscall.SIGTERM {
				fmt.Println("Received expected value:", value)
				// 等待时间太短，就看不到打印Received expected value: terminated
				time.Sleep(20 * time.Second)
				done <- true
			} else {
				fmt.Printf("Received unexpected value:%v, continue to read\n", value)
			}
		}

	}()

	pid := os.Getpid()

	fmt.Printf("The PID of the caller process is %d.\n", pid)
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

/*运行结果
The PID of the caller process is 25739.
awaiting signal
^CReceived unexpected value:interrupt, continue to read （ctrl+c  的结果）
Received expected value: terminated ，然后退出 （kill pid结果）
signal: killed  (kill -9 的结果)
*/
