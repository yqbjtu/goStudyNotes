package main

import (
	"fmt"
	"reflect"

	//"slices" go 1.21 has this
	"golang.org/x/exp/slices"
)

type Person struct {
	Name string
	Age  int
}

type PersonV2 struct {
	Name     string
	Age      int
	Interest []string
}

func Compare1(p1V2, p2V2 *PersonV2) bool {
	if !reflect.DeepEqual(p1V2.Name, p2V2.Name) {
		return false
	}

	if !reflect.DeepEqual(p1V2.Age, p2V2.Age) {
		return false
	}

	if len(p1V2.Interest) != len(p2V2.Interest) {
		return false
	}

	for _, v := range p1V2.Interest {
		if !slices.Contains(p2V2.Interest, v) {
			return false
		}
	}
}

func main() {
	p1 := Person{Name: "John", Age: 30}
	p2 := Person{Name: "John", Age: 30}

	if reflect.DeepEqual(p1, p2) {
		fmt.Println("p1 and p2 are equal")
	} else {
		fmt.Println("p1 and p2 are not equal")
	}

	s1 := []string{"foo", "bar", "baz"}
	s2 := []string{"foo", "bar", "baz"}

	if reflect.DeepEqual(s1, s2) {
		fmt.Println("s1 and s2 are equal")
	} else {
		fmt.Println("s1 and s2 are not equal")
	}

	p1V2 := PersonV2{Name: "John", Age: 30, Interest: []string{"badminton", "basketball"}}
	p2V2 := PersonV2{Name: "John", Age: 30, Interest: []string{"basketball", "badminton"}}

	if reflect.DeepEqual(p1V2, p2V2) {
		fmt.Println("p1V2 and p2V2 are equal")
	} else {
		fmt.Println("p1V2 and p2V2 are not equal")
	}

	if Compare1(&p1V2, &p2V2) {
		fmt.Println("cmp p1V2 and p2V2 are equal")
	} else {
		fmt.Println("cmp p1V2 and p2V2 are not equal")
	}

	m1 := map[string]string{"foo": "bar", "baz": "qux"}
	m2 := map[string]string{"foo": "bar", "baz": "qux"}

	if reflect.DeepEqual(m1, m2) {
		fmt.Println("m1 and m2 are equal")
	} else {
		fmt.Println("m1 and m2 are not equal")
	}
}
