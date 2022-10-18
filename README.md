# go-standard

![Go](https://raw.githubusercontent.com/zc2638/material/master/go/go.png)

Go常用规范定义，标准库方法使用示例，请注意这不是Go的中文版标准库

## [Golang 泛型](https://github.com/akutz/go-generics-the-hard-way)
## [Golang 常见指令](./docs/instruction.md)
## [Uber编码规范](./docs/style.md)

## Go Module代理
设置环境变量
```
GOPROXY=https://proxy.golang.org           # 官方推荐，国内无法正常使用
GOPROXY=https://proxy.golang.com.cn        # 全球通用
GOPROXY=https://goproxy.cn                 # 国内友好
```

## TODO
- 改写为中文版pkg，支持常用第三方包
- 更新到 `1.18+` 版本

## 简介

- [**archive**](https://github.com/zc2638/go-standard/tree/master/src/archive) &emsp;&emsp;&emsp;&emsp; tar/zip压缩操作
- [**bufio**](https://github.com/zc2638/go-standard/tree/master/src/bufio/example.go) &emsp;&emsp;&emsp;&emsp;&emsp; 有缓冲的I/O
- [**bytes**](https://github.com/zc2638/go-standard/tree/master/src/bytes/example.go) &emsp;&emsp;&emsp;&emsp;&emsp; 操作[]byte字节片
- [**compress**](https://github.com/zc2638/go-standard/tree/master/src/compress) &emsp;&emsp;&emsp; bzip2/flate/gzip/lzw/zlib压缩操作
- [**container**](https://github.com/zc2638/go-standard/tree/master/src/container) &emsp;&emsp;&emsp;&ensp;堆操作/双向链表/环形链表
- [**context**](https://github.com/zc2638/go-standard/tree/master/src/context/example.go) &emsp;&emsp;&emsp;&emsp;&nbsp;上下文类型
- [**crypto**](https://github.com/zc2638/go-standard/tree/master/src/crypto) &emsp;&emsp;&emsp;&emsp;&emsp;常用的密码（算法）
- [**database/sql**](https://github.com/zc2638/go-standard/tree/master/src/database/sql/example.go) &emsp;&emsp;数据库接口
- [**encoding**](https://github.com/zc2638/go-standard/tree/master/src/encoding) &emsp;&emsp;&emsp;&emsp;数据编码
- [**errors**](https://github.com/zc2638/go-standard/tree/master/src/errors/example.go) &emsp;&emsp;&emsp;&emsp;&emsp; 创建错误函数
- [**expvar**](https://github.com/zc2638/go-standard/tree/master/src/expvar/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;公共变量
- [**flag**](https://github.com/zc2638/go-standard/tree/master/src/flag/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&ensp;命令行参数解析
- [**fmt**](https://github.com/zc2638/go-standard/tree/master/src/fmt/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&ensp; 格式化I/O
- [**go**](https://github.com/zc2638/go-standard/tree/master/src/go/example.md) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;GO常用命令
- [**hash**](https://github.com/zc2638/go-standard/tree/master/src/hash) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;提供hash函数的接口
- [**html**](https://github.com/zc2638/go-standard/tree/master/src/html/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp; 转义和解转义HTML文本
- [**image**](https://github.com/zc2638/go-standard/tree/master/src/image/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&ensp;实现了基本的2D图片库
- [**index/suffixarray**](https://github.com/zc2638/go-standard/tree/master/src/index/suffixarray/example.go) &ensp;使用内存后缀数组以对数时间实现子字符串搜索
- [**io**](https://github.com/zc2638/go-standard/tree/master/src/io/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&ensp; I/O操作
- [**log**](https://github.com/zc2638/go-standard/tree/master/src/log) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;简单的日志服务
- [**math**](https://github.com/zc2638/go-standard/tree/master/src/math) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;基本的数学常数和数学函数
- [**mime**](https://github.com/zc2638/go-standard/tree/master/src/mime/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;实现了MIME的部分规定
- [**net**](https://github.com/zc2638/go-standard/tree/master/src/net) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket
- [**os**](https://github.com/zc2638/go-standard/tree/master/src/os) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp; 操作系统函数
- [**path**](https://github.com/zc2638/go-standard/tree/master/src/path/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp; 对斜杠分隔的路径的实用操作
- [**plugin**](https://github.com/zc2638/go-standard/tree/master/src/plugin/example.go) &emsp;&emsp;&emsp;&emsp;&emsp; 插件生成和加载
- [**reflect**](https://github.com/zc2638/go-standard/tree/master/src/reflect/example.go) &emsp;&emsp;&emsp;&emsp;&emsp; 反射操作任意类型对象
- [**regexp**](https://github.com/zc2638/go-standard/tree/master/src/regexp/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;正则表达式
- [**rutime**](https://github.com/zc2638/go-standard/tree/master/src/runtime) &emsp;&emsp;&emsp;&emsp;&emsp;&nbsp;提供和go运行时环境的互操作，如控制go程的函数
- [**sort**](https://github.com/zc2638/go-standard/tree/master/src/sort/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&ensp;排序切片和用户自定义数据集
- [**strconv**](https://github.com/zc2638/go-standard/tree/master/src/strconv/example.go) &emsp;&emsp;&emsp;&emsp;&ensp; 基本数据类型和其字符串类型的相互转换
- [**strings**](https://github.com/zc2638/go-standard/tree/master/src/strings/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;操作字符串
- [**sync**](https://github.com/zc2638/go-standard/tree/master/src/sync) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;提供了基本的同步基元，如互斥锁
- [**text**](https://github.com/zc2638/go-standard/tree/master/src/text) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&ensp;文本操作
- [**time**](https://github.com/zc2638/go-standard/tree/master/src/time/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp; 时间操作
- [**unicode**](https://github.com/zc2638/go-standard/tree/master/src/unicode) &emsp;&emsp;&emsp;&emsp;&ensp;unicode操作
- [**unsafe**](https://github.com/zc2638/go-standard/tree/master/src/unsafe/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;提供了一些跳过go语言类型安全限制的操作

### Tip
channel作为参数
```go
chan int   // 可读可写
<-chan int // 只读
chan<- int // 只写
```

### JetBrains 开源证书支持

`go-standard` 项目一直以来都是在 JetBrains 公司旗下的 GoLand 集成开发环境中进行开发，基于 **free JetBrains Open Source license(s)** 正版免费授权，在此表达谢意。

<a href="https://www.jetbrains.com/?from=go-standard" target="_blank"><img src="https://raw.githubusercontent.com/zc2638/material/master/jetbrains/jetbrains.png" width="250" align="middle"/></a>
