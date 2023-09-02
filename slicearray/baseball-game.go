package main

import (
	"fmt"
	"strconv"
)

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		return element, true
	}
}

// leetcode 682. 棒球比赛 计算得分
func main() {
	ops := []string{"5", "2", "C", "D", "+"} // 结果30
	printGameScore(ops)

	ops = []string{"5", "-2", "4", "C", "D", "9", "+", "+"} //结果27
	printGameScore(ops)

	ops = []string{"5", "C", "4", "C", "D", "9", "+", "+"} // -1, 因为第一个C将5取消了，第二个C将4取消了，到了D前面没有数字
	printGameScore(ops)
}

func printGameScore(str []string) {
	ret := computerByStack(str)
	fmt.Printf("%s, sum:%d\n", str, ret)
}

func computerByStack(ops []string) int {
	if len(ops) == 0 {
		return 0
	}

	var myStack Stack

	for _, value := range ops {
		num, err := strconv.Atoi(value)
		if err == nil {
			//如果是纯数字直接放入栈中
			myStack.Push(num)
		} else {
			switch value {
			//只有括号的左边需要入栈
			case "C":
				//表示前一次得分无效，将其从记录中移除。题目数据保证记录此操作时前面总是存在一个有效的分数
				if myStack.IsEmpty() {
					fmt.Printf("invalud array %s since there is no num before C\n", ops)
					return -1
				}
				myStack.Pop()
			case "D":
				// 表示本回合新获得的得分是前一次得分的两倍。题目数据保证记录此操作时前面总是存在一个有效的分数。
				peekVal, success := myStack.Peek()
				if success {
					newVal := peekVal * 2
					myStack.Push(newVal)
				} else {
					fmt.Printf("invalud array %s since there is no num before D\n", ops)
					return -1
				}
			case "+":
				//表示本回合新获得的得分是前两次得分的总和。题目数据保证记录此操作时前面总是存在两个有效的分数。
				popVal, popOk := myStack.Pop()
				peekVal, peekOk := myStack.Peek()
				if popOk && peekOk {
					newVal := popVal + peekVal
					//老的栈顶元素先push
					myStack.Push(popVal)
					myStack.Push(newVal)
				} else {
					fmt.Printf("invalud array %s since there is no two num before +\n", ops)
					return -1
				}
			}
		}
	}

	if myStack.IsEmpty() {
		return 0
	} else {
		//栈中所有元素进行相加
		sum := 0
		for !myStack.IsEmpty() {
			popVal, _ := myStack.Pop()
			sum = sum + popVal
		}
		return sum
	}
}

