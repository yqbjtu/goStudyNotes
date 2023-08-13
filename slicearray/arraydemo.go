package main

import "fmt"

func main() {
    a := [...]string{"beijing", "shanghai", "wuhan", "xian"}
	b := a //b 是a的一份copy
	b[2] = "najing"
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

//运行结果
//a: [beijing shanghai wuhan xian]
//b: [beijing shanghai najing xian]
