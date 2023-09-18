package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {

    sigs := make(chan os.Signal, 1)

    // https://man7.org/linux/man-pages/man1/killall.1.html
    // SIGTERM gracefully kills the process whereas SIGKILL kills the process immediately. SIGTERM signal can be handled, 
    // ignored and blocked but SIGKILL cannot be handled or blocked
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    done := make(chan bool, 1)

    go func() {

        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
