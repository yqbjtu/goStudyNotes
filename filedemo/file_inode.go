package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("fail to get current directory, msg:", err)
		return
	}
	fmt.Println("CurrentDir:", dir)

	fileName := "archive.zip"
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		fmt.Println("Not a syscall.Stat_t")
		return
	}

	fmt.Printf("Inode: %+v\n",  stat)
}
