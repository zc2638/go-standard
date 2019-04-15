package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
)

// fcgi包实现了FastCGI协议。目前只支持响应器的角色。
// 协议定义的地址：http://www.fastcgi.com/drupal/node/6?q=node/22
func main() {

	// 服务监听
	listener, err:= net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}

	// 初始化一个自定义handler
	srv := new(FastCGIHandler)

	// 接受从监视器l传入的FastCGI连接，为每一个FastCGI连接创建一个新的go程，该go程读取请求然后调用handler回复请求
	// 如果l是nil，Serve将从os.Stdin接受连接
	// 如果handler是nil，将使用http.DefaultServeMux
	if err := fcgi.Serve(listener, srv); err != nil {
		log.Fatal(err)
	}
}

type FastCGIHandler struct{}

func (s FastCGIHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	fmt.Println(123)
	resp.Write([]byte("<h1>Hello, World!</h1>\n<p>Hello Gopher.</p>"))
}

/*
	需要nginx配置fastcgi
	server {
        listen 80;
        server_name go.dev;
        root /root/go/src/godev;
        index index.html;
        #gzip off;
        #proxy_buffering off;

        location / {
                 try_files $uri $uri/;
        }

        location ~ /.* {
                include         fastcgi.conf;
                fastcgi_pass    127.0.0.1:8081;
        }

        try_files $uri $uri.html =404;
	}
 */