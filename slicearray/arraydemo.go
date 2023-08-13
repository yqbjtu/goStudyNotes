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
	sliceB[2] = 5
	fmt.Printf("sliceA address:%p, value:%v, typeAndValue:%+v\n", &sliceA, sliceA, sliceA)
	fmt.Printf("sliceA address:%p, value:%v, typeAndValue:%+v\n", &sliceB, sliceB, sliceB)
}

//运行结果
//a address:0xc000040080, value:[beijing shanghai wuhan xian], typeAndValue:[beijing shanghai wuhan xian]
//a address:0xc0000400c0, value:[beijing shanghai nanjing xian], typeAndValue:[beijing shanghai nanjing xian]
//sliceA address:0xc0000044c0, value:[2 3 5], typeAndValue:[2 3 5]
//sliceA address:0xc0000044e0, value:[2 3 5], typeAndValue:[2 3 5]
