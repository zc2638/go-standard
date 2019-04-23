package main

import (
	"fmt"
	"strings"
	"unicode"
)

// strings包实现了用于操作字符的简单函数
func main() {

	strDemo()

	// 字符串倒序
	s := reverse("Hello, World!")
	fmt.Println(s)
}

func strDemo() {

	var a, b, c string

	// 根据字符串创建reader
	strings.NewReader("Hello World!")

	// 使用提供的多组old、new字符串对创建并返回一个*strings.Replacer
	// 替换是依次进行的，匹配时不会重叠
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))

	// 比较a和b, 返回 0: a等于b, 1: a大于b, -1: a小于b
	strings.Compare(a, b)

	// 判断a与b是否相同(将unicode大写、小写、标题三种格式字符视为相同)
	strings.EqualFold(a, b)

	// 判断a是否以b开头，当b为空时true
	strings.HasPrefix(a, b)

	// 判断a是否以b结尾，当b为空时true
	strings.HasSuffix(a, b)

	// 如果a以b结尾，则返回a去掉b结尾部分的新string。如果不是，返回a
	strings.TrimSuffix(a, b)

	// 如果a以b开头，则返回a去掉b开头部分的新string。如果不是，返回a
	strings.TrimPrefix(a, b)

	// 去除开头结尾所有的 空格换行回车缩进
	strings.TrimSpace(a)

	// 去除开头结尾所有的 指定字符串中的任意字符
	strings.Trim(a, " ")

	// 按自定义方法 去除开头结尾所有指定内容
	strings.TrimFunc(a, unicode.IsLetter)

	// 去除开头所有的 指定字符串中的任意字符
	strings.TrimLeft(a, "0123456789")

	// 按自定义方法 去除开头所有指定内容
	strings.TrimLeftFunc(a, unicode.IsLetter)

	// 去除结尾所有的 指定字符串中的任意字符
	strings.TrimRight(a, "0123456789")

	// 按自定义方法 去除结尾所有指定内容
	strings.TrimRightFunc(a, unicode.IsLetter)

	// 以一个或者多个空格分割成切片
	strings.Fields("  foo bar  baz   ")

	// 根据指定方法分割成切片
	strings.FieldsFunc("  foo1;bar2,baz3...", func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) // 以 不是字符或者数字 进行分割
	})

	// 判断a是否包含b
	strings.Contains(a, b)

	// 判断是否包含字符串中任意字符，只要包含字符串中一个及以上字符返回true，否则false
	strings.ContainsAny("I like seafood.", "fÄo!")

	// 判断是否包含rune字符
	strings.ContainsRune("I like seafood.", 'f')

	// 统计a中包含所有b的个数，如果b为空则返回a的长度
	strings.Count(a, b)

	// 检索b在a中第一次出现的位置，未检索到返回-1
	strings.Index(a, b)

	// 检索字符c在s中第一次出现的位置，不存在则返回-1
	strings.IndexByte(a, byte('k'))

	// 自定义方法检索首个字符的位置，未检索到返回-1
	strings.IndexFunc("Hello, 世界", func(c rune) bool {
		return unicode.Is(unicode.Han, c) // 是否包含中文字符
	})

	// 检索a中首个 字符串中任意字符 的位置，未检索到返回-1
	strings.IndexAny(a, "abc")

	// 检索a中首个 rune类型字符 的位置，未检索到返回-1
	strings.IndexRune("chicken", 'k')

	// 检索a中最后个b的位置，未检索到返回-1
	strings.LastIndex(a, b)

	// 检索a中最后个 byte类型字符 的位置，未检索到返回-1
	strings.LastIndexByte(a, byte('k'))

	// 自定义方法检索最后个字符的位置，未检索到返回-1
	strings.LastIndexFunc(a, unicode.IsLetter)

	// 将string数组以指定字符连接成一个新的string
	s := []string{a, b}
	strings.Join(s, ",")

	// 返回count个s串联的字符串
	// 例如：a = "abc"，返回 "abcabc"
	strings.Repeat(a, 2)

	// 返回一个 将a中的b替换为c 的新string，n为替换个数，-1替换所有
	strings.Replace(a, b, c, -1)

	// 返回一个 将a中的b替换为c 的新string
	strings.ReplaceAll(a, b, c)

	// 将a以指定字符分割成string数组
	strings.Split(a, b)

	// 将a以指定字符分割成string数组, n为分割个数，-1分割所有
	strings.SplitN(a, b, 2)

	// 将a以指定字符分割成string数组，保留b。
	strings.SplitAfter(a, b)

	// 将a以指定字符分割成string数组，保留b。n为分割个数，-1分割所有
	strings.SplitAfterN(a, b, 2)

	// 返回一个 以空格为界限，所有首个字母大写 的标题格式
	strings.Title(a)

	// 返回一个 所有字母大写 的标题格式
	strings.ToTitle(a)

	// 使用指定的映射表将 a 中的所有字符修改为标题格式返回。
	strings.ToTitleSpecial(unicode.SpecialCase{}, a)

	// 所有字母大写
	strings.ToUpper(a)

	// 使用指定的映射表将 a 中的所有字符修改为大写格式返回。
	strings.ToUpperSpecial(unicode.SpecialCase{}, a)

	// 所有字母小写
	strings.ToLower(a)

	// 使用指定的映射表将 a 中的所有字符修改为大写格式返回。
	strings.ToLowerSpecial(unicode.SpecialCase{}, a)

	// 遍历a按指定的rune方法处理每个字符
	strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return r
		} else {
			return 'a'
		}
	}, a)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}