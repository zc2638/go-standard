package main

import (
	"log"
	"os"
)

const LogFile = "testdata/testLog"

// log包实现了简单的日志服务
func main() {

	// 创建日志文件
	file, err := os.Create(LogFile)
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个Logger
	// 参数out设置日志信息写入的目的地
	// 参数prefix会添加到生成的每一条日志前面
	// 参数flag定义日志的属性（时间、文件等等）
	l := log.New(file, "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)

	// 设置logger的输出选项(日志属性)
	l.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	// 设置logger的输出前缀
	l.SetPrefix("")
	// 设置标准logger的输出目的地，默认是标准错误输出
	l.SetOutput(file)

	// 写入输出一次日志事件
	// 参数s包含在Logger根据选项生成的前缀之后要打印的文本，如果s末尾没有换行会添加换行符
	// calldepth用于恢复PC，出于一般性而提供，但目前在所有预定义的路径上它的值都为2
	l.Output(2, "Hello World!")
	// 返回标准logger的输出前缀
	l.Prefix()
	// 返回标准logger的输出选项
	l.Flags()

	l.Println("ok")
	// 其它用法大致与fmt包类似
}
