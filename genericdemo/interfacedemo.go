package main

import "fmt"

type MyInterface[T int|string] interface {
	~int|float64
	WriteOne(data T) T
	ReadOne() T
}
type Note int

func (n Note) WriteOne(one string) string {
	return one
}

func (n Note) ReadOne() string {
	return "small"
}

// 泛型接口，只能被当做类型参数来使用，无法被实例化。
func add[T MyInterface[string]] (a, b T) T {
	return a + b
}

func main() {
	var a Note = 1
	var b Note = 2
	res := add(a, b)
	fmt.Println(res)
	res1 := res.WriteOne("123")
	fmt.Println(res1)
}

//运行结果
// 3
// 123
