package main

import (
	"context"
	"time"
)

// Context包定义了上下文类型
func main() {

	// 返回一个不是nil的空的Context。当不清楚使用哪个Context或者不可用时（因为功能尚未扩展到接受Context参数）。
	context.TODO()

	// 返回一个不是nil的空Context。它永远不会被取消，没有值，也没有存活时间。 它通常由主函数，初始化和测试使用，并作为传入请求的顶级Context。
	ctx := context.Background()

	// 根据Context，返回一个Context副本和一个取消方法，以方便完成工作后释放资源
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	// 根据Context和最迟存活时间，返回一个Context副本和一个取消方法，以方便完成工作后释放资源
	context.WithDeadline(ctx, time.Now().Add(50 * time.Millisecond))

	// 根据Context和超时时间，返回一个Context副本和一个取消方法，以方便完成工作后释放资源
	context.WithTimeout(ctx, time.Millisecond * 50)

	// 给Context添加自定义类型的键值对，返回一个新的Context
	context.WithValue(ctx, "language", "Go")
}