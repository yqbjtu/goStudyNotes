package main

import "fmt"

func testReturnWithoutName() int {
	i := 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2, i:", i)
	}()
	return i
}

func testReturnWithName() (i int) {
	i = 0
	defer func() {
		i += 1
		fmt.Println("defer3, i:", i)
	}()
	return i
}
func main() {
	fmt.Println("有名返回值的函数")
	fmt.Println("return", testReturnWithoutName())

	fmt.Println("有名返回值的函数")
	fmt.Println("return", testReturnWithName())
}

//执行结果
//有名返回值的函数
//defer2, i: 1
//defer1
//return 0
//有名返回值的函数
//defer3, i: 1
//return 1
