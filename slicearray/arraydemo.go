package main

import "fmt"

func main() {
    a := [...]string{"beijing", "shanghai", "wuhan", "xian"}
	b := a //b 是a的一份copy
	b[2] = "nanjing"
	fmt.Printf("a address:%p, value:%v, typeAndValue:%+v\n", &a, a, a)
	fmt.Printf("a address:%p, value:%v, typeAndValue:%+v\n", &b, b, b)

	sliceA := []int{2 , 3, 4}
	sliceB := sliceA //sliceB 是sliceA的引用

	sliceCopyFromA := make([]int, 20)
	//sliceCopyFromA 是sliceA的copy, 不是引用，从copy完成后，他们之间就没有关系
	copyCount := copy(sliceA, sliceCopyFromA)
	fmt.Printf("copy sliceA address:%p to sliceCopyFromA address:%p, copyCount:%d, type:%T\n",
	 &sliceA, sliceCopyFromA, copyCount, sliceCopyFromA)

	sliceB[2] = 5
	sliceCopyFromA[2] = 200
	fmt.Printf("sliceA address:%p, value:%v, typeAndValue:%+v\n", &sliceA, sliceA, sliceA)
	fmt.Printf("sliceB address:%p, value:%v, typeAndValue:%+v\n", &sliceB, sliceB, sliceB)
	fmt.Printf("sliceCopyFromA address:%p, value:%v, typeAndValue:%+v\n", &sliceCopyFromA, sliceCopyFromA, sliceCopyFromA)
}
//运行结果
//a address:0xc000040080, value:[beijing shanghai wuhan xian], typeAndValue:[beijing shanghai wuhan xian]
//a address:0xc0000400c0, value:[beijing shanghai nanjing xian], typeAndValue:[beijing shanghai nanjing xian] 
//copy sliceA address:0xc0000044c0 to sliceCopyFromA address:0xc000088000, copyCount:3, type:[]int
//sliceA address:0xc0000044c0, value:[0 0 5], typeAndValue:[0 0 5]
//sliceB address:0xc0000044e0, value:[0 0 5], typeAndValue:[0 0 5]
//sliceCopyFromA address:0xc000004500, value:[0 0 200 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], typeAndValue:[0 0 200 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
