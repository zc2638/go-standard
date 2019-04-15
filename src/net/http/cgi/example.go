package main

import (
	"log"
	"net/http"
	"net/http/cgi"
)

// cgi包实现了CGI（通用网关接口）
// 请注意，使用CGI意味着启动一个新的进程来处理每个请求，这通常比使用长时间运行的服务器效率低
// 此软件包主要用于与现有系统兼容

// CGI(Common Gateway Interface)是能让web服务器和CGI脚本共同处理客户的请求的协议。它的协议定义文档是http://www.ietf.org/rfc/rfc3875
// 其中Web服务器负责管理连接，数据传输，网络交互等。至于CGI脚本就负责管理具体的业务逻辑
// Web服务器的功能是将客户端请求（HTTP Request）转换成CGI脚本请求，然后执行脚本，接着将CGI脚本回复转换为客户端的回复（HTTP Response）
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// 初始化一个cgi.Handler，用于在子进程中执行具有一个CGI环境的可执行程序
		handler := new(cgi.Handler)

		// 设置CGI可执行文件的路径
		handler.Path = "/usr/local/opt/go/libexec/bin/go"

		// Dir指定CGI程序的工作目录
		// 如果Dir为""则使用Path的基目录；如果Path没有基目录则使用当前工作目录
		handler.Dir = "/Users/zc/go/project/standard-library/"

		// 指定文件路径
		script := "testdata/cgi/" + r.URL.Path

		// 设置进程参数
		args := []string{"run", script}

		// 可选的传递给子进程的参数
		handler.Args = append(handler.Args, args...)

		// 额外设置的环境变量（如果有），格式为"key=value"
		handler.Env = append(handler.Env, "GOPATH=/Users/zc/go", "GOROOT=/usr/local/opt/go/libexec")

		// 从host继承的环境变量，只有"key"
		handler.InheritEnv = []string{"HOME", "GOCACHE"}

		// 可选的logger接口切片，如为nil则使用log.Print
		handler.Logger = nil

		// handler的根URI前缀，""代表"/"
		handler.Root = ""

		handler.ServeHTTP(w, r)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}