package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// 计算文件的md5 如果文件内容较小，可以一次全部读取，如果文件内容加大，不能全部加载，否则内部不够

func calcFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	md5Handle := md5.New()
	_, _ = io.Copy(md5Handle, file)

	md5Byte := md5Handle.Sum(nil)        //计算 MD5 值，返回 []byte
	md5str := fmt.Sprintf("%x", md5Byte) //将 []byte 转为 string
	fmt.Printf("md5Byte:%v, md5str:%v\n", md5Byte, md5str)
	return hex.EncodeToString(md5Byte), nil
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory: ", err)
		return
	}

	fmt.Println("Current dir: ", dir)

	var fileName = "file1.txt"
	md5Val, err := calcFileMD5(fileName)
	if err != nil {
		fmt.Printf("fail to calc md5 for file %v, errMsg:%v\n", fileName, err)
	} else {
		fmt.Printf("md5 for file %v is %v\n", fileName, md5Val)
	}
}

//运行结果
//Current dir:  /Users/ericyang/GolandProjects/mygotutorials2/typecast
//md5Byte:[143 226 153 178 250 103 61 141 29 89 151 253 122 103 122 29], md5str:8fe299b2fa673d8d1d5997fd7a677a1d
//md5 for file file1.txt is 8fe299b2fa673d8d1d5997fd7a677a1d
