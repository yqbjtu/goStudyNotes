package main

import (
	"fmt"
)

var (
	//实例化一个长度容量为5的slice
	//从slice切出一个slice2
	slice  = make([]int, 5)
	slice2 = slice[2:]
)

func main() {
	fmt.Printf("初始容量\n")
	printSlice()

	fmt.Printf("修改slice2 第0个元素\n")
	slice2[0] = 1
	printSlice()

	fmt.Printf("slice2 append 2个元素, 可以发现slice2 cap从之前的3 变为6\n")
	slice2 = append(slice2, 2)
	printSlice()

	fmt.Printf("slice2 修改slice2第0个元素\n")
	slice2[0] = 3
	printSlice()

	i := 0
	for i = 0; i < 1024; i++ {
		slice2 = append(slice2, 2)
	}
	fmt.Printf("slice2 append 1024个元素\n")
	printSliceLengthAndCapacity()

	slice2 = append(slice2, 5)
	fmt.Printf("slice2 append 1024后再app 5个元素\n")
	printSliceLengthAndCapacity()
}

func printSlice() {
	fmt.Printf("slice  len=%d cap=%d %v\n", len(slice), cap(slice), slice)
	fmt.Printf("slice2 len=%d cap=%d %v\n", len(slice2), cap(slice2), slice2)
}
func printSliceLengthAndCapacity() {
	fmt.Printf("slice  len=%d cap=%d\n", len(slice), cap(slice))
	fmt.Printf("slice2 len=%d cap=%d\n", len(slice2), cap(slice2))
}

//运行结果
//初始容量
//slice  len=5 cap=5 [0 0 0 0 0]
//slice2 len=3 cap=3 [0 0 0]
//修改slice2 第0个元素
//slice  len=5 cap=5 [0 0 1 0 0]
//slice2 len=3 cap=3 [1 0 0]
//slice2 append 2个元素, 可以发现slice2 cap从之前的3 变为6
//slice  len=5 cap=5 [0 0 1 0 0]
//slice2 len=4 cap=6 [1 0 0 2]
//slice2 修改slice2第0个元素
//slice  len=5 cap=5 [0 0 1 0 0]
//slice2 len=4 cap=6 [3 0 0 2]
//slice2 append 1024个元素
//slice  len=5 cap=5
//slice2 len=1028 cap=1184
//slice2 append 1024后再app 5个元素
//slice  len=5 cap=5
//slice2 len=1029 cap=1184
