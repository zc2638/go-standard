package main

import (
	"fmt"
	"html"
)

// html包提供了用于转义和解转义HTML文本的函数
func main() {

	var data = "<div>Hello World!</div>"

	// 将特定的一些字符转为逸码后的字符实体，如"<"变成"&lt;"
	// 它只会修改五个字符：<、>、&、'、"
	// UnescapeString(EscapeString(s)) == s总是成立，但是两个函数顺序反过来则不一定成立
	str := html.EscapeString(data)
	fmt.Println(str)

	// 将逸码的字符实体如"&lt;"修改为原字符"<"
	// 它会解码一个很大范围内的字符实体，远比函数EscapeString转码范围大得多。例如"&aacute;"解码为"á"，"&#225;"和"&xE1;"也会解码为该字符
	origin := html.UnescapeString(str)
	fmt.Println(origin)
}