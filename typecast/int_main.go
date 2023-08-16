package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 可以同以下三种方式将int转为string
// Sprintf() from fmt
// Itoa() from strconv
// FormatInt() from strconv

// 通过如下将int转为sting
//strconv.Atoi
//fmt.Sscan
//strconv.ParseInt
//string()
//strconv.Itoa
//strconv.FormatInt

func main() {
	var res1 int64 = 3
	str1 := fmt.Sprintf("The value is %d of type %T, converted to a string is %v", res1, res1, res1)
	println(str1)

	str1Int := "111"
	fmt.Printf("%v %T\n", str1Int, str1Int) //111 string

	// Atoi 等价于 ParseInt(s, 10, 0),
	valInt, _ := strconv.Atoi(str1Int)
	fmt.Printf("%v %T\n", valInt, valInt) // 111 int

	var res2 = 20
	fmt.Println(res2, reflect.TypeOf(res2))
	// Itoa 等同于 FormatInt(int64(i), 10).， 也就是按照10进制将int转为string
	str2 := strconv.Itoa(res2)
	fmt.Printf("Value %d 通过Itoa函数转为string：", res2)
	fmt.Println(str2, reflect.TypeOf(str2))

	var res3 int64 = 21
	fmt.Println(res3, reflect.TypeOf(res3))

	st3Binary := strconv.FormatInt(res3, 2)
	fmt.Printf("Value %d 按照2进制转为string：", res3)
	fmt.Println(st3Binary, reflect.TypeOf(st3Binary))

	st3Decimal := strconv.FormatInt(res3, 10)
	fmt.Printf("Value %d 按照10进制转为string：", res3)
	fmt.Println(st3Decimal, reflect.TypeOf(st3Decimal))

	st3Hexadecimal := strconv.FormatInt(res3, 16)
	fmt.Printf("Value %d 按照16进制转为string：", res3)
	fmt.Println(st3Hexadecimal, reflect.TypeOf(st3Hexadecimal))

	str3Int := "100"
	flex, _ := strconv.ParseInt(str3Int, 10, 0)
	fmt.Printf("按照10进制 %v %T\n", flex, flex)

	flex, _ = strconv.ParseInt(str3Int, 16, 0)
	fmt.Printf("按照16进制 %v %T\n", flex, flex)
}

//运行输出
//The value is 3 of type int64, converted to a string is 3
//111 string
//111 int
//20 int
//Value 20 通过Itoa函数转为string：20 string
//21 int64
//Value 21 按照2进制转为string：10101 string
//Value 21 按照10进制转为string：21 string
//Value 21 按照16进制转为string：15 string
//按照10进制 100 int64
//按照16进制 256 int64
