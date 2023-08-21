package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("fail to get current directory, msg:", err)
		return
	}
	fmt.Println("CurrentDir:", dir)

	var compressFile = "archive.zip"
	var toBeCompressedDir = "dir1"
	file, err := os.Create(compressFile)
	if err != nil {
		fmt.Printf("fail to create compressFile:%v, errMsg:%v\n", compressFile, err)
		return
	}
	defer file.Close()

	archive := zip.NewWriter(file)
	defer archive.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, _ := zip.FileInfoHeader(info)
		header.Name = path

		if info.IsDir() {
			// 没有这行， 会导致压缩文件中不包含空文件夹
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	}

	_, err = os.Stat(toBeCompressedDir)
	if os.IsNotExist(err) {
		fullPath, _ := filepath.Abs(toBeCompressedDir)
		fmt.Printf("File '%v' does not exist.\v", fullPath)
		return
	}

	err = filepath.Walk(toBeCompressedDir, walker)
	if err != nil {
		fmt.Printf("fail to find toBeCompressedDir:%v, errMsg:%v\n", toBeCompressedDir, err)
	}

	// 这里必须调用close， 否则会因为压缩文件还没有关闭，导致读取报异常 :zip: not a valid zip file
	archive.Close()
	fmt.Printf("compress file content\n")
	listCompressContent(compressFile)
}

func listCompressContent(compressedFile string) {
	zipListing, err := zip.OpenReader(compressedFile)
	if err != nil {
		fmt.Printf("fail to open compressedFile:%v, errMsg:%v\n", compressedFile, err)
	}
	defer zipListing.Close()

	for _, file := range zipListing.File {
		fmt.Println(file.Name)
	}
}
