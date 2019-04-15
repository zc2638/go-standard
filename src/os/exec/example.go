package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// exec包执行外部命令
// 它包装了os.StartProcess函数以便更容易的修正输入和输出，使用管道连接I/O，以及作其它的一些调整
func main() {

	// 执行命令
	exampleCommand()
	// 查询环境变量
	exampleLookPath()
}

func exampleLookPath() {

	// 在环境变量PATH指定的目录中搜索可执行文件，如file中有斜杠，则只在当前目录搜索
	// 返回完整路径或者相对于当前目录的一个相对路径
	path, err := exec.LookPath("fortune")
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Printf("fortune is available at %s\n", path)
	}
}

func exampleCommand() {

	cmd := exec.Command("go", "doc", "exec")

	// 设置输入内容
	cmd.Stdin = strings.NewReader("")

	// 声明buffer
	var out bytes.Buffer

	// 设置输出内容填充地址
	cmd.Stdout = &out

	// 执行c包含的命令，并阻塞直到完成
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
}

