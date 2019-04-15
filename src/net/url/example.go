package main

import (
	"fmt"
	"log"
	"net/url"
)

// url包解析URL并实现了查询的逸码
func main() {

	// url编码/解码
	example()

	// 创建/解析url
	exampleURL()

	// 创建/解析url参数
	exampleQuery()

	// 创建url用户名和密码信息
	exampleUser()
}

func example() {

	s := "https://www.baidu.com?test=1&name=测试"

	// url编码
	se := url.QueryEscape(s)
	fmt.Println("urlEncode:", se)

	// url解码
	sd, err := url.QueryUnescape(se)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("urlDecode:", sd)
}

func exampleURL() {

	// 创建URL，结构体各参数意义参见 解析URL部分
	//ul := url.URL{}

	// 解析URL，本函数会假设rawurl是在一个HTTP请求里，因此会假设该参数是一个绝对URL或者绝对路径，并会假设该URL没有#fragment后缀。（网页浏览器会在去掉该后缀后才将网址发送到网页服务器）
	//url.ParseRequestURI("http://bing.com/search?q=dotnet")

	// 解析URL，rawurl可以是绝对地址，也可以是相对地址
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}

	// 协议名称http或https
	fmt.Println("Scheme:", u.Scheme)

	// 域名加端口
	fmt.Println("Host:", u.Host)

	// 编码后的不透明数据
	fmt.Println("Opaque", u.Opaque)

	// 用户名和密码信息
	fmt.Println("User:", u.User)

	// path路径
	fmt.Println("Path:", u.Path)

	// 编码后的查询字符串，没有'?'
	fmt.Println("RawQuery:", u.RawQuery)

	// 引用的片段（文档位置），没有'#'
	fmt.Println("Fragment:", u.Fragment)

	u.Scheme = "https"
	u.Host = "google.com"

	// 判断url是否绝对地址
	fmt.Println("IsAbs:", u.IsAbs())

	// 域名
	fmt.Println("Hostname:", u.Hostname())

	// 端口
	fmt.Println("Port:", u.Port())

	// 转字符串显示url
	fmt.Println("String:", u.String())

	// 返回 u.Path 的转义形式
	fmt.Println("EscapedPath:", u.EscapedPath())

	// 返回在 u 的 HTTP 请求中使用的编码 path?query 或 opaque?query 字符串
	fmt.Println("RequestURI:", u.RequestURI())

	text, err := u.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("MarshalBinary: %s\n", string(text))

	if err := u.UnmarshalBinary(text); err != nil {
		log.Fatal(err)
	}
	fmt.Println("UnmarshalBinary: ok")

	// 解析RawQuery字段并返回其表示的url.Values类型键值对
	q := u.Query()

	// 将key对应的值集设为value，它会替换掉已有的值集
	q.Set("q", "golang")

	// 将value添加到key关联的值集里原有的值的后面
	q.Add("say", "hello")

	// 获取key对应的值集的第一个值。如果没有对应key的值集会返回空字符串
	fmt.Println(q.Get("say"))

	// 将v编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序
	fmt.Println(q.Encode())

	// 删除key关联的值集
	q.Del("say")

	fmt.Println(u)

	// 解析一个相对路径的url
	uri, err := url.Parse("../../..//search?q=test")
	if err != nil {
		log.Fatal(err)
	}
	// 根据一个绝对URI将一个URI补全为一个绝对URI
	// 参数ref可以是绝对URI或者相对URI，返回一个新的URL实例，即使该实例和u或者ref完全一样
	// 如果ref是绝对URI，本方法会忽略参照URI并返回ref的一个拷贝
	uriNew := u.ResolveReference(uri)
	fmt.Println(uriNew.String())
}

func exampleQuery() {

	// 解析url参数
	// 解析一个URL编码的查询字符串，并返回可以表示该查询的Values类型的字典
	// 本函数总是返回一个包含了所有合法查询参数的非nil字典，err用来描述解码时遇到的（如果有）第一个错误
	m, err := url.ParseQuery(`x=1&y=2&y=3;z`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)

	// 创建url参数, 可以直接用url.Value结构
	params := url.Values{}
	// 设置参数对
	params.Set("say", "hello")
	params.Set("test", "123")
	params.Set("q", "Go")
	fmt.Println(params.Encode())
}

func exampleUser() {

	// 返回一个用户名设置为username的不设置密码的*url.Userinfo
	u := url.User("gopher")

	// 返回string格式
	fmt.Println(u.String())

	// 返回用户名
	fmt.Println(u.Username())

	// 如果设置了密码，返回密码和true，否则会返回false
	fmt.Println(u.Password())

	// 返回一个用户名设置为username、密码设置为password的*url.Userinfo
	up := url.UserPassword("gopher", "123123")
	fmt.Println(up.String())
	fmt.Println(up.Password())
}
