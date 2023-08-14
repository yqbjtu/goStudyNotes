package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"
)

func main() {
	var a interface{} = 10
	switch a.(type) {
	case int:
		fmt.Println("a is an int")
	case string:
		fmt.Println("a is a string")
	}

	var num1 = 10
	fmt.Printf("num1=%v, num1=%+v, num1=%#v\n", num1, num1)
	fmt.Println(reflect.TypeOf(num1)) // "int"

	var num2 = 10.00
	fmt.Printf("num2=%v, num2=%+v\n", num2, num2)
	fmt.Println(reflect.TypeOf(num2)) // "float64"

	var intStrValue = "10.5"
	value, err := strconv.Atoi(intStrValue)
	if err != nil {
		fmt.Printf("fail to convert string '%s' to int, err: %v\n", intStrValue, err)
	}
	fmt.Printf("Value %s of type %T to be converted to data type of %T the value is %02d. \n", intStrValue, intStrValue, value, value)

	intStrValue = "10"
	value, err = strconv.Atoi(intStrValue)
	if err != nil {
		fmt.Printf("fail to convert string '%s' to int, err: %v\n", intStrValue, err)
	}
	fmt.Printf("Value %s of type %T to be converted to data type of %T the value is %02d. \n", intStrValue, intStrValue, value, value)

	var name = "amit"
	fmt.Printf("name=%v, name=%+v\n", name, name)
	fmt.Println(reflect.TypeOf(name)) // "string"

	var w = os.Stdout
	fmt.Printf("w=%v, w=%+v\n", w, w)
	fmt.Println(reflect.TypeOf(w)) // "*os.File"

	start := time.Now()
	fmt.Printf("start=%v, start=%+v\n", start, start)
	fmt.Println(reflect.TypeOf(start)) // "time.Time"
}

//运行输出
//a is an int
//num1=10, num1=10, num1=%!v(MISSING)
//int
//num2=10, num2=10
//float64
//fail to convert string '10.5' to int, err: strconv.Atoi: parsing "10.5": invalid syntax
//Value 10.5 of type string to be converted to data type of int the value is 00.
//Value 10 of type string to be converted to data type of int the value is 10.
//name=amit, name=amit
//string
//w=&{0x140000ae060}, w=&{file:0x140000ae060}
//*os.File
//start=2023-08-14 21:09:53.653219 +0800 CST m=+0.000332085, start=2023-08-14 21:09:53.653219 +0800 CST m=+0.000332085
//time.Time
