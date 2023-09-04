package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	slow := 1
	for fast := 1; fast < n; fast++ {
		// 快指针和自身下元素不相等，怎将它放到慢指针对应值，并且慢指针加1
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

// leetcode 26. 删除有序数组中的重复项
func main() {
	//此处使用了切片
	ops := []int{1, 1, 2} // 结果 2, [1,2,_]
	fmt.Printf("originalArray%v\n", ops)
	printRemoveDuplicateResult(ops)

	ops = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4} //结果5, [0,1,2,3,4]
	fmt.Printf("originalArray%v\n", ops)
	printRemoveDuplicateResult(ops)
}

func printRemoveDuplicateResult(intArray []int) {
	ret := removeDuplicates(intArray)
	fmt.Printf("%v, sum:%v\n", intArray, ret)
}
