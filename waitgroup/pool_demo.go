package main

import (
	"fmt"
	"sync"
)

// pool 本身就是线程安全的，多个 goroutine 可以并发地存取对象
var pool *sync.Pool

type Person struct {
	Name string
}

func initPool() {
	// Pool.New 函数（用户自定义的）不一定是线程安全的。并且 Pool.New 函数可能会被并发调用，
	pool = &sync.Pool{
		New: func() interface{} {
			//return new(Person)
			return &Person{Name: "initName"} // 创建一个新的 Person 实例，并指定 Name 字段的值
		},
	}
}

func main() {
	initPool()

	p := pool.Get().(*Person)
	fmt.Printf("get one obj from pool:%v \n", p)
	p.Name = "first"

	// 需要在 Put 前，清除 p 的各个成员，这里为了展示，就不清除了
	// 放回对象池中以供其他goroutine调用
	fmt.Println("after chanage its name field to first, put it obj to pool:", p)
	pool.Put(p)
	pFirst := pool.Get().(*Person)
	fmt.Println("Get pFirst from pool:", pFirst)

	p = pool.Get().(*Person)
	p.Name = "second"
	pool.Put(p)
	fmt.Println("after chanage its name field to second, put it obj to pool:", p)

	fmt.Println("get from pool", pool.Get().(*Person))
	fmt.Println("get from pool", pool.Get().(*Person))
	fmt.Println("get from pool", pool.Get().(*Person))

	pool.Put(pFirst)
	fmt.Println("after pFrist is return, get from pool", pool.Get().(*Person))
}
//运行结果
//get one obj from pool:&{initName} 
//after chanage its name field to first, put it obj to pool: &{first}
//Get pFirst from pool: &{first}
//after chanage its name field to second, put it obj to pool: &{second}
//get from pool &{second}
//get from pool &{initName}
//get from pool &{initName}
//after pFrist is return, get from pool &{first}
