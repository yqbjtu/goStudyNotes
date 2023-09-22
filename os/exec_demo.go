package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println("output to std1")
	printToStd()

	fmt.Println("output to std2")
	printToStd2()

	fmt.Println("read from std")
	readFromStd()

	fmt.Println("run with multi args")
	execWithMultiArgs()

	fmt.Println("主函数执行完毕")
}

func printToStd() {
	cmd := exec.Command("ls", "-sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "/"
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to call cmd.Run(): %v", err)
	}
	fmt.Println()
}

func printToStd2() {
	//Output runs the command and returns its standard output
	out, err := exec.Command("ls", "-l").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func readFromStd() {
	// tr - translate or delete characters
	cmd := exec.Command("tr", "a-z", "A-Z")

	cmd.Stdin = strings.NewReader("and old falcon")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("translated phrase: %q\n\n", out.String())
}
func execWithMultiArgs() {
	prg := "echo"

	arg1 := "there"
	arg2 := "are three"
	arg3 := "falcons"

	cmd := exec.Command(prg, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}

//运行结果
//output to std1
//total 10
//0 AppleInternal         0 Volumes               0 etc                   0 osmdb-core            0 usr
//0 Library               0 coredata              0 opt                   0 private
//0 System                0 cores                 0 osmdb-checkin         0 sbin
//0 Users                10 dev                   0 osmdb-checkout        0 tmp
//
//output to std2
//total 24
//-rw-r--r--  1 ericyang  staff  1357  9 22 19:36 exec_demo.go
//-rw-r--r--  1 ericyang  staff  1887  9 20 12:44 string1111main.go
//-rw-r--r--  1 ericyang  staff   698  9 19 20:49 toString_demo.go
//
//read from std
//translated phrase: "AND OLD FALCON"
//
//run with multi args
//there are three falcons
//
//主 函 数 执 行 完 毕
