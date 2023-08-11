package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func WriterCSV(path string) {

	//OpenFile读取文件，不存在时则创建，使用追加模式
	File, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("文件打开失败！")
	}
	defer File.Close()

	//创建写入接口
	WriterCsv := csv.NewWriter(File)
	str1 := []string{"zhangsan", "beijing", "2020"} //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	err = WriterCsv.Write(str1)
	if err != nil {
		log.Println("WriterCsv写入文件失败")
	}

	str2 := []string{"lisi", "shanghai", "2021"}
	str3 := []string{"wangwu", "taibei", "2022"}
	str4 := []string{"zhaoliu", "xian", "2023"}
	var values [][]string
	values = append(values, str2)
	values = append(values, str3)
	values = append(values, str4)

	// 循环输出
	for i := range values {
		fmt.Printf("Row: %v\n", i)
		err = WriterCsv.Write(values[i])
		if err != nil {
			log.Println("WriterCsv写入文件失败")
		}
	}

	WriterCsv.Flush() //刷新，不刷新是无法写入的
	log.Println("数据写入成功...")
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory: ", err)
		return
	}

	fmt.Println("Current directory: ", dir)
	// 打开 CSV 文件
	var fileName = "demo1.csv"
	WriterCSV(fileName)

	fmt.Println("done. write csv file: ", fileName)
}

//运行结果：
//Current directory:  /Users/ericyang/GolandProjects/mygotutorials2/csvdemo
//Row: 0
//Row: 1
//Row: 2
//2023/08/11 20:58:59 数据写入成功...
//done. write csv file:  demo1.csv
