package main

import (
	"fmt"
	"strings"
)

// leetcode 20. 有效的括号
func main() {
	str1 := "([]])"
	printValid(str1)

	str1 = "([])"
	printValid(str1)

	str1 = "([{}])"
	printValid(str1)

	str1 = "()[{}]"
	printValid(str1)

	str1 = "()[{}]]"
	printValid(str1)

	str1 = "()[][{}]"
	printValid(str1)

	str1 = "()[]{}(){[]}"
	printValid(str1)
}

func printValid(str string) {
	ret := isValidByStringReplace(str)
	ret2 := isValidBySlice(str)
	ret3 := isValidByStack(str)
	fmt.Printf("%s, isValid:%t, %t, %t\n", str, ret, ret2, ret3)
}
func isValidByStringReplace(str string) bool {
	if len(str) == 0 {
		return true
	}

	if len(str)%2 != 0 {
		return false
	}

	for {
		find1 := strings.Contains(str, "()")
		find2 := strings.Contains(str, "[]")
		find3 := strings.Contains(str, "{}")
		if find1 || find2 || find3 {
			str = strings.Replace(str, "()", "", -1)
			str = strings.Replace(str, "[]", "", -1)
			str = strings.Replace(str, "{}", "", -1)
		} else {
			break
		}
	}

	return len(str) == 0
}

func isValidBySlice(str string) bool {
	if str == "" {
		return true
	}
	if len(str)%2 != 0 {
		return false
	}
	// rune 类型通常用来处理 Unicode 或 UTF-8 字符
	stack := make([]rune, len(str))
	i := 0
	for _, value := range str {
		switch value {
		//只有括号的左边需要入栈
		case '(', '[', '{':
			stack[i] = value
			i++
		case ')':
			if i-1 >= 0 && stack[i-1] == '(' {
				i--
			} else {
				return false
			}
		case ']':
			if i-1 >= 0 && stack[i-1] == '[' {
				i--
			} else {
				return false
			}
		case '}':
			if i-1 >= 0 && stack[i-1] == '{' {
				i--
			} else {
				return false
			}
		}
	}

	if i == 0 {
		return true
	} else {
		return false
	}
}

func isValidByStack(str string) bool {
	if str == "" {
		return true
	}
	if len(str)%2 != 0 {
		return false
	}
	var myStack Stack

	for _, value := range str {
		switch value {
		//只有括号的左边需要入栈
		case '(', '[', '{':
			myStack.Push(value)
		case ')':
			peekVal, success := myStack.Peek()
			if success && peekVal == '(' {
				myStack.Pop()
			} else {
				return false
			}
		case ']':
			peekVal, success := myStack.Peek()
			if success && peekVal == '[' {
				myStack.Pop()
			} else {
				return false
			}
		case '}':
			peekVal, success := myStack.Peek()
			if success && peekVal == '{' {
				myStack.Pop()
			} else {
				return false
			}
		}
	}

	if myStack.IsEmpty() {
		return true
	} else {
		return false
	}
}

type Stack []rune

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(val rune) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Peek() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		return element, true
	}
}

/* 运行结果
[]]), isValid:false, false, false
([]), isValid:true, true, true
([{}]), isValid:true, true, true
()[{}], isValid:true, true, true
()[{}]], isValid:false, false, false
()[][{}], isValid:true, true, true
()[]{}(){[]}, isValid:true, true, true
*/
