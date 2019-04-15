package main

import "unicode/utf16"

// utf16包实现了UTF-16序列的编解码
func main() {

	var b = []rune("Hello World!")

	// 判断r是否可以编码为一个utf16的代理对
	utf16.IsSurrogate('H')

	// 将unicode码值序列编码为utf-16序列
	e := utf16.Encode(b)

	// 将utf-16序列解码为unicode码值序列
	utf16.Decode(e)

	// 将unicode码值r编码为一个utf-16的代理对。如果不能编码，会返回(U+FFFD, U+FFFD)
	r1, r2 := utf16.EncodeRune('H')

	// 将utf-16代理对(r1, r2)解码为unicode码值。如果代理对不合法，会返回U+FFFD
	utf16.DecodeRune(r1, r2)
}