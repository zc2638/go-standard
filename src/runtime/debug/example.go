package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

// debug包提供了程序在运行时进行调试的功能
func main() {

	// 强制进行一次垃圾收集，以释放尽量多的内存回操作系统。（即使没有调用，运行时环境也会在后台任务里逐渐将内存释放给系统）
	debug.FreeOSMemory()

	// 返回嵌入在运行二进制文件中的构建信息
	// 该信息仅在使用模块支持构建的二进制文件中可用
	if info, ok := debug.ReadBuildInfo(); ok {

		// main包路径
		fmt.Println(info.Path)
		// module信息
		fmt.Println(info.Main)
		// module依赖
		fmt.Println(info.Deps)
	}

	// 设置垃圾回收百分比
	exampleGCPercent()

	// 设置被单个go协程调用栈可使用内存值
	exampleSetMaxStack()

	// 设置go程序可以使用的最大操作系统线程数
	exampleSetMaxThreads()

	// 设置程序请求运行是只触发panic,而不崩溃
	exampleSetPanic()

	// 垃圾收集信息
	exampleStats()

	// 将内存分配堆和其中对象的描述写入文件中
	exampleHeapDump()

	// 获取go协程调用栈踪迹
	exampleStack()
}

func exampleGCPercent() {

	// 设定垃圾收集的目标百分比
	// 当新申请的内存大小占前次垃圾收集剩余可用内存大小的比率达到设定值时，就会触发垃圾收集
	// SetGCPercent返回之前的设定。初始值为环境变量GOGC的值；如果没有设置该环境变量，初始值为100
	// percent参数如果是负数值，会关闭垃圾收集
	fmt.Println(debug.SetGCPercent(1))

	// 初始化切片
	var dic = make([]byte, 100, 100)

	// 将x的终止器设置为f
	// 当垃圾收集器发现一个不能接触的（即引用计数为零，程序中不能再直接或间接访问该对象）具有终止器的块时，它会清理该关联（对象到终止器）并在独立go程调用f(x)
	// 这使x再次可以接触，但没有了绑定的终止器。如果SetFinalizer没有被再次调用，下一次垃圾收集器将视x为不可接触的，并释放x
	// SetFinalizer(x, nil)会清理任何绑定到x的终止器
	runtime.SetFinalizer(&dic, func(dic *[]byte) {
		fmt.Println("内存回收1")
	})

	// 立即回收
	runtime.GC()

	var s = make([]byte, 100, 100)
	runtime.SetFinalizer(&s, func(dic *[]byte) {
		fmt.Println("内存回收2")
	})

	d := make([]byte, 300, 300)

	for index, _ := range d {
		d[index] = 'a'
	}
	fmt.Println(d)

	time.Sleep(time.Second)
}

func exampleSetMaxStack() {

	// 设置被单个go协程调用栈可使用的内存最大值
	// 修改该函数设置值为1，再次运行 exampleSetMaxStack()即可看到效果
	debug.SetMaxStack(102400)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(1)
		}()
	}
	time.Sleep(time.Second)
}

func exampleSetMaxThreads() {

	// 置go程序可以使用的最大操作系统线程数
	// 如果程序试图使用超过该限制的线程数，就会导致程序崩溃
	// SetMaxThreads返回之前的设置，初始设置为10000个线程
	// 主要用于限制程序无限制的创造线程导致的灾难。目的是让程序在干掉操作系统之前，先干掉它自己
	debug.SetMaxThreads(10)

	go func() {
		fmt.Println("Hello World!")
	}()
	time.Sleep(time.Second)
}

func exampleSetPanic() {

	go func() {
		defer func() { recover() }()

		// 控制程序在不期望（非nil）的地址出错时的运行时行为
		// 这些错误一般是因为运行时内存破坏的bug引起的，因此默认反应是使程序崩溃
		// 使用内存映射的文件或进行内存的不安全操作的程序可能会在非nil的地址出现错误；SetPanicOnFault允许这些程序请求运行时只触发一个panic，而不是崩溃
		// SetPanicOnFault只用于当前的go程。它返回之前的设置
		fmt.Println(debug.SetPanicOnFault(true))
		var s *int = nil
		*s = 34
	}()

	time.Sleep(time.Second)

	fmt.Println("ddd")
}

func exampleStats() {

	// 初始化切片
	data := make([]byte, 1000, 1000)
	println(data)

	// 回收
	runtime.GC()

	// 声明一个GCStats，用于收集了近期垃圾收集的信息
	var stats debug.GCStats

	// 将垃圾收集信息填入stats里
	debug.ReadGCStats(&stats)

	// 垃圾收集的次数
	fmt.Println(stats.NumGC)

	// 最近一次垃圾收集的时间
	fmt.Println(stats.LastGC)

	// 每次暂停收集垃圾的消耗的时间
	fmt.Println(stats.Pause)

	// 所有暂停收集垃圾消耗的总时间
	fmt.Println(stats.PauseTotal)

	// 暂停结束时间历史记录，最近的第一个
	fmt.Println(stats.PauseEnd)
}

func exampleHeapDump() {

	// 打开文件
	f, _ := os.OpenFile("testdata/debug_heapDump.txt", os.O_RDWR|os.O_CREATE, 0666)

	// 返回与文件f对应的整数类型的Unix文件描述符
	fd := f.Fd()

	// 将内存分配堆和其中对象的描述写入给定文件描述符fd指定的文件
	debug.WriteHeapDump(fd)

	data := make([]byte, 10, 10)
	println(data)
	runtime.GC()

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func exampleStack() {

	go func() {
		// 返回格式化的go程的调用栈踪迹
		// 对于每一个调用栈，它包括原文件的行信息和PC值；对go函数还会尝试获取调用该函数的函数或方法，及调用所在行的文本
		fmt.Println(string(debug.Stack()))

		// 将Stack返回信息打印到标准错误输出
		// debug.PrintStack()
	}()
	time.Sleep(time.Second)
}