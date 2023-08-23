package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func getOpenFilesByPid(pid int) ([]string, error) {
	path := fmt.Sprintf("/proc/%d/fd", pid)
	filenames := []string{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		filePath := path + "/" + file.Name()
		realPath, err := os.Readlink(filePath)
		if err != nil {
			continue // 当遇到错误时，忽略并处理下一个文件描述符
		}
		filenames = append(filenames, realPath)
	}

	return filenames, nil
}

// GOOS=linux GOARCH=amd64 go build -o file_fd_demo file_fd_demo.go
// 查看一个进程都打开了哪些文件，linux可以运行，mac不行
func main() {
	pid := os.Getpid() // 得到当前程序的进程 ID
	fmt.Println("Current PID:", pid)

	if len(os.Args) < 2 {
		fmt.Println("Please provide a PID")
		return
	}

	pidStr := os.Args[1]
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		fmt.Printf("Invalid PID: %s\n", pidStr)
		return
	}

	openFiles, err := getOpenFilesByPid(pid)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for i, filename := range openFiles {
		fmt.Printf("%d: %s\n", i+1, filename)
	}
}
