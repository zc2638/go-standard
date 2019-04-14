package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// http包提供了HTTP客户端和服务端的实现
func main() {

	// 指定func注册hander到DefaultServeMux(包含Hijacker接口用法)
	exampleHandleFunc()

	// 注册handler到DefaultServeMux
	exampleHandler()

	// Request对象示例
	exampleRequest()

	// Response对象示例
	exampleResponse()

	// http curl示例
	exampleHttp()

	// 客户端请求
	exampleClient()

	// 静态文件服务监听
	exampleFileServer()

	// 使用ServeMux创建服务
	exampleServeMux()

	// 服务端创建
	exampleServer()

	// 服务器监听返回html内容
	exampleServerHtml()

	// transport代理
	exampleTransport()
}

type CustomHandler struct{}

func (CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 限制Body的大小
	r.Body = http.MaxBytesReader(w, r.Body, 10485760) // 10M

	if r.Method == "GET" {
		w.Write([]byte("Hello World!"))
	} else {
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
	return
}

func exampleHandler() {

	// 返回一个简单的请求处理器，该处理器会对每个请求都回复"404 page not found"
	notFound := http.NotFoundHandler()

	// 返回一个请求处理器，该处理器会对每个请求都使用状态码code重定向到网址url
	redirect := http.RedirectHandler("http://www.baidu.com", http.StatusFound)

	// 返回一个采用指定时间限制的请求处理器
	// 返回的Handler会调用h.ServeHTTP去处理每个请求，但如果某一次调用耗时超过了时间限制，该处理器会回复请求状态码503 Service Unavailable，并将msg作为回复的主体（如果msg为空字符串，将发送一个合理的默认信息）
	// 在超时后，h对它的ResponseWriter接口参数的写入操作会返回ErrHandlerTimeout
	timeout := http.TimeoutHandler(redirect, time.Minute*2, "")

	// 返回一个处理器，该处理器会将请求的URL.Path字段中给定前缀prefix去除后再交由h处理
	// StripPrefix会向URL.Path字段中没有给定前缀的请求回复404 page not found
	stripPrefix := http.StripPrefix("/temp/", redirect)

	// 自定义handler
	custom := CustomHandler{}

	// 注册HTTP处理器handler和对应的模式pattern（注册到DefaultServeMux）
	// 如果该模式已经注册有一个处理器，Handle会panic
	http.Handle("/404", notFound)
	http.Handle("/redirect", redirect)
	http.Handle("/timeout", timeout)
	http.Handle("/stripPrefix", stripPrefix)
	http.Handle("/custom", custom)
}

func exampleHandleFunc() {

	// 注册一个处理器函数handler和对应的模式pattern（注册到DefaultServeMux）
	http.HandleFunc("/hijack", func(w http.ResponseWriter, r *http.Request) {

		// 指定Hijacker接口
		// HTTP处理器ResponseWriter接口参数的下层如果实现了Hijacker接口，可以让HTTP处理器接管该连接
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
			return
		}

		// Hijack让调用者接管连接，返回连接和关联到该连接的一个缓冲读写器
		// 调用本方法后，HTTP服务端将不再对连接进行任何操作
		// 调用者有责任管理、关闭返回的连接
		conn, buf, err := hj.Hijack()
		if err != nil {

			// 用指定的错误信息和状态码回复请求，将数据写入w。错误信息必须是明文
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 关闭连接
		defer conn.Close()

		// 写入
		if _, err := buf.WriteString("Now we're speaking raw TCP. Say hi: "); err != nil {
			log.Printf("error write string: %v", err)
			return
		}

		// 将所有缓冲数据写入
		if err := buf.Flush(); err != nil {
			log.Printf("error flush: %v", err)
			return
		}

		// 读取直到输入中第一次出现delim，返回一个包含分隔符之前的数据的字符串
		// 如果在查找分隔符之前遇到错误，则返回错误之前读取的数据和错误本身（通常为io.EOF）
		s, err := buf.ReadString('\n')
		if err != nil {
			log.Printf("error reading string: %v", err)
			return
		}

		fmt.Fprintf(buf, "You said: %q\nBye.\n", s)
		if err := buf.Flush(); err != nil {
			log.Printf("error flush: %v", err)
			return
		}
	})
}

func exampleRequest() {

	body := "a=1&b=2&c=3"

	// 解析url
	urls, err := url.Parse("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	// body转io.ReaderCloser
	rc := func(reader io.Reader) io.ReadCloser {
		rcr, ok := reader.(io.ReadCloser)
		if !ok && reader != nil {
			rcr = ioutil.NopCloser(reader)
		}
		return rcr
	}(strings.NewReader(body))

	req := &http.Request{
		//  Method指定HTTP方法（GET、POST、PUT等）。对客户端，""代表GET
		Method: "GET",

		// URL在服务端表示被请求的URI，在客户端表示要访问的URL。
		// 在服务端，URL字段是解析请求行的URI（保存在RequestURI字段）得到的，
		// 对大多数请求来说，除了Path和RawQuery之外的字段都是空字符串
		// 在客户端，URL的Host字段指定了要连接的服务器，而Request的Host字段（可选地）指定要发送的HTTP请求的Host头的值
		URL: urls,

		// 接收到的请求的协议版本。本包生产的Request总是使用HTTP/1.1
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,

		/*
			Header字段用来表示HTTP请求的头域。如果头域（多行键值对格式）为：
				accept-encoding: gzip, deflate
				Accept-Language: en-us
				Connection: keep-alive
			则：
				Header = map[string][]string{
					"Accept-Encoding": {"gzip, deflate"},
					"Accept-Language": {"en-us"},
					"Connection": {"keep-alive"},
				}
			HTTP规定头域的键名（头名）是大小写敏感的，请求的解析器通过规范化头域的键名来实现这点。
			在客户端的请求，可能会被自动添加或重写Header中的特定的头，参见Request.Write方法
		*/
		Header: make(http.Header),

		// Body是请求的主体
		// 在客户端，如果Body是nil表示该请求没有主体，相当于GET请求
		// Client的Transport字段会负责调用Body的Close方法。
		// 在服务端，Body字段总是非nil的；但在没有主体时，读取Body会立刻返回EOF。
		// Server会关闭请求的主体，ServeHTTP处理器不需要关闭Body字段
		Body: rc,

		// 定义一个可选的func以返回body的新副本，使用GetBody仍然需要设置body
		// 当重定向需要多次读取主体时，它用于客户端请求
		GetBody: func() (closer io.ReadCloser, e error) {
			return rc, nil
		},

		// ContentLength记录相关内容的长度。
		// 如果为-1，表示长度未知，如果>=0，表示可以从Body字段读取ContentLength字节数据
		// 在客户端，如果Body非nil而该字段为0，表示不知道Body的长度
		ContentLength: 0,

		// Close在服务端指定是否在回复请求后关闭连接，在客户端指定是否在发送请求后关闭连接
		Close: false,

		// 在服务端，Host指定URL会在其上寻找资源的主机
		// 该值可以是Host头的值，或者URL自身提供的主机名
		// Host的格式可以是"host:port"
		// 在客户端，请求的Host字段（可选地）用来重写请求的Host头
		// 如过该字段为""，Request.Write方法会使用URL字段的Host
		Host: urls.Host,

		// Form是解析好的表单数据，包括URL字段的query参数和POST或PUT的表单数据
		// 本字段只有在调用ParseForm后才有效
		// 在客户端，会忽略请求中的本字段而使用Body替代
		Form: nil,

		// PostForm是解析好的POST或PUT的表单数据
		// 本字段只有在调用ParseForm后才有效
		// 在客户端，会忽略请求中的本字段而使用Body替代
		PostForm: nil,

		// MultipartForm是解析好的多部件表单，包括上传的文件
		// 本字段只有在调用ParseMultipartForm后才有效
		// 在客户端，会忽略请求中的本字段而使用Body替代
		MultipartForm: nil,

		// RemoteAddr允许HTTP服务器和其他软件记录该请求的来源地址，一般用于日志
		// 本字段不是ReadRequest函数填写的，也没有定义格式
		// 本包的HTTP服务器会在调用处理器之前设置RemoteAddr为"IP:port"格式的地址
		// 客户端会忽略请求中的RemoteAddr字段
		RemoteAddr: "",

		// RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
		// 一般应使用URI字段，在客户端设置请求的本字段会导致错误
		RequestURI: "",

		// TLS字段允许HTTP服务器和其他软件记录接收到该请求的TLS连接的信息
		// 本字段不是ReadRequest函数填写的
		// 对启用了TLS的连接，本包的HTTP服务器会在调用处理器之前设置TLS字段，否则将设TLS为nil。
		// 客户端会忽略请求中的TLS字段
		TLS: nil,
	}

	// 判断该请求使用的HTTP协议版本至少是major.minor
	req.ProtoAtLeast(1, 1)

	// 返回请求中的客户端用户代理信息（请求的User-Agent头）
	req.UserAgent()

	// 返回请求中的访问来路信息。（请求的Referer头）
	// Referer在请求中就是拼错了的，这是HTTP早期就有的错误
	// 该值也可以从用Header["Referer"]获取；让获取Referer字段变成方法的好处是，编译器可以诊断使用正确单词拼法的req.Referrer()的程序，但却不能诊断使用Header["Referrer"]的程序
	req.Referer()

	// 向请求中添加一个cookie
	// 所有的cookie都写在同一行，用分号分隔（cookie内部用逗号分隔属性）
	req.AddCookie(&http.Cookie{
		Name:  "GoSession",
		Value: "dasghdjashfjgasidhqwi321u34h12bdusag",
	})

	// 使用提供的用户名和密码，采用HTTP基本认证，设置请求的Authorization头。HTTP基本认证会明码传送用户名和密码
	req.SetBasicAuth("gopher", "123123")

	var buf bytes.Buffer
	// 以有线格式将HTTP/1.1请求写入w（用于将请求写入下层TCPConn等）
	// 本方法会考虑请求的如下字段：Host、URL、Method (defaults to "GET")、Header、ContentLength、TransferEncoding、Body
	// 如果存在Body，ContentLength字段<= 0且TransferEncoding字段未显式设置为["identity"]，Write方法会显式添加"Transfer-Encoding: chunked"到请求的头域。Body字段会在发送完请求后关闭
	req.Write(&buf)

	buf.Reset()
	// 类似Write但会将请求以HTTP代理期望的格式发送
	// 会使用绝对URI（包括协议和主机名）来初始化请求的第1行（Request-URI行）。无论何种情况，WriteProxy都会使用r.Host或r.URL.Host设置Host头
	req.WriteProxy(&buf)

	// 解析并返回该请求的Cookie头设置的cookie
	req.Cookies()

	// 返回请求中名为name的cookie，如果未找到该cookie会返回nil, ErrNoCookie
	req.Cookie("GoSession")

	// 解析URL中的查询字符串，并将解析结果更新到r.Form字段
	// 对于POST或PUT请求，ParseForm还会将body当作表单解析，并将结果既更新到r.PostForm也更新到r.Form
	// 解析结果中，POST或PUT请求主体要优先于URL查询字符串（同名变量，主体的值在查询字符串的值前面）
	// 如果请求的主体的大小没有被MaxBytesReader函数设定限制，其大小默认限制为开头10MB
	// ParseMultipartForm会自动调用ParseForm。重复调用本方法是无意义的
	req.ParseForm()

	// 将请求的主体作为multipart/form-data解析
	// 请求的整个主体都会被解析，得到的文件记录最多maxMemery字节保存在内存，其余部分保存在硬盘的temp文件里
	// 如果必要，ParseMultipartForm会自行调用ParseForm。重复调用本方法是无意义的
	req.ParseMultipartForm(10485760) // 10M

	// 返回key为键查询r.Form字段得到结果[]string切片的第一个值
	// POST和PUT主体中的同名参数优先于URL查询字符串
	// 如果必要，本函数会隐式调用ParseMultipartForm和ParseForm
	req.FormValue("a")

	// 返回key为键查询r.PostForm字段得到结果[]string切片的第一个值
	// 如果必要，本函数会隐式调用ParseMultipartForm和ParseForm
	req.PostFormValue("a")

	// 返回以key为键查询r.MultipartForm字段得到结果中的第一个文件和它的信息
	// 如果必要，本函数会隐式调用ParseMultipartForm和ParseForm。查询失败会返回ErrMissingFile错误
	req.FormFile("b")

	// 如果请求是multipart/form-data POST请求，MultipartReader返回一个multipart.Reader接口，否则返回nil和一个错误
	// 使用本函数代替ParseMultipartForm，可以将r.Body作为流处理
	req.MultipartReader()
}

func exampleResponse() {

	resp := &http.Response{
		// 响应状态，例如"200 OK"
		Status: "200 OK",

		// 响应状态码，例如200
		StatusCode: http.StatusOK,

		// 协议版本
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 1,

		// Header保管头域的键值对。
		// 如果回复中有多个头的键相同，Header中保存为该键对应用逗号分隔串联起来的这些头的值
		// 被本结构体中的其他字段复制保管的头（如ContentLength）会从Header中删掉
		// Header中的键都是规范化的，参见CanonicalHeaderKey函数
		Header: make(http.Header),

		// Body代表回复的主体
		// Client类型和Transport类型会保证Body字段总是非nil的，即使回复没有主体或主体长度为0
		// 关闭主体是调用者的责任
		// 如果服务端采用"chunked"传输编码发送的回复，Body字段会自动进行解码
		Body: nil,

		// ContentLength记录相关内容的长度
		// 其值为-1表示长度未知（采用chunked传输编码）
		// 除非对应的Request.Method是"HEAD"，否则其值>=0表示可以从Body读取的字节数
		ContentLength: -1,

		// Close记录头域是否指定应在读取完主体后关闭连接。（即Connection头）
		// 该值是给客户端的建议，Response.Write方法的ReadResponse函数都不会关闭连接
		Close: false,

		// Trailer字段保存和头域相同格式的trailer键值对，和Header字段相同类型
		Trailer: make(http.Header),

		// Request是用来获取此回复的请求
		// Request的Body字段是nil（因为已经被用掉了）
		// 这个字段是被Client类型发出请求并获得回复后填充的
		Request: nil,

		// TLS包含接收到该回复的TLS连接的信息。 对未加密的回复，本字段为nil
		// 返回的指针是被（同一TLS连接接收到的）回复共享的，不应被修改
		TLS: nil,
	}

	// 判断该回复使用的HTTP协议版本至少是major.minor
	resp.ProtoAtLeast(1, 1)

	// 解析并返回该回复中的Set-Cookie头设置的cookie
	resp.Cookies()

	// 返回该回复的Location头设置的URL
	// 相对地址的重定向会相对于该回复对应的请求来确定绝对地址
	// 如果回复中没有Location头，会返回nil, ErrNoLocation
	resp.Location()

	var buf bytes.Buffer
	// 以有线格式将回复写入w（用于将回复写入下层TCPConn等）
	// 本方法会考虑如下字段：StatusCode、ProtoMajor、ProtoMinor、Request.Method、TransferEncoding、Trailer、Body、ContentLength、Header（不规范的键名和它对应的值会导致不可预知的行为）
	// Body字段在发送完回复后会被关闭
	resp.Write(&buf)
}

func exampleHttp() {

	var uri = "http://www.google.com/robots.txt"
	var body = "a=1&b=2&c=3"
	values, _ := url.ParseQuery(body)

	// GET请求
	http.Get(uri)

	// POST请求，可指定ContentType
	http.Post(uri, "application/x-www-form-urlencoded", strings.NewReader(body))

	// POST请求，默认使用application/x-www-form-urlencoded
	http.PostForm(uri, values)

	// Head请求
	http.Head(uri)
}

func exampleClient() {

	body := "a=1&b=2&c=3"

	// 使用指定的方法、网址和可选的主题创建并返回一个新的*http.Request请求
	// 如果body参数实现了io.Closer接口，Request返回值的Body 字段会被设置为body，并会被Client类型的Do、Post和PostFOrm方法以及Transport.RoundTrip方法关闭
	// 可以直接使用http.Request
	req, err := http.NewRequest("GET", "http://www.google.com/robots.txt", strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	// 初始化http客户端
	// 等于直接使用&http.Client{}
	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// 关闭body
	defer resp.Body.Close()

	// 读取body内容
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}

func exampleFileServer() {

	// 指定目录。空Dir被视为"."，即代表当前目录
	dir := http.Dir("/tmp")

	// 返回一个使用FileSystem接口root提供文件访问服务的HTTP处理器
	fileServer := http.FileServer(dir)

	// 返回一个处理器，该处理器会将请求的URL.Path字段中给定前缀prefix去除后再交由h处理
	// StripPrefix会向URL.Path字段中没有给定前缀的请求回复404 page not found
	fileHandler := http.StripPrefix("/tmpfiles/", fileServer)

	// 注册HTTP处理器handler和对应的模式pattern（注册到DefaultServeMux）
	// 如果该模式已经注册有一个处理器，Handle会panic
	http.Handle("/tmpfiles/", fileHandler)

	// 监听TCP地址addr，并且会使用handler参数调用Serve函数处理接收到的连接
	// handler参数一般会设为nil，此时会使用DefaultServeMux
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func exampleServeMux() {

	// 创建并返回一个新的*http.ServeMux
	// ServeMux类型是HTTP请求的多路转接器。它会将每一个接收的请求的URL与一个注册模式的列表进行匹配，并调用和URL最匹配的模式的处理器
	// 模式是固定的、由根开始的路径，如"/favicon.ico"，或由根开始的子树，如"/images/"（注意结尾的斜杠）。较长的模式优先于较短的模式，因此如果模式"/images/"和"/images/thumbnails/"都注册了处理器，后一个处理器会用于路径以"/images/thumbnails/"开始的请求，前一个处理器会接收到其余的路径在"/images/"子树下的请求
	// 注意，因为以斜杠结尾的模式代表一个由根开始的子树，模式"/"会匹配所有的未被其他注册的模式匹配的路径，而不仅仅是路径"/"
	// 模式也能（可选地）以主机名开始，表示只匹配该主机上的路径。指定主机的模式优先于一般的模式，因此一个注册了两个模式"/codesearch"和"codesearch.google.com/"的处理器不会接管目标为"http://www.google.com/"的请求。
	// ServeMux还会注意到请求的URL路径的无害化，将任何路径中包含"."或".."元素的请求重定向到等价的没有这两种元素的URL
	mux := http.NewServeMux()

	mux.Handle("/custom/", CustomHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	// 监听srv.Addr确定的TCP地址，并且会调用Serve方法处理接收到的连接
	// 必须提供证书文件和对应的私钥文件
	// 如果证书是由权威机构签发的，certFile参数必须是顺序串联的服务端证书和CA证书
	// 如果srv.Addr为空字符串，会使用":https"
	if err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", mux); err != nil {
		log.Fatal(err)
	}
}

func exampleServer() {

	// 初始化server，定义了运行HTTP服务端的参数。Server的零值是合法的配置
	srv := &http.Server{
		// 监听的TCP地址，如果为空字符串会使用":http"
		Addr: ":8080",

		// 调用的处理器，如为nil会调用http.DefaultServeMux
		Handler: nil,

		// 请求的读取操作在超时前的最大持续时间
		ReadTimeout: 0,

		// 回复的写入操作在超时前的最大持续时间
		WriteTimeout: 0,

		// 允许读取请求头的时间。连接的读取截止时间在读取头之后重置，处理程序可以决定什么对主体来说太慢
		ReadHeaderTimeout: 0,

		// 启用keep alives时等待下一个请求的最大时间
		// 如果idleTimeout为0，则使用readTimeout的值。如果两者都为零，则使用readHeaderTimeout
		IdleTimeout: 0,

		// 请求的头域最大长度，如为0则用DefaultMaxHeaderBytes
		MaxHeaderBytes: 0,

		// 可选的TLS配置，用于ListenAndServeTLS方法
		TLSConfig: &tls.Config{},

		// 指定一个可选的回调函数，该函数会在一个与客户端的连接改变状态时被调用
		ConnState: func(conn net.Conn, state http.ConnState) {
			if err := conn.Close(); err != nil {
				log.Fatal(err)
			}
		},

		// ErrorLog指定一个可选的日志记录器，用于记录接收连接时的错误和处理器不正常的行为
		// 如果本字段为nil，日志会通过log包的标准日志记录器写入os.Stderr
		ErrorLog: nil,
	}

	// 控制是否允许HTTP闲置连接重用（keep-alive）功能，默认该功能总是被启用的
	// 只有资源非常紧张的环境或者服务端在关闭进程中时，才应该关闭该功能
	srv.SetKeepAlivesEnabled(true)

	// 创建TCP服务监听
	l, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}

	// 接手监听器l收到的每一个连接，并为每一个连接创建一个新的服务go程
	// 该go程会读取请求，然后调用srv.Handler回复请求
	if err := srv.Serve(l); err != nil {
		log.Fatal(err)
	}

	// 接手l上的传入连接，为每个连接创建一个新的服务例程。服务程序读取请求，然后调用srv.Handler来回复它们
	// 此外，如果服务器的TLSConfig.Certificates和TLSConfig.GetCertificate都未被填充，则必须提供包含服务器的证书和匹配私钥的文件
	// 如果证书由证书颁发机构签署，则certFile应该是服务器的串联证书，任何中间件和CA的证书
	// 对于HTTP/2支持，在调用Serve之前，应将srv.TLSConfig初始化为提供的侦听器的TLS配置
	// 如果srv.TLSConfig非零，并且在Config.NextProtos中不包含字符串“h2”，则不启用HTTP/2支持
	if err := srv.ServeTLS(l, "cert.pem", "key.pem"); err != nil {
		log.Fatal(err)
	}

	// 监听srv.Addr指定的TCP地址，并且会调用Serve方法接收到的连接
	// 如果srv.Addr为空字符串，会使用":http"
	// 效果等同于http.ListenAndServe()
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	// 监听srv.Addr确定的TCP地址，并且会调用Serve方法处理接收到的连接
	// 必须提供证书文件和对应的私钥文件
	// 如果证书是由权威机构签发的，certFile参数必须是顺序串联的服务端证书和CA证书
	// 如果srv.Addr为空字符串，会使用":https"
	// 效果等同于http.ListenAndServeTLS
	if err := srv.ListenAndServeTLS("cert.pem", "key.pem"); err != nil {
		log.Fatal(err)
	}

	// 正常关闭服务器而不中断任何活动连接
	// 首先关闭所有打开的监听程序，然后关闭所有空闲连接，然后无限期地等待连接返回到空闲状态然后关闭
	// 如果提供的Context在关闭完成之前到期，Shutdown将返回Context的错误，否则返回关闭服务器的底层侦听器返回的任何错误
	// 当调用Shutdown时，Serve，ListenAndServe和ListenAndServeTLS立即返回ErrServerClosed。确保程序不会退出，而是等待Shutdown返回
	// 不会尝试关闭或等待被劫持的连接，如WebSockets。如果需要，Shutdown的调用者应该分别通知关闭的这种长期连接并等待
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}

	// 注册一个函数来调用Shutdown
	// 这可以用于正常关闭经过NPN/ALPN协议升级或已被劫持的连接
	// 此功能应启动特定于协议的正常关机，但不应等待关机完成
	srv.RegisterOnShutdown(func() {
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	})

	// 立即关闭所有活动的net.listener和state statenew、state active或stateidle中的任何连接
	// 要正常关机，请使用Shutdown
	// 不试图关闭（甚至不知道）任何被劫持的连接，如websockets。
	// 返回关闭服务器的基础侦听器返回的任何错误
	if err := srv.Close(); err != nil {
		log.Fatal(err)
	}
}

func exampleServerHtml() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<p>Hello World!</p>"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func exampleTransport() {

	uri, _ := url.Parse("http://www.baidu.com")

	// 实现了RoundTripper接口，支持http、https和http/https代理。Transport类型可以缓存连接以在未来重用
	// RoundTripper接口是具有执行单次HTTP事务的能力（接收指定请求的回复）的接口。
	// RoundTripper接口的类型必须可以安全的被多线程同时使用
	transport := &http.Transport{

		// 指定一个对给定请求返回代理的函数
		// 如果该函数返回了非nil的错误值，请求的执行就会中断并返回该错误。
		// 如果Proxy为nil或返回nil的*URL置，将不使用代理
		Proxy: http.ProxyURL(uri),

		// 指定一个Dial函数，用于创建未加密TCP连接
		// 如果DialContext为nil（下面不推荐使用的Dial也为nil），则使用net.Dial
		DialContext: nil,

		// 指定一个可选的Dial函数，用于为非代理的https请求创建TLS连接
		DialTLS: nil,

		// 指定用于tls.Client的TLS配置信息
		// 如果该字段为nil，会使用默认的配置信息
		TLSClientConfig: nil,

		// 指定等待TLS握手完成的最长时间。零值表示不设置超时
		TLSHandshakeTimeout: 0,

		// 如果DisableKeepAlives为true，会禁止不同HTTP请求之间TCP连接的重用
		DisableKeepAlives: false,

		// 如果DisableCompression为true，会禁止Transport在请求中没有Accept-Encoding头时，主动添加"Accept-Encoding: gzip"头，以获取压缩数据
		// 如果Transport自己请求gzip并得到了压缩后的回复，它会主动解压缩回复的主体
		// 但如果用户显式的请求gzip压缩数据，Transport是不会主动解压缩的
		DisableCompression: false,

		// 控制所有主机的最大空闲（保持活动）连接数。零表示无限制
		MaxIdleConns: 0,

		// 如果MaxIdleConnsPerHost!=0，会控制每个主机下的最大闲置连接
		// 如果MaxIdleConnsPerHost==0，会使用DefaultMaxIdleConnsPerHost
		MaxIdleConnsPerHost: 0,

		// 可以选择限制每个主机的连接总数，包括拨号、活动和空闲状态的连接
		// 如果违反限制，拨号将被阻止。零值无限制
		MaxConnsPerHost: 0,

		// 空闲（保持活动）连接在关闭之前保持空闲状态的最长时间。零值无限制
		IdleConnTimeout: 0,

		// 指定在发送完请求（包括其可能的主体）之后，等待接收服务端的回复的头域的最大时间
		// 零值表示不设置超时。该时间不包括获取回复主体的时间
		ResponseHeaderTimeout: 0,

		// 指定在连接请求期间发送到代理的头(可选)
		ProxyConnectHeader: nil,

		// 指定对服务器响应头中允许的响应字节数的限制。零值无限制
		MaxResponseHeaderBytes: 0,
	}

	// 注册一个新的名为scheme的协议。transport会将使用scheme协议的请求转交给rt。rt有责任模拟HTTP请求的语义。
	// RegisterProtocol可以被其他包用于提供"ftp"或"file"等协议的实现
	transport.RegisterProtocol("test", http.NewFileTransport(http.Dir("/tmp/")))

	// 关闭所有之前的请求建立但目前处于闲置状态的连接。本方法不会中断正在使用的连接
	transport.CloseIdleConnections()

	// 创建一个*http.Request
	req, err := http.NewRequest("GET", "http://www.google.com/robots.txt", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 高层次的HTTP客户端支持（如管理cookie和重定向）请参见Get、Post等函数和Client类型
	transport.RoundTrip(req)
}