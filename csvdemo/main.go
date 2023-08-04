package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory: ", err)
		return
	}

	fmt.Println("Current directory: ", dir)
	// 打开 CSV 文件
	file, err := os.Open("./csvdemo/dataset_0801.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建新的 CSV reader
	reader := csv.NewReader(file)

	// 读取 CSV 数据
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 遍历 CSV 记录
	// 遍历 CSV 记录
	for i, record := range records {
		fmt.Printf("Row %d has %d columns:\n", i+1, len(record))
		// 遍历行中的每一列
		for j, value := range record {
			fmt.Printf("Column %d: %s\n", j+1, value)
		}
		fmt.Println("-----------------")
	}
}
