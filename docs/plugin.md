# Go Plugin

## 一、概念

Golang是静态编译型语言，在编译时就将所有引用的包（库）全部加载打包到最终的可执行程序（或库文件）中，因此并不能在运行时动态加载其他共享库。Go
Plugin提供了这样一种方式，能够让你在运行时动态加载外部功能。

## 二、优势

- 可插拔：有了Plugin，我的程序可以根据需要随时替换其中某些部件而不用修改我的程序；
- 动态加载的需要：有些模块只有在运行时才能确定，需要动态加载外部的功能模块；
- 独立开发：Plugin 可以和主程序独立建设，主程序只需要制定好框架，实现默认（模版）功能。Plugin
  可根据用户需求随时自行扩展开发，运行时随意替换，提高了程序的可定制性；

## 三、方法

Golang 对 Plugin 的实现在标准库`plugin`中

```go
type Plugin struct{ ... }
func Open(path string) (*Plugin, error)
func (p *Plugin) Lookup(symName string) (Symbol, error)

type Symbol interface{}
```

### Plugin

Golang加载的插件结构，与之有关的两个方法：

- `Open`: 根据参数`path`提供的插件路径加载这个插件，并返回插件这个插件结构的指针`*Glugin`
- `Lookup`:`*Plugin`的惟一方法，通过名称`symName`在插件中寻找对应的变量或方法，以`Symbol`的形式返回

### Symbol

根据定义`type Symbol interface{}`，`Symbol`是`interface`的别名，也就是说，我们可以从插件里面拿到任何类型的可导出元素

## 三、构建插件

```go
package main

import (
	"fmt"
)

func Hello() {
	fmt.Println("Hello World From Plugin!")
}
```

执行构建命令，会生成对应的 `.so` 文件

```shell
go build --buildmode=plugin -o pluginhello.so pluginhello.go
```

## 四、使用插件

```go
package main

import (
	"log"
	"os"
	"plugin"
)

func main() {
	p, err := plugin.Open("./pluginhello.so")
	if err != nil {
		log.Fatalf("open plugin failed: %v", err)
	}
	s, err := p.Lookup("Hello")
	if err != nil {
		log.Fatalf("lookup Hello symbol failed: %v", err)
	}
	if hello, ok := s.(func()); ok {
		hello()
	}
}
```

## 五、原理解析

https://zhuanlan.zhihu.com/p/385530871

## 六、注意点

#### 1、插件 和 主程序 需使用相同的 `golang` 版本

由于插件提供的代码将与主代码在相同的进程空间中运行，因此编译的二进制文件应与主应用程序 100％ 兼容。

#### 2、插件 和 主程序 使用的依赖库版本必须一致

#### 3、不要使用 vendor 目录

#### 4、不支持静态编译

因为需要使用CGO

#### 5、方法或者变量名必须是导出类型，即首字母大写

#### 6、源码需要在main包中，否则无法编译

#### 7、插件对主程序的函数依赖问题

可使用定义接口的方式、将实现了接口的实例传入插件中使用

## 七、其它

RPC插件：[GitHub - hashicorp/go-plugin: Golang plugin system over RPC.](https://github.com/hashicorp/go-plugin)

Go原生插件使用问题解析： https://www.sofastack.tech/blog/go-native-plug-in-use-problem-full-analysis/
