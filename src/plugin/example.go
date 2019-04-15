package main

import (
	"log"
	"plugin"
)

const PluginFile = "testdata/testPlugin.so"

// plugin包实现 Go 插件的加载和符号解析
// 注意plugin的源码需要在main包中，否则无法编译
func main() {

	// 编译go plugin只需要在go build的时候带上--buildmode=plugin即可。命令如下：
	// go build --buildmode=plugin -o pluginHello.so pluginHello.go

	// 打开一个 Go 插件。如果路径已被打开，则返回现有的*plugin.Plugin。由多个 goroutines 并行使用是安全的
	p ,err := plugin.Open(PluginFile)
	if err != nil {
		log.Fatal(err)
	}

	// 查找在插件 p 中搜索名为 symName 的符号
	// 符号是任何导出的变量或函数，如果找不到该符号，它会报告错误。由多个 goroutines 并行使用是安全的
	s, err := p.Lookup("test")
	if err != nil {
		log.Fatal(err)
	}
	if test, ok := s.(func()); ok {
		test()
	}
}