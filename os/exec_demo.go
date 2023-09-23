package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("output to std1, cmd can run normally")
	printToStdNormal()

	fmt.Println("output to std2, cmd can't run normally for the cmd args is not right")
	printToStdError1()

	fmt.Println("output to std3, cmd can't run normally for the cmd dir does not exist")
	printToStdError2()

	fmt.Println("output to std2")
	printToStd2()

	fmt.Println("read from std")
	readFromStd()

	fmt.Println("run with multi args")
	execWithMultiArgs()

	fmt.Println("主函数执行完毕")
}

func printToStdNormal() {
	//设置需要执行的命令
	cmd := exec.Command("ls", "-sh")
	//将命令的输出和错误与标准的os.Stdout关联
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//设置该命令执行的上下文目录信息
	// 不设置dir时， dir默认值为当前程序启动的位置
	cmd.Dir = "/"

	err := cmd.Run()
	if err != nil {
		fmt.Printf("failed to call cmd.Run(): %v\n", err)
	}
	fmt.Println()
}

func printToStdError1() {
	cmd := exec.Command("ls", "aaaa-sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "/"
	err := cmd.Run()
	if err != nil {
		//错误信息输出该cmd的参数
		fmt.Printf("failed to call cmd.Run() with args:%v,  err:%v\n", cmd.Args, err)
	}
	fmt.Println()
}

func printToStdError2() {
	cmd := exec.Command("ls", "-sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "/not-exist"
	err := cmd.Run()
	if err != nil {
		//错误信息输出该cmd的执行的目录
		fmt.Printf("failed to call cmd.Run() with args:%v,  err:%v\n", cmd.Dir, err)
	}
	fmt.Println()
}

func printToStd2() {
	//Output runs the command and returns its standard output
	arg1 := os.ExpandEnv("${MY_ENV_VAR}")
	// 目前读取动态自定义的环境变量有问题，通过os.ExpandEnv读取os已有的的环境环境变量没问题
	//或者可以先通过os.Setenv("MY_ENV_VAR", "some_value1")， 然后os.ExpandEnv
	cmd := exec.Command("echo", arg1)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "MY_ENV_VAR=some_value")

	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("failed to call cmd.Run() with args:%v, err:%v\n", cmd.Dir, err)
		return
	}

	fmt.Printf("cmd '%s' executed in dir '%s' (命令执行的默认目录), env:%s, out:'%s'\n",
		cmd.Args, cmd.Dir, cmd.Env, string(out))
}

func readFromStd() {
	// tr - translate or delete characters
	cmd := exec.Command("tr", "a-z", "A-Z")

	//设置执行命令的input信息
	cmd.Stdin = strings.NewReader("and old falcon")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		fmt.Printf("failed to call cmd.Run() with args:%v, err:%v\n", cmd.Args, err)
	}

	fmt.Printf("translated phrase: %q\n\n", out.String())
}
func execWithMultiArgs() {
	prg := "echo"

	arg1 := "there"
	arg2 := "are three"
	arg3 := os.ExpandEnv("${JAVA_HOME}")

	cmd := exec.Command(prg, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("cmd '%s' executed in dir '%s' (命令执行的默认目录), env:%s, out:%s\n",
		cmd.Args, cmd.Dir, cmd.Env, string(stdout))
}

/*运行结果
output to std1, cmd can run normally
total 10
0 AppleInternal         0 Volumes               0 etc                   0 osmdb-core            0 usr
0 Applications          0 bin                   0 home                  0 osmdb-output          0 var
0 Library               0 coredata              0 opt                   0 private
0 System                0 cores                 0 osmdb-checkin         0 sbin
0 Users                10 dev                   0 osmdb-checkout        0 tmp

output to std2, cmd can't run normally for the cmd args is not right
ls: aaaa-sh: No such file or directory
failed to call cmd.Run() with args:[ls aaaa-sh],  err:exit status 1

output to std3, cmd can't run normally for the cmd dir does not exist
failed to call cmd.Run() with args:/not-exist,  err:chdir /not-exist: no such file or directory

output to std2
cmd '[echo ]' executed in dir '' (命 令 执 行 的 默 认 目 录 ), env:[ALACRITTY_WINDOW_ID=0 COLORTERM=truecolor COMMAND_MODE=unix2003 DISABLE_AUTO_UPDATE=true GRADLE_HOME=/Users/ericyang/.sdkman/candidates/gradle/current HOME=/Users/ericyang HOMEBREW_CELLAR=/opt/homebrew/Cellar HOMEBREW_PREFIX=/opt/homebrew HOMEBREW_REPOSITORY=/opt/homebrew INFOPATH=/opt/homebrew/share/info:/opt/homebrew/share/info: INTELLIJ_TERMINAL_COMMAND_BLOCKS=1 JAVA_HOME=/Users/ericyang/.sdkman/candidates/java/current JAVA_MAIN_CLASS_8868=fleet.dock.bootstrap.FleetBootstrapKt LC_ALL=zh_CN.UTF-8 LOGIN_SHELL=1 LOGNAME=ericyang MANPATH=/opt/homebrew/share/man:/usr/share/man:/usr/local/share/man:/opt/homebrew/share/man: MAVEN_HOME=/Users/ericyang/software/apache-maven-3.8.6 OLDPWD=/Users/ericyang/GolandProjects/mygotutorials2/stringdemo PATH=/Users/ericyang/software/apache-maven-3.8.6/bin:/opt/homebrew/bin:/opt/homebrew/sbin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin:/Users/ericyang/software/apache-maven-3.8.6/bin:/Users/ericyang/.sdkman/candidates/java/current/bin:/Users/ericyang/.sdkman/candidates/gradle/current/bin:/opt/homebrew/bin:/opt/homebrew/sbin:/Users/ericyang/Library/Application Support/JetBrains/Toolbox/scripts:/Users/ericyang/Library/Application Support/JetBrains/Toolbox/scripts PWD=/Users/ericyang/GolandProjects/mygotutorials2/string SDKMAN_CANDIDATES_API=https://api.sdkman.io/2 SDKMAN_CANDIDATES_DIR=/Users/ericyang/.sdkman/candidates SDKMAN_DIR=/Users/ericyang/.sdkman SDKMAN_PLATFORM=darwinarm64 SDKMAN_VERSION=5.15.0 SHELL=/bin/zsh SHLVL=1 SSH_AUTH_SOCK=/private/tmp/com.apple.launchd.XATn15hl3j/Listeners TERM=xterm-256color TERM_PROGRAM=Jetbrains.Fleet TERM_PROGRAM_VERSION=1.22.113 TMPDIR=/var/folders/ls/8nvj13l13b59rpk8c1jw_tg00000gn/T/ USER=ericyang WINDOWID=0 XPC_FLAGS=0x0 XPC_SERVICE_NAME=0 __CFBundleIdentifier=Fleet.app __CF_USER_TEXT_ENCODING=0x1F5:0x19:0x34 _=/usr/local/go/bin/go MY_ENV_VAR=some_value], out:'
'
read from std
translated phrase: "AND OLD FALCON"

run with multi args
cmd '[echo there are three /Users/ericyang/.sdkman/candidates/java/current]' executed in dir '' (命 令 执 行 的 默 认 目 录 ), env:[], out:there are three /Users/ericyang/.sdkman/candidates/java/current

主 函 数 执 行 完 毕
*/
