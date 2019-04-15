package main

import (
	"bytes"
	"fmt"
	"unicode"
)

// bytes包实现了操作[]byte的常用函数。本包的函数和strings包的函数相当类似
func main() {

	// buffer写入
	write()
	// buffer直接创建
	bufferNew()
	// byte操作demo
	byteDemo()
}

func write() {

	// 声明buffer
	var b bytes.Buffer

	// 增加buffer容量
	b.Grow(64)

	// 写入字符串
	b.WriteString("Hello ")
	// 写入单个byte
	b.WriteByte('W')
	// 写入单个rune
	b.WriteRune('o')
	// 写入byte
	b.Write([]byte("rld!"))

	fmt.Printf("%s\n", b.Bytes())
}

func bufferNew() {

	str := "Hello World!"
	// 根据字符串创建buffer
	buf := bytes.NewBufferString(str)
	// 根据byte创建buffer
	buf = bytes.NewBuffer([]byte(str))

	fmt.Printf("%s\n", buf.Bytes())
}

func byteDemo() {

	var a, b, c []byte

	// 根据[]byte创建reader
	bytes.NewReader([]byte("Hello World!"))

	// 比较a和b, 返回 0: a等于b, 1: a包含b, -1: a不包含b
	bytes.Compare(a, b)

	// 判断a与b是否相同
	bytes.Equal(a, b)

	// 判断a与b是否相同，忽略大小写
	bytes.EqualFold(a, b)

	// 判断a是否以b开头，当b为空时true
	bytes.HasPrefix(a, b)

	// 判断a是否以b结尾，当b为空时true
	bytes.HasSuffix(a, b)

	// 如果a以b结尾，则返回a去掉b结尾部分的新byte。如果不是，返回a
	bytes.TrimSuffix(a, b)

	// 如果a以b开头，则返回a去掉b开头部分的新byte。如果不是，返回a
	bytes.TrimPrefix(a, b)

	// 去除开头结尾所有的 空格换行回车缩进
	bytes.TrimSpace(a)

	// 去除开头结尾所有的 指定字符串中的任意字符
	bytes.Trim(a, " ")

	// 按自定义方法 去除开头结尾所有指定内容
	bytes.TrimFunc(a, unicode.IsLetter)

	// 去除开头所有的 指定字符串中的任意字符
	bytes.TrimLeft(a, "0123456789")

	// 按自定义方法 去除开头所有指定内容
	bytes.TrimLeftFunc(a, unicode.IsLetter)

	// 去除结尾所有的 指定字符串中的任意字符
	bytes.TrimRight(a, "0123456789")

	// 按自定义方法 去除结尾所有指定内容
	bytes.TrimRightFunc(a, unicode.IsLetter)

	// 以一个或者多个空格分割成切片
	bytes.Fields([]byte("  foo bar  baz   "))

	// 根据指定方法分割成切片
	bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) // 以 不是字符或者数字 进行分割
	})

	// 判断a是否包含b
	bytes.Contains(a, b)

	// 判断byte是否包含字符串中任意字符，只要包含字符串中一个及以上字符返回true，否则false
	bytes.ContainsAny([]byte("I like seafood."), "fÄo!")

	// 判断byte是否包含rune字符
	bytes.ContainsRune([]byte("I like seafood."), 'f')

	// 统计a中包含所有b的个数，如果b为空则返回a的长度
	bytes.Count(a, b)

	// 检索a中首个b的位置，未检索到返回-1
	bytes.Index(a, b)

	// 检索a中首个 byte类型字符 的位置，未检索到返回-1
	bytes.IndexByte(a, byte('k'))

	// 自定义方法检索首个字符的位置，未检索到返回-1
	bytes.IndexFunc([]byte("Hello, 世界"), func(c rune) bool {
		return unicode.Is(unicode.Han, c) // 是否包含中文字符
	})

	// 检索a中首个 字符串中任意字符 的位置，未检索到返回-1
	bytes.IndexAny(a, "abc")

	// 检索a中首个 rune类型字符 的位置，未检索到返回-1
	bytes.IndexRune([]byte("chicken"), 'k')

	// 检索a中最后个b的位置，未检索到返回-1
	bytes.LastIndex(a, b)

	// 检索a中最后个 byte类型字符 的位置，未检索到返回-1
	bytes.LastIndexByte(a, byte('k'))

	// 自定义方法检索最后个字符的位置，未检索到返回-1
	bytes.LastIndexFunc(a, unicode.IsLetter)

	// 将byte数组以指定 byte字符 连接成一个新的byte
	s := [][]byte{a, b}
	bytes.Join(s, []byte(","))

	// 返回一个重复n次a的新byte
	// 例如：a = []byte("abc")，返回 []byte("abcabc")
	bytes.Repeat(a, 2)

	// 返回一个 将a中的b替换为c 的新byte，n为替换个数，-1替换所有
	bytes.Replace(a, b, c, -1)

	// 返回一个 将a中的b替换为c 的新byte
	bytes.ReplaceAll(a, b, c)

	// byte类型转rune类型
	bytes.Runes(a)

	// 将a以指定字符byte分割成byte数组
	bytes.Split(a, []byte(","))

	// 将a以指定字符byte分割成byte数组, n为分割个数，-1分割所有
	bytes.SplitN(a, []byte(","), 2)

	// 将a以指定字符byte分割成byte数组，保留b。
	bytes.SplitAfter(a, []byte(","))

	// 将a以指定字符byte分割成byte数组，保留b。n为分割个数，-1分割所有
	bytes.SplitAfterN(a, []byte(","), 2)

	// 返回一个 以空格为界限，所有首个字母大写 的标题格式
	bytes.Title(a)

	// 返回一个 所有字母大写 的标题格式
	bytes.ToTitle(a)

	// 使用指定的映射表将 a 中的所有字符修改为标题格式返回。
	bytes.ToTitleSpecial(unicode.SpecialCase{}, a)

	// 所有字母大写
	bytes.ToUpper(a)

	// 使用指定的映射表将 a 中的所有字符修改为大写格式返回。
	bytes.ToUpperSpecial(unicode.SpecialCase{}, a)

	// 所有字母小写
	bytes.ToLower(a)

	// 使用指定的映射表将 a 中的所有字符修改为大写格式返回。
	bytes.ToLowerSpecial(unicode.SpecialCase{}, a)

	// 遍历a按指定的rune方法处理每个字符
	bytes.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return r
		} else {
			return 'a'
		}
	}, a)
}
