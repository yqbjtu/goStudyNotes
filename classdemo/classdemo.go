package main

import (
	"fmt"
)


type Person struct {
	name, phone string
	age int
}

type Student struct {
	Person //匿名字段
	school string
}

type Employee struct {
	Person //匿名字段
	company string
}

func (p *Person) introduce() {
    fmt.Printf("大家好， 我是%s, 今年%d岁， 我的联系方式是：%s \n", p.name, p.age,  p.phone)
}

func main() {
	stu1 := Student{Person{"lilei", "13266885522", 12}, "高新一中"}
	emp1 := Employee{Person{"张三", "13133551234", 37}, "大厂"}

	stu1.introduce()
	emp1.introduce()
}

// 运行结果
//大家好， 我是lilei, 今年12岁， 我的联系方式是：13266885522
//大家好， 我是张三, 今年37岁， 我的联系方式是：13133551234
