package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用
// 这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好
// 应通过通信来共享内存，而不通过共享内存实现通信
func main() {

	// 自定义value
	exampleValue()
	// 类型示例
	example()
}

func exampleValue() {

	cfg := make(map[string]int)
	// 根基当前时间的unix值初始化rand状态
	rand.Seed(time.Now().Unix())
	// 随机生成500以内的随机数
	cfg["test"] = rand.Intn(500)

	// 声明atomic.Value
	// 提供了一个自动加载和一个一致的类型值的存储。Value 的零值从 Load 返回 nil 。一旦 Store 被调用，Value 不能被复制
	// 首次使用后不得复制 Value
	var config atomic.Value

	// 将 Value 的值设置为 x
	// 对于给定值的所有对Store的调用都必须使用相同具体类型的值。存储不一致的类型 会panic，就像 Store（nil） 一样
	config.Store(cfg)

	// 返回最近的存储设置的值。如果此值没有存储调用，则返回 nil
	c := config.Load()
	fmt.Println(c)
}

func example() {

	var i32 int32

	// 更新 int32类型 内存值
	atomic.StoreInt32(&i32, 15)
	// 获取 int32类型 内存值
	fmt.Println(atomic.LoadInt32(&i32))
	// 将值添加到int32内存中（值+内存值）
	fmt.Println(atomic.AddInt32(&i32, 12))
}