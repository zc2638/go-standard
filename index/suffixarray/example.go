package main

import (
	"bytes"
	"index/suffixarray"
	"regexp"
)

// suffixarray包通过使用内存中的后缀树实现了对数级时间消耗的子字符串搜索。
func main() {

	// 声明buffer
	var buf bytes.Buffer

	// 声明内容
	var data = []byte("Hello Gopher!")
	var src = []byte("Gopher")

	// 生成一个*Index，时间复杂度O(N*log(N))。
	index := suffixarray.New(data)

	// 返回创建x时提供的[]byte数据，注意不能修改返回值
	index.Bytes()

	// 从reader中读取一个index写入x，x不能为nil
	index.Read(&buf)

	// 将x中的index写入writer中，x不能为nil
	index.Write(&buf)

	// 返回一个未排序的列表，内为s在被索引为index的切片数据中出现的位置
	// 如果n<0，返回全部匹配；如果n==0或s为空，返回nil；否则n为result的最大长度
	// 时间复杂度O(log(N)*len(s) + len(result))，其中N是被索引的数据的大小
	index.Lookup(src, -1)

	// 返回一个正则表达式r的不重叠的匹配的经过排序的列表，一个匹配表示为一对指定了匹配结果的切片的索引（相对于x.Bytes())
	// 如果n<0，返回全部匹配；如果n==0或匹配失败，返回nil；否则n为result的最大长度
	index.FindAllIndex(regexp.MustCompile(`\d*`), -1)
}