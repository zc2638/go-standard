package main

import (
	"unicode/utf8"
)

// utf8包实现了对utf-8文本的常用函数和常数的支持，包括rune和utf-8编码byte序列之间互相翻译的函数
func main() {

	var b = []byte("Hello World!")

	// 判断切片p是否包含完整且合法的utf-8编码序列
	utf8.Valid(b)

	// 判断r是否可以编码为合法的utf-8序列
	utf8.ValidRune('H')

	// 判断s是否包含完整且合法的utf-8编码序列
	utf8.ValidString(string(b))

	// 将r的utf-8编码序列写入p（p必须有足够的长度），并返回写入的字节数
	utf8.EncodeRune(b, 'H')

	// 解码p开始位置的第一个utf-8编码的码值，返回该码值和编码的字节数
	// 如果编码不合法，会返回(RuneError, 1)。该返回值在正确的utf-8编码情况下是不可能返回的
	utf8.DecodeRune(b)

	// 类似DecodeRune但输入参数是字符串
	utf8.DecodeRuneInString(string(b))

	// 解码p中最后一个utf-8编码序列，返回该码值和编码序列的长度
	utf8.DecodeLastRune(b)

	// 类似DecodeLastRune但输入参数是字符串
	utf8.DecodeLastRuneInString(string(b))

	// 判断切片p是否以一个码值的完整utf-8编码开始
	// 不合法的编码因为会被转换为宽度1的错误码值而被视为完整的
	// 如中文字符占3位byte，一位byte判断为false，完整的3位为true
	utf8.FullRune(b)

	// 类似FullRune但输入参数是字符串
	utf8.FullRuneInString(string(b))

	// 返回p中的utf-8编码的码值的个数。错误或者不完整的编码会被视为宽度1字节的单个码值
	utf8.RuneCount(b)

	// 类似RuneCount但输入参数是一个字符串
	utf8.RuneCountInString(string(b))

	// 返回r编码后的字节数。如果r不是一个合法的可编码为utf-8序列的值，会返回-1
	utf8.RuneLen('世')

	// 判断字节b是否可以作为某个rune编码后的第一个字节。第二个即之后的字节总是将左端两个字位设为10
	utf8.RuneStart('世')
}