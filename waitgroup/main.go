package main

import (
  "fmt"
  "math/rand"
  "strings"
  "sync"
  "time"
)

func main() {

  var wg sync.WaitGroup
  fmt.Printf("%T\n", wg) 
  fmt.Println(wg) 

  wg.Add(3)
  rand.Seed(time.Now().UnixNano())

  go printNum(&wg, 1)
  go printNum(&wg, 2)
  go printNum(&wg, 3)

  wg.Wait() //进入阻塞
  defer fmt.Println("main over ...")
}

func printNum(wg *sync.WaitGroup, num int) {
  for i := 1; i < 3 : i++ {
     pre := strings.Repeat("\t", num -1) 
     fmt.Printf("%s %d goroutine , %d", pre, num, i)
     time.Sleep(time.Second)
  }
  
  wg.Done()
}
