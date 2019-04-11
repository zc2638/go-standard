package main

import (
	"fmt"
	"path"
)

// path实现了对斜杠分隔的路径的实用操作函数
func main() {

	var pathStr = "/testdata/test/"

	// 返回路径的最后一个元素，在提取元素前会删掉末尾的斜杠
	// 如果路径是""，会返回"."；如果路径是只有一个斜杆构成，会返回"/"
	path.Base(pathStr) // test

	// 通过单纯的词法操作返回和path代表同一地址的最短路径
	// 它会不断的依次应用如下的规则，直到不能再进行任何处理：
	// 1. 将连续的多个斜杠替换为单个斜杠
	// 2. 剔除每一个.路径名元素（代表当前目录）
	// 3. 剔除每一个路径内的..路径名元素（代表父目录）和它前面的非..路径名元素
	// 4. 剔除开始一个根路径的..路径名元素，即将路径开始处的"/.."替换为"/"
	path.Clean("a//c") // a/c

	// 返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录
	// 在使用Split去掉最后一个元素后，会简化路径并去掉末尾的斜杠
	// 如果路径是空字符串，会返回"."；如果路径由1到多个斜杠后跟0到多个非斜杠字符组成，会返回"/"；其他任何情况下都不会返回以斜杠结尾的路径
	path.Dir(pathStr) // /testdata/test

	// 返回path文件扩展名
	path.Ext("/a/test.html") // html

	// 返回路径是否是一个绝对路径
	path.IsAbs("/dev/null")

	// 将路径从最后一个斜杠后面位置分隔为两个部分（dir和file）并返回
	// 如果路径中没有斜杠，函数返回值dir会设为空字符串，file会设为path
	// 两个返回值满足path == dir+file
	dir, file := path.Split(pathStr)
	fmt.Println(dir, file)

	// 可以将任意数量的路径元素放入一个单一路径里，会根据需要添加斜杠
	// 结果是经过简化的，所有的空字符串元素会被忽略
	path.Join("a", "b/c") // /a/b/c

	// 如果name匹配shell文件名模式匹配字符串，Match函数返回真
	// Match要求匹配整个name字符串，而不是它的一部分。只有pattern语法错误时，会返回ErrBadPattern
	/*
	匹配字符串语法：
		pattern:
			{ term }
		term:
			'*'                                  匹配0或多个非/的字符
			'?'                                  匹配1个非/的字符
			'[' [ '^' ] { character-range } ']'  字符组（必须非空）
			c                                    匹配字符c（c != '*', '?', '\\', '['）
			'\\' c                               匹配字符c
		character-range:
			c           匹配字符c（c != '\\', '-', ']'）
			'\\' c      匹配字符c
			lo '-' hi   匹配区间[lo, hi]内的字符
	 */
	path.Match("*", "abc")
}