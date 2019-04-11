package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// regexp 包实现正则表达式搜索
// 接受的正则表达式的语法与 Perl，Python 和其他语言使用的通用语法相同。更准确地说，它是  RE2 接受的语法，并在https://golang.org/s/re2syntax中进行了描述，\C除外
// 有关语法的概述，请运行 `go doc regexp/syntax`
// 所有字符都是 UTF-8 编码的代码点
func main() {

	// 简单匹配示例
	exampleMatch()

	// 转义特殊字符
	exampleQuoteMeta()

	// 匹配示例
	exampleCompile()

	// 最长匹配示例
	exampleCompilePOSTIX()

	exampleSubMatch()

	// 匹配追加
	exampleExpand()
}

func exampleMatch() {

	data := "Hello World![test]"
	pattern := `\d*`

	// 匹配检查文本正则表达式是否与字节片匹配。更复杂的查询需要使用 Compile 和完整的 Regexp 接口
	ok, err := regexp.Match(pattern, []byte(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)

	// 检查文本正则表达式是否与 RuneReader 读取的文本匹配。更复杂的查询需要使用 Compile 和完整的 Regexp 接口
	ok, err = regexp.MatchReader(pattern, strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)

	// 检查文本正则表达式是否匹配字符串。更复杂的查询需要使用 Compile 和完整的 Regexp 接口
	ok, err = regexp.MatchString(pattern, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)
}

func exampleQuoteMeta() {

	// 返回一个字符串，它引用参数文本中的所有正则表达式元字符； 返回的字符串是一个匹配文本文本的正则表达式。例如，QuoteMeta（[foo]）返回\[foo\]
	// 特殊字符有：\.+*?()|[]{}^$
	// 这些字符用于实现正则语法，所以当作普通字符使用时需要转换
	str := regexp.QuoteMeta(`Escaping symbols like: .+*?()|[]{}^$`)
	fmt.Println(str)
}

func exampleCompile() {

	data := "Hello foo!"

	// 编译解析一个正则表达式，并且如果成功返回一个可以用来匹配文本的 Regexp 对象
	r, err := regexp.Compile(`foo.?`)
	if err != nil {
		log.Fatal(err)
	}

	// 等同于Compile。但如果表达式不能被解析就会panic
	r = regexp.MustCompile(`foo(.?)`)

	// 匹配字符串
	fmt.Println(r.MatchString(data))

	// 匹配字节切片
	fmt.Println(r.Match([]byte(data)))

	// 匹配文本
	fmt.Println(r.MatchReader(strings.NewReader(data)))

	// 返回正则表达式
	fmt.Println("正则表达式：", r.String())

	// 查找字节切片，返回第一个匹配的内容
	fmt.Printf("%s\n", r.Find([]byte(data)))

	// 查找字节切片，返回第一个匹配的位置
	// [起始位置, 结束位置]
	fmt.Println(r.FindIndex([]byte(data)))

	// 查找字符串，返回第一个匹配的内容
	fmt.Println(r.FindString(data))

	// 查找字符串，返回第一个匹配的位置
	// [起始位置, 结束位置]
	fmt.Println(r.FindStringIndex(data))

	// 查找字节切片，返回所有匹配的内容
	// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
	fmt.Printf("%q\n", r.FindAll([]byte(data), -1))

	// 查找字节切片，返回所有匹配的位置
	// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
	// [[起始位置, 结束位置], [起始位置, 结束位置]...]
	fmt.Println(r.FindAllIndex([]byte(data), -1))

	// 查找字符串，返回所有匹配的内容
	// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
	fmt.Println(r.FindAllString(data, -1))

	// 查找字符串，返回所有匹配的位置
	// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
	// [[起始位置, 结束位置], [起始位置, 结束位置]...]
	fmt.Println(r.FindAllStringIndex(data, -1))

	// 查找文本，返回第一个匹配的位置
	// [起始位置, 结束位置]
	fmt.Println(r.FindReaderIndex(strings.NewReader(data)))


	// 查找字符串，以匹配项为分割符 分割成多个子串
	// 最多分割出 n 个子串，第 n 个子串不再进行分割
	// 如果 n < 0，则分割所有子串
	// 返回分割后的子串列表
	fmt.Println(r.Split(data, -1))

	// 返回所有匹配项都共同拥有的前缀（去除可变元素）
	// prefix：共同拥有的前缀
	// complete：如果 prefix 就是正则表达式本身，则返回 true，否则返回 false
	fmt.Println(r.LiteralPrefix())

	// 在 src 中搜索匹配项，并替换为 repl 指定的内容
	// 全部替换，并返回替换后的结果
	fmt.Printf("%s\n", r.ReplaceAll([]byte(data), []byte("World!")))
	fmt.Printf("%s\n", r.ReplaceAllString(data, "World!"))

	// 在 src 中搜索匹配项，并替换为 repl 指定的内容
	// 如果 repl 中有“分组引用符”（$1、$name），则将“分组引用符”当普通字符处理
	// 全部替换，并返回替换后的结果
	fmt.Printf("%s\n", r.ReplaceAllLiteral([]byte(data), []byte("$World")))
	fmt.Printf("%s\n", r.ReplaceAllLiteralString(data, "$World"))

	// 在 src 中搜索匹配项，然后将匹配的内容经过 repl方法 处理后，替换 src 中的匹配项
	// 如果 repl 的返回值中有“分组引用符”（$1、$name），则将“分组引用符”当普通字符处理
	// 全部替换，并返回替换后的结果
	fmt.Printf("%s\n", r.ReplaceAllFunc([]byte(data), func(b []byte) []byte { return []byte("World!") }))
	fmt.Printf("%s\n", r.ReplaceAllStringFunc(data, func(s string) string { return "World!" }))

	// 切换到贪婪模式
	// 正则标记“非贪婪模式”(?U)，如`(?U)H[\w\s]+o`
	r.Longest()
}

func exampleCompilePOSTIX() {

	data := "Hello World!"

	// CompilePOSIX 的作用和 Compile 一样
	// 不同的是，CompilePOSIX 使用 POSIX 语法，
	// 同时，它采用最左最长方式搜索，
	// 而 Compile 采用最左最短方式搜索
	// POSIX 语法不支持 Perl 的语法格式：\d、\D、\s、\S、\w、\W
	r, err := regexp.CompilePOSIX(`foo.?`)
	if err != nil {
		log.Fatal(err)
	}

	// 等同于CompilePOSIX。但如果表达式不能被解析就会panic
	r = regexp.MustCompilePOSIX(`foo.?`)

	fmt.Println(r.MatchString(data))

	// 其它用法参见Compile
}

func exampleSubMatch() {

	data := "seafood fool"

	r := regexp.MustCompile(`foo(.?)`)

	// 统计正则表达式中的分组个数（不包括“非捕获的分组”）
	fmt.Println(r.NumSubexp())

	// 返回分组名称列表，未命名的分组返回空字符串
	// 返回值[0] 为整个正则表达式的名称
	// 返回值[1] 是分组 1 的名称
	// 返回值[2] 是分组 2 的名称
	fmt.Println(r.SubexpNames())

	// 查找字节切片，返回第一个匹配的内容
	// 同时返回子表达式匹配的内容
	// [[完整匹配项], [子匹配项], [子匹配项], ...]
	fmt.Printf("%s\n", r.FindSubmatch([]byte(data)))

	// 查找字节切片，返回第一个匹配的位置
	// 同时返回子表达式匹配的位置
	// [完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...]
	fmt.Println(r.FindSubmatchIndex([]byte(data)))

	// 查找字符串，返回第一个匹配的内容
	// 同时返回子表达式匹配的内容
	// [[完整匹配项], [子匹配项], [子匹配项], ...]
	fmt.Println(r.FindStringSubmatch(data))

	// 查找字符串，返回第一个匹配的位置
	// 同时返回子表达式匹配的位置
	// [完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...]
	fmt.Println(r.FindStringSubmatchIndex(data))


	// 查找字节切片，返回所有匹配的内容
	// 同时返回子表达式匹配的内容
	// [ [[完整匹配项], [子匹配项], [子匹配项], ...], [[完整匹配项], [子匹配项], [子匹配项], ...] ]
	fmt.Printf("%q\n", r.FindAllSubmatch([]byte(data), -1))

	// 查找字节切片，返回所有匹配的位置
	// 同时返回子表达式匹配的位置
	// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
	// [ [完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...], [完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...]... ]
	fmt.Println(r.FindAllSubmatchIndex([]byte(data), -1))

	// 查找字符串，返回所有匹配的内容
	// 同时返回子表达式匹配的内容
	// [ [[完整匹配项], [子匹配项], [子匹配项], ...], [[完整匹配项], [子匹配项], [子匹配项], ...] ]
	fmt.Println(r.FindAllStringSubmatch(data, -1))

	// 查找字节切片，返回所有匹配的位置
	// 同时返回子表达式匹配的位置
	// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
	// [ [完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...], [完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...]... ]
	fmt.Println(r.FindAllStringSubmatchIndex(data, -1))


	// 查找文本，返回第一个匹配的位置
	// 同时返回子表达式匹配的位置
	// [完整项起始, 完整项结束, 子项起始, 子项结束, 子项起始, 子项结束, ...]
	fmt.Println(r.FindReaderSubmatchIndex(strings.NewReader(data)))
}

// 追加
func exampleExpand() {

	content := `
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`

	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	template := "$key=$value\n"

	var result []byte
	var resultStr []byte

	for _, s := range pattern.FindAllSubmatchIndex([]byte(content), -1) {

		// 将 template 的内容经过处理后，追加到 dst 的尾部
		// template 中要有 $1、$2、${name1}、${name2} 这样的“分组引用符”
		// match 是由 FindSubmatchIndex 方法返回的结果，里面存放了各个分组的位置信息
		// 如果 template 中有“分组引用符”，则以 match 为标准，
		// 在 src 中取出相应的子串，替换掉 template 中的 $1、$2 等引用符号。
		result = pattern.Expand(result, []byte(template), []byte(content), s)
	}
	fmt.Println(string(result))

	for _, s := range pattern.FindAllStringSubmatchIndex(content, -1) {

		// 同Expand类似
		resultStr = pattern.ExpandString(resultStr, template, content, s)
	}
	fmt.Println(string(resultStr))
}
