package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

// pprof包以pprof可视化工具期望的格式书写运行时剖面数据
// 采集程序的运行数据进行分析，使用 'go tool pprof' 命令进行分析
func main() {

	// 创建pprof文件
	// 使用 "go tool pprof testdata/test.prof" 进入，如 使用top查看数据，web生成svg文件
	f, err := os.Create("testdata/test.prof")
	if err != nil {
		log.Fatal(err)
	}

	// 设置CPU profile记录的速率为平均每秒hz次
	// 如果hz<=0，SetCPUProfileRate会关闭profile的记录。如果记录器在执行，该速率必须在关闭之后才能修改
	runtime.SetCPUProfileRate(10000)

	// 为当前进程开启CPU profile
	// 在分析时，分析报告会缓存并写入到w中。若分析已经开启，StartCPUProfile就会返回错误
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	// 停止当前的CPU profile（如果有）
	// 只会在所有的分析报告写入完毕后才会返回
	defer pprof.StopCPUProfile()

	// 下面写自己的程序

	for i := 0; i < 30; i++ {
		// 递归实现的斐波纳契数列
		num := fibonacci(i)
		fmt.Println(num)
	}
}

func fibonacci(num int) int {

	if num < 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
