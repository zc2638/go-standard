package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

// scanner包提供对utf-8文本的token扫描服务。它会从一个io.Reader获取utf-8文本，通过对Scan方法的重复调用获取一个个token
// 为了兼容已有的工具，NUL字符不被接受。如果第一个字符是表示utf-8编码格式的BOM标记，会自动忽略该标记
// 一般Scanner会跳过空白和Go注释，并会识别所有go语言规格的字面量。它可以定制为只识别这些字面量的一个子集，也可以识别不同的空白字符
func main() {

	const src = `
// This is scanned code.
if a > 10 {
	someParsable = text
}`

	// 创建一个扫描服务
	s := new(scanner.Scanner)

	// 初始化
	s.Init(strings.NewReader(src))

	// 设置扫描的文件名
	s.Filename = "example"

	// 从资源循环token或者unicode字符并返回它
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Println(s.Position, s.TokenText())
	}
}