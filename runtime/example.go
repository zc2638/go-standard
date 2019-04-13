package main

import (
	"fmt"
	"runtime"
	"strings"
)

// runtime包提供和go运行时环境的互操作，如控制go程的函数
// 它也包括用于reflect包的低层次类型信息
func main() {

	// 基础
	example()

	// 显示调用过程
	exampleFrames()
}

func example() {

	// 返回Go的根目录。如果存在GOROOT环境变量，返回该变量的值；否则，返回创建Go时的根目录
	fmt.Println(runtime.GOROOT())

	// 返回Go的版本字符串。它要么是递交的hash和创建时的日期；要么是发行标签如"go1.3"
	fmt.Println(runtime.Version())

	// 返回本地机器的逻辑CPU个数
	fmt.Println(runtime.NumCPU())

	// 返回当前进程执行的cgo调用次数
	fmt.Println(runtime.NumCgoCall())

	// 返回当前存在的Go程数
	fmt.Println(runtime.NumGoroutine())

	// 执行一次垃圾回收
	runtime.GC()

	// 设置可同时执行的最大CPU数，并返回先前的设置
	// 若 n < 1，它就不会更改当前设置
	// 本地机器的逻辑CPU数可通过 NumCPU 查询。本函数在调度程序优化后会去掉
	runtime.GOMAXPROCS(2)

	go func() {

		fmt.Println("Start Goroutine1")

		// 终止调用它的go程。其它go程不会受影响。Goexit会在终止该go程前执行所有defer的函数。
		// 在程序的main go程调用本函数，会终结该go程，而不会让main返回
		// 因为main函数没有返回，程序会继续执行其它的go程
		// 如果所有其它go程都退出了，程序会panic
		runtime.Goexit()

		fmt.Println("End Goroutine1")
	}()

	go func() {

		// 使当前go程放弃处理器，以让其它go程运行。它不会挂起当前go程，因此当前go程未来会恢复执行
		runtime.Gosched()
		fmt.Println("Start Goroutine2")
		fmt.Println("End Goroutine2")
	}()

	go func() {

		// 将调用的go程绑定到它当前所在的操作系统线程
		// 除非调用的go程退出或调用UnlockOSThread，否则它将总是在该线程中执行，而其它go程则不能进入该线程
		runtime.LockOSThread()
		fmt.Println("Start Goroutine3")
		fmt.Println("End Goroutine3")

		// 将调用的go程解除和它绑定的操作系统线程
		// 若调用的go程未调用LockOSThread，UnlockOSThread不做操作
		runtime.UnlockOSThread()
	}()
}

func exampleFrames() {

	c := func() {

		// 初始化一个长度为10的整数切片
		pc := make([]uintptr, 10)

		// 把当前go程调用栈上的调用栈标识符填入切片pc中，返回写入到pc中的项数
		// 实参skip为开始在pc中记录之前所要跳过的栈帧数，0表示Callers自身的调用栈，1表示Callers所在的调用栈
		n := runtime.Callers(0, pc)
		if n == 0 {
			return
		}

		// 截取有效部分
		pc = pc[:n]

		// 获取被调用方返回的PC值切片，并准备返回 函数/文件/行 信息
		// 在完成帧处理之前，不要更改切片
		frames := runtime.CallersFrames(pc)

		for {
			// 返回下一个调用者的帧信息
			// 如果 more为false，则无下一个（Frame 值有效）
			frame, more := frames.Next()

			// 判断文件名是否包含 "runtime/" 字符串
			if !strings.Contains(frame.File, "runtime/") {
				break
			}
			fmt.Printf("- more:%v | %s\n", more, frame.Function)
			if !more {
				break
			}
		}
	}
	b := func() { c() }
	a := func() { b() }

	a()
}