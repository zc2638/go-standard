package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"
)

// rpc包提供了通过网络或其他I/O连接对一个对象的导出方法的访问
// 服务端注册一个对象，使它作为一个服务被暴露，服务的名字是该对象的类型名
// 注册之后，对象的导出方法就可以被远程访问
// 服务端可以注册多个不同类型的对象（服务），但注册具有相同类型的多个对象是错误的
func main() {

	// rpc服务端到客户端的完整示例
	example()
	example2()
}

type Hello struct {}

func (h *Hello) Say(args *[]string, reply *string) error {

	*reply = strings.Join(*args, " ")
	return nil
}

func example() {

	// 初始化Hello
	hello := new(Hello)

	// 初始化服务端
	// 创建并返回一个*rpc.Server
	server := rpc.NewServer()

	// 在server注册并公布rcvr的方法集
	// 满足: 方法是导出的；方法有两个参数，都是导出类型或内建类型；方法的第二个参数是指针；方法只有一个error接口类型的返回值
	server.Register(hello)

	// 类似Register，但使用提供的name代替rcvr的具体类型名作为服务名
	server.RegisterName("Hello", hello)

	// 监听端口
	l, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("net.Listen tcp :0: %v", err)
	}

	// 接收监听器l获取的连接，然后服务每一个连接。Accept会阻塞，调用者应另开线程："go server.Accept(l)"
	go server.Accept(l)

	// 注册server的RPC信息HTTP处理器对应到rpcPath，注册server的debug信息HTTP处理器对应到debugPath
	// HandleHTTP会注册到http.DefaultServeMux。之后，仍需要调用http.Serve()，一般会另开线程："go http.Serve(l, nil)"
	server.HandleHTTP("/hello", "/debug")

	// 将addr作为TCP地址解析并返回
	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("ResolveTCPAddr error: ", err)
	}

	// 在网络协议net上连接本地地址laddr和远端地址raddr
	// net必须是"tcp"、"tcp4"、"tcp6"；如果laddr不是nil，将使用它作为本地地址，否则自动选择一个本地地址
	conn, _ := net.DialTCP("tcp", nil, address)
	defer conn.Close()

	// 初始化客户端
	// 返回一个新的rpc.Client，以管理对连接另一端的服务的请求。它添加缓冲到连接的写入侧，以便将回复的头域和有效负载作为一个单元发送
	client := rpc.NewClient(conn)
	defer client.Close()

	// 设置参数
	args := &[]string{"Hello", "World!"}
	// 初始化接收
	reply := new(string)
	err = client.Call("Hello.Say", args, reply)
	if err != nil {
		log.Fatal("Hello error:", err)
	}
	log.Println(*reply)
}

func example2() {

	hello := new(Hello)
	rpc.Register(hello)
	rpc.HandleHTTP()

	// 设置服务端监听
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	// 客户端连接服务端
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &[]string{"Hello", "Gopher!"}
	reply := new(string)
	err = client.Call("Hello.Say", args, reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(*reply)
}