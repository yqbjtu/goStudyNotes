package main

import "fmt"

func main() {

	// 定义一个切片, 数组声明需要指定元素类型及元素个数. 而切片不用
	var slice1 []string
	fmt.Printf("未make, len=%d, cap=%d, (slice1 == nil):%t, slice1=%v\n", len(slice1), cap(slice1), slice1 == nil, slice1)

	slice1 = make([]string, 3)
	fmt.Printf("make后，len=%d, cap=%d, (slice1 == nil):%t, slice1=%v\n", len(slice1), cap(slice1), slice1 == nil, slice1)

	var intSlice1 = make([]int, 3, 5)
	fmt.Printf("len=%d, cap=%d, intSlice1=%v\n", len(intSlice1), cap(intSlice1), intSlice1)

	slice2 := make([]string, 3, 5)
	//也可以指定容量，其中 capacity 为可选参数。
	//make([]T, length, capacity)
	fmt.Printf("len=%d, cap=%d, slice2=%v\n", len(slice2), cap(slice2), slice2)

	slice1[0] = "a"
	slice1[1] = "b"
	slice1[2] = "c"
	fmt.Println("set:", slice1)
	fmt.Println("get[2]:", slice1[2])

	fmt.Println("slice1 len:", len(slice1))

	slice1 = append(slice1, "d")
	slice1 = append(slice1, "e", "f", "g", "a")
	fmt.Printf("slice1 append more item, len=%d, cap=%d, slice2=%v\n", len(slice1), cap(slice1), slice1)

	c := make([]string, len(slice1))
	copy(c, slice1)
	fmt.Println("copy slice1 to c:", c)

	l := slice1[2:5]
	fmt.Println("截取 slice1[2:5]内容 sl1:", l)

	l = slice1[:5]
	fmt.Println("截取 slice1[:5]内容 sl2:", l)

	l = slice1[2:]
	fmt.Println("截取 slice1[2:]内容 sl3:", l)

	initSlice3 := []string{"g", "h", "i"}
	fmt.Printf("initSlice3, len=%d, cap=%d, initSlice3=%v\n", len(initSlice3), cap(initSlice3), initSlice3)
	initSlice3 = append(initSlice3, "d")
	fmt.Printf("initSlice3 append one item, len=%d, cap=%d, slice2=%v\n", len(initSlice3), cap(initSlice3), initSlice3)

	twoDimensionSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDimensionSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDimensionSlice[i][j] = i + j
		}
	}

	fmt.Printf("twoDimensionSlice, len=%d, cap=%d, slice2=%v\n", len(twoDimensionSlice), cap(twoDimensionSlice), twoDimensionSlice)
}

//运行结果
////未make, len=0, cap=0, (slice1 == nil):true, slice1=[]
////make后，len=3, cap=3, (slice1 == nil):false, slice1=[  ]
////len=3, cap=5, intSlice1=[0 0 0]
////len=3, cap=5, slice2=[  ]
////set: [a b c]
////get[2]: c
////slice1 len: 3
////slice1 append more item, len=8, cap=12, slice2=[a b c d e f g a]
////copy slice1 to c: [a b c d e f g a]
////截取 slice1[2:5]内容 sl1: [c d e]
////截取 slice1[:5]内容 sl2: [a b c d e]
////截取 slice1[2:]内容 sl3: [c d e f g a]
////initSlice3, len=3, cap=3, initSlice3=[g h i]
////initSlice3 append one item, len=4, cap=6, slice2=[g h i d]
////twoDimensionSlice, len=3, cap=3, slice2=[[0] [1 2] [2 3 4]]
