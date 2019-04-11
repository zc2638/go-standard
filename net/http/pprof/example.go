package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// pprof包通过它的HTTP服务端提供pprof可视化工具期望格式的运行时剖面文件数据服务
// 本包一般只需导入获取其注册HTTP处理器的副作用。处理器的路径以/debug/pprof/开始
func main() {

	// 启动服务
	log.Println(http.ListenAndServe("localhost:6060", nil))

	/*
		使用pprof工具查看堆剖面
		go tool pprof http://localhost:6060/debug/pprof/heap

		查看周期30秒的CPU剖面
		go tool pprof http://localhost:6060/debug/pprof/profile

		查看go程阻塞剖面
		go tool pprof http://localhost:6060/debug/pprof/block

		浏览器可直接访问http://localhost:6060/debug/pprof/查看所有信息
	*/


}