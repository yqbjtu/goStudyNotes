package main

import (
	"fmt"
	"math"
)

func roundToDecimalPlace(num float64, decimalPlace int) float64 {
	divisor := math.Pow(10, float64(decimalPlace))
	return math.Round(num*divisor) / divisor
}

func main() {
	num := 1.45678
	fmt.Println(roundToDecimalPlace(num, 3)) // 输出: 1.457

	num = -1.45
	fmt.Println(roundToDecimalPlace(num, 3)) // 输出: -1.45

	num = -1.002817
	fmt.Println(roundToDecimalPlace(num, 3)) // 输出: -1.003

	num = 5000
	fmt.Println(roundToDecimalPlace(num, 3)) // 输出: 5000

	// 向上取整 Ceil 示例
	fmt.Println(math.Ceil(1.2)) // 输出: 2 （1.4 向上取整为2）

	// Round 示例
	fmt.Println(math.Round(1.4)) // 输出: 1 （1.4 四舍五入为1）
	fmt.Println(math.Round(1.5)) // 输出: 2 （1.5 四舍五入为2）

	//向下取整
	// Floor 示例
	fmt.Println(math.Floor(1.4)) // 输出: 1 （1.4 向下取整为1）
	fmt.Println(math.Floor(1.6)) // 输出: 1 （1.6 向下取整为1）

}
