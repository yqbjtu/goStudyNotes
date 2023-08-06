 package main
 
 
 import (
 "fmt"
 )
 
 func main() {
    ch1 := make(chan string)
    go sendData(ch1)
	
	  for value := range ch1 {
	      fmt.Println("read data from ch1:", value)
	  }
 }
 
 func sendData(cha1 chan string) 
 {   
     //循环从channel中接收数据，需要陪着关闭channel 以及for range读取数据
     defer close(ch1)
	   for i :=0; i<3; i++ {
	       ch1 <- fmt.Sprintf("send data%d \n", i)
	   }
	 
	   fmt.Println("send data to ch1 done")
 }
