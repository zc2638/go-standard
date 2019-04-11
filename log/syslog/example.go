package main

import (
	"log"
	"log/syslog"
)

// syslog 包提供一个简单的系统日志服务的接口
func main() {

	// 通过连接到指定网络上的地址来建立到日志守护程序的连接
	// 每次写入返回的时候都会发送一条日志消息，其中包含设施和严重性（来自优先级）和标记
	// 如果标签为空，则使用 os.Args0。如果网络为空，拨号将连接到本地系统日志服务器。否则，请参阅 net.Dial 的文档以获取网络和 raddr 的有效值
	sl, err := syslog.Dial("tcp", "localhost:1234",
		syslog.LOG_WARNING|syslog.LOG_DAEMON|syslog.LOG_INFO, "testTag")
	if err != nil {
		log.Fatal(err)
	}
	// 写入
	if _, err := sl.Write([]byte("Hello World!")); err != nil {
		log.Fatal(err)
	}
	// 关闭连接
	if err := sl.Close(); err != nil {
		log.Fatal(err)
	}

	// 建立到系统日志守护进程的新连接
	// 每次写入返回的写入程序都会发送一条具有给定优先级（syslog 设施和严重性的组合）和前缀标记的日志消息
	// 如果标签为空，则使用 os.Args0
	news, err := syslog.New(syslog.LOG_DEBUG, "testNew")
	defer news.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个 log.Logger，它的输出以指定的优先级写入系统日志服务
	// 这是 syslog 设施和严重性的组合。logFlag 参数是通过 log.New 创建记录器的标志集
	l, err := syslog.NewLogger(syslog.LOG_DEBUG, log.Ldate|log.Ltime|log.Lshortfile)
	if err != nil {
		log.Fatal(err)
	}
	l.Fatal("退出")
}