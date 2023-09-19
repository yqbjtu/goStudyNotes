package main

import "fmt"

type Person struct {
	name    string
	address string
	age     int
}

func (p *Person) String() string {
	str := fmt.Sprintf("\n"+
		"[name=%s, "+
		"address=%s, "+
		"age=%d]\n", p.name, p.address, p.age)

	return str
}

func main() {
	tom := &Person{"Tom", "bj haidain xueyuanlu", 30}
	fmt.Println(tom)
	fmt.Printf("tom: %s \n", tom) // OK
	//fmt.Println("t : " + t) // Compiler error !!!
	fmt.Println("tom.String(): " + tom.String()) // OK if calling the function explicitly
}

/* 运行结果

[name=Tom, address=bj haidain xueyuanlu, age=30]

tom:
	[name=Tom, address=bj haidain xueyuanlu, age=30]

tom.String():
	[name=Tom, address=bj haidain xueyuanlu, age=30]
*/
