package main

import (
	"io"
	"log"
	"net"
	"os"
)

// net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket
// 虽然本包提供了对网络原语的访问，大部分使用者只需要Dial、Listen和Accept函数提供的基本接口；以及相关的Conn和Listener接口
// crypto/tls包提供了相同的接口和类似的Dial和Listen函数
func main() {

	// 创建服务
	exampleServer()

	// 客户端请求
	exampleClient()
}

func exampleServer() {

	// 创建tcp监听
	// network必须是"tcp", "tcp4", "tcp6", "unix", "unixpacket".
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// 等待连接
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {

			// 拷贝接收内容并输出
			io.Copy(os.Stdout, c)

			// 关闭链接
			c.Close()
		}(conn)
	}
}

func exampleClient() {

	// 创建tcp连接
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// 写入内容
	if _, err := conn.Write([]byte("Hello World!")); err != nil {
		log.Fatal(err)
	}

	// 关闭连接
	if err := conn.Close(); err != nil {
		log.Fatal(err)
	}
}
