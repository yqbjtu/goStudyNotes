package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("fail to get current directory, msg:", err)
		return
	}
	fmt.Println("CurrentDir:", dir)

	unZipDst := "output"

	var zipFile = "archive.zip"
	_, err = os.Stat(zipFile)
	if os.IsNotExist(err) {
		fullPath, _ := filepath.Abs(zipFile)
		fmt.Printf("File '%v' does not exist.\v", fullPath)
		return
	}

	archive, err := zip.OpenReader(zipFile)
	if err != nil {
		fmt.Printf("fail to open compressedFile:%v, errMsg:%v\n", zipFile, err)
		return
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(unZipDst, f.Name)
		fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(unZipDst)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			return
		}
		if f.FileInfo().IsDir() {
			fmt.Printf("creating directory:%v\n", f.Name)
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
}
