package main

import "fmt"

// leetcode 128 
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题
func longestConsecutive(nums []int) int {
	numsMap := make(map[int]int)

	lenNums := len(nums)

	for i := 0; i < lenNums; i++ {
		numsMap[nums[i]] = 0
	}

	maxLCS := 0
	for i := 0; i < lenNums; i++ {
		currentLen := 1
		counter := 1

		for {
			val, ok := numsMap[nums[i]+counter]
			if ok {
				if val != 0 {
					currentLen += val
					break
				} else {
					currentLen += 1
					counter += 1
				}
			} else {
				break
			}
		}

		if currentLen > maxLCS {
			maxLCS = currentLen
		}
		numsMap[nums[i]] = currentLen
	}

	return maxLCS
}

func main() {
	output := longestConsecutive([]int{4, 7, 3, 8, 2, 1})
	fmt.Println(output)

	output = longestConsecutive([]int{4, 7, 3, 8, 2, 1, 9, 24, 10, 11})
	fmt.Println(output)
}
