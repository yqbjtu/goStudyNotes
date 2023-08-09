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

	stringOne := ""
	stringTwo := "LastOfUs"
	stringThree := "   "
	stringFour := " \t\n  "

	if stringOne == "" {
		fmt.Printf("stringOne is empty. it is '%s'.\n", stringOne)
	}

	if stringTwo == "" {
		fmt.Printf("stringTwo is empty. it is '%s'.\n", stringTwo)
	} else {
		fmt.Printf("stringTwo is not empty. it is '%s'.\n", stringTwo)
	}

	if strings.TrimSpace(stringThree) == "" {
		fmt.Printf("stringThree is empty. it is '%s'.\n", stringThree)
	} else {
		fmt.Printf("stringThree is not empty. it is '%s'.\n", stringThree)
	}

	if strings.TrimSpace(stringFour) == "" {
		fmt.Printf("stringFour is empty. it is '%s'.\n", stringFour)
	} else {
		fmt.Printf("stringFour is not empty. it is '%s'.\n", stringFour)
	}
}

//运行结果
//原始消息:ZhangSan is a programmer.，通过空格split，使用| join结果:ZhangSan|is|a|programmer.
//原始消息:ZhangSan is a programmer.，通过中划线-split， 使用下划线'_'join结果:ZhangSan is a programmer.
//原始消息:ZhangSan-is-a-programmer.，通过中划线-split，使用下划线'_'join结果:ZhangSan_is_a_programmer.
//stringOne is empty. it is ''.
//stringTwo is not empty. it is 'LastOfUs'.
//stringThree is empty. it is '   '.
//stringFour is empty. it is '    
//'.
