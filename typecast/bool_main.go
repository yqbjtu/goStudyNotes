package main

import (
	"fmt"
	"strconv"
)

// 可以同以下三种方式将bool转为string
// strconv 包的func FormatBool(b bool) string， 返回 "true" or "false"
// fmt.Sprintf(string, bool)使用格式化 "%t" or "%v"
// func ParseBool(str string) (bool, error) 将字符转为bool

func main() {
	v := true
	// bool to string
	str1 := strconv.FormatBool(v)
	fmt.Printf("原始值       %v %T\n", v, v)
	fmt.Printf("转换为string %v %T\n", str1, str1)
	fmt.Println()

	boolVal := false
	// convert bool to string value
	str2 := fmt.Sprintf("%v", boolVal)
	fmt.Printf("原始值       %v %T\n", boolVal, boolVal)
	fmt.Printf("转换为string %v %T\n", str2, str2)
	fmt.Println()

	//接受 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False 等字符串；
	//其他形式的字符串会返回错误
	strB1 := "1"
	b1, _ := strconv.ParseBool(strB1)
	fmt.Printf("原始值  %v %T\n", strB1, strB1)
	fmt.Printf("转换为  %v %T\n", b1, b1)
	fmt.Println()

	strB2 := "false"
	b2, _ := strconv.ParseBool(strB2)
	fmt.Printf("原始值  %v %T\n", strB2, strB2)
	fmt.Printf("转换为  %v %T\n", b2, b2)
	fmt.Println()

	strB3 := "fal1xse"
	b3, err := strconv.ParseBool(strB3)
	fmt.Printf("原始值  %v %T\n", strB3, strB3)
	if err != nil {
		fmt.Printf("fail to convert %v to bool, errMsg:%v \n", strB3, err)
	} else {
		fmt.Printf("转换为  %v %T\n", b3, b3)
	}
}

//运行输出
//原始值       true bool
//转换为string true string
//
//原始值       false bool
//转换为string false string
//
//原始值  1 string
//转换为  true bool
//
//原始值  false string
//转换为  false bool
//
//原始值  fal1xse string
//fail to convert fal1xse to bool, errMsg:strconv.ParseBool: parsing "fal1xse": invalid syntax
