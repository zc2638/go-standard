package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
)

// 执行追踪器
// 跟踪器捕获各种各样的执行事件，如 goroutine 创建/阻塞/解锁，系统调用进入/退出/块，GC 相关事件，堆大小变化，处理器启动/停止等，并将它们以紧凑的形式写入 io.Writer 中
// 大多数事件都会捕获精确的纳秒精度时间戳和堆栈跟踪。跟踪可以稍后使用 'go tool trace' 命令进行分析
func main() {

	// 创建trace.out文件
	// 跟踪完毕后可以使用 "go tool trace testdata/trace.out" 命令分析
	f, err := os.Create("testdata/trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	// 启用当前程序的跟踪
	// 跟踪时，跟踪将被缓冲并写入 w
	// 如果跟踪已启用，则启动将返回错误
	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	// 停止当前跟踪，如果有的话。在完成跟踪的所有写入后，仅停止返回
	defer trace.Stop()

	// 下面写自己的程序

	traceTest()
}

func traceTest() {
	fmt.Printf("this function will be traced\n")
}
