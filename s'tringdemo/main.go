package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "ZhangSan is a programmer."
	list1 := strings.Split(msg, " ")
	msgJoin := strings.Join(list1, "|")
	fmt.Printf("原始消息:%s，通过空格split，使用| join结果:%s\n", msg, msgJoin)

	list2 := strings.Split(msg, "-")
	msgJoin = strings.Join(list2, "_")
	fmt.Printf("原始消息:%s，通过中划线-split， 使用下划线'_'join结果:%s\n", msg, msgJoin)

	msg2 := "ZhangSan-is-a-programmer."
	list2 = strings.Split(msg2, "-")
	msgJoin2 := strings.Join(list2, "_")
	fmt.Printf("原始消息:%s，通过中划线-split，使用下划线'_'join结果:%s\n", msg2, msgJoin2)
}

//运行结果
//原始消息:ZhangSan is a programmer.，通过空格split，使用| join结果:ZhangSan|is|a|programmer.
//原始消息:ZhangSan is a programmer.，通过中划线-split， 使用下划线'_'join结果:ZhangSan is a programmer.
//原始消息:ZhangSan-is-a-programmer.，通过中划线-split，使用下划线'_'join结果:ZhangSan_is_a_programmer.
