package main
//     基于lock
//     基于init函数
//     基于sync.Once  这里选择基于once

 import "sync"
 type singleton struct {
     // 单例对象的状态
 }
 var (
     instance *singleton
     once     sync.Once
 )
 func GetInstance() *singleton {
     once.Do(func() {
         instance = &singleton{}
         // 初始化单例对象的状态
     })
     return instance
 }

func main() {

    for i := 0; i < 30; i++ {
        go GetInstance()
    }
    // 简单起见时候使用sleep， 正式场合应该waitgroup 或者channel 
    time.Sleep(1000 * time.Millisecond)
}
