package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// signal包实现了对输入信号的访问
func main() {

	// 初始化一个信号channel
	c := make(chan os.Signal, 1)

	// 让signal包将输入信号转发到c。如果没有列出要传递的信号，会将所有输入信号传递到c；否则只传递列出的输入信号
	// signal包不会为了向c发送信息而阻塞（就是说如果发送时c阻塞了，signal包会直接放弃）
	// 调用者应该保证c有足够的缓存空间可以跟上期望的信号频率。对使用单一信号用于通知的通道，缓存为1就足够了
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL)

	s := <-c
	fmt.Println("Got signal:", s)

	// 让signal包停止向c转发信号
	// 它会取消之前使用c调用的所有Notify的效果。当Stop返回后，会保证c不再接收到任何信号
	signal.Stop(c)
}