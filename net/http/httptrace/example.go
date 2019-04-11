package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
)

// httptrace包提供跟踪HTTP客户端请求中的事件的机制
// http钩子
func main() {

	// 创建GET请求
	req, _ := http.NewRequest("GET", "http://example.com", nil)

	// 声明一组挂钩
	// 可以在传出的HTTP请求的各个阶段运行，任何特定的钩子可能是零
	// 函数可以从不同的goroutine同时调用，有些可能在请求完成或失败后调用
	trace := &httptrace.ClientTrace{

		// 在成功连接后调用
		// 没有用于连接失败的钩子；相反，使用transport.roundtrip中的错误
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},

		// 在DNS查找结束时调用
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
	}

	// 根据提供的ctx返回一个新的Context
	// 除了使用ctx注册的任何以前的挂钩外，使用返回的Context创建的HTTP客户端请求将使用提供的跟踪挂钩。在提供的轨迹中定义的任何钩子将首先被调用
	ctx := httptrace.WithClientTrace(req.Context(), trace)

	// 修改r中的Context并返回一个浅表副本
	// ctx不能为nil
	// 对于传出客户端请求，上下文控制请求及其响应的整个生命周期：获取连接，发送请求以及读取响应头和正文
	req = req.WithContext(ctx)

	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
}