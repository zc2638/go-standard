package main

import (
	"fmt"
	"log"
	"strconv"
)

// strconv包实现了基本数据类型和其字符串表示的相互转换
func main() {

	// bool操作
	exampleBool()
	// float操作
	exampleFloat()
	// int操作
	exampleInt()
	// uint操作
	exampleUint()
	// 转义操作
	exampleQuote()
	// 反转义操作
	exampleUnquote()
	// 数字解析错误操作
	exampleNumError()
}

func exampleBool() {

	// bool转string
	s := strconv.FormatBool(true)
	fmt.Printf("%T, %v\n", s, s)

	// string转bool
	bl, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T, %v\n", bl, bl)

	// 添加string类型的bool值
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))
}

func exampleFloat() {

	v := 3.1415926535

	// float32转string
	// fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）
	// prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f
	s32 := strconv.FormatFloat(float64(v), 'f', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)

	// float64转string
	s64 := strconv.FormatFloat(v, 'f', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)

	// string转float32
	if s, err := strconv.ParseFloat(s32, 32); err == nil {
		s1 := float32(s)
		fmt.Printf("%T, %v\n", s1, s1)
	}

	// string转float64
	if s, err := strconv.ParseFloat(s64, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	// 添加string类型的float32值
	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, v, 'f', -1, 32)
	fmt.Println(string(b32))

	// 添加string类型的float64值
	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, v, 'f', -1, 64)
	fmt.Println(string(b64))
}

func exampleInt() {

	v := 12

	// int转string
	s := strconv.Itoa(v)
	fmt.Printf("%T, %v\n", s, s)

	// string转int
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T, %v\n", i, i)

	// int64转十进制string
	// base 必须在2到36之间
	s = strconv.FormatInt(int64(v), 10)
	fmt.Printf("%T, %v\n", s, s)

	// 十进制string转int64
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T, %v\n", i64, i64)

	// 添加string类型的int十进制值
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	// 添加string类型的int十六进制值
	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
}

func exampleUint() {

	v := uint64(42)

	// uint64转十进制string
	s := strconv.FormatUint(v, 10)
	fmt.Printf("%T, %v\n", s, s)

	// 十进制string转uint64
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T, %v\n", u, u)

	// 添加string类型的uint十进制值
	b10 := []byte("uint (base 10):")
	b10 = strconv.AppendUint(b10, 42, 10)
	fmt.Println(string(b10))
}

func exampleQuote() {

	v := `"Fran & Freddie's Diner	☺"`
	vr := '☺'

	// 判断字符串 s 是否可以表示为一个单行的“反引号”字符串
	// 字符串中不能含有控制字符（除了 \t）和“反引号”字符，否则返回 false
	// 判断是否可以转义
	ok := strconv.CanBackquote(v)
	fmt.Println(ok)

	// 返回字符串s在go语法下的双引号字面值表示，控制字符、不可打印字符会进行转义。（如\t，\n，\xFF，\u0100）
	// 转义字符串
	s := strconv.Quote(v)
	fmt.Println(s)

	// 返回字符串s在go语法下的双引号字面值表示，控制字符和不可打印字符、非ASCII字符会进行转义
	// 转义字符串为ASCII
	sa := strconv.QuoteToASCII(v)
	fmt.Println(sa)

	// 判断 Unicode 字符 r 是否是一个可显示的字符
	// 可否显示并不是你想象的那样，比如空格可以显示，而\t则不能显示
	// 判断rune是否可以转义
	rok := strconv.IsPrint(vr)
	fmt.Println(rok)

	// 返回字符r在go语法下的单引号字面值表示，控制字符、不可打印字符会进行转义。（如\t，\n，\xFF，\u0100）
	// rune转string
	sr := strconv.QuoteRune(vr)
	fmt.Println(sr)

	// 返回字符串s在go语法下的双引号字面值表示，控制字符和不可打印字符、非ASCII字符会进行转义
	// rune转ASCII
	sra := strconv.QuoteRuneToASCII(vr)
	fmt.Println(sra)

	// 将字符串转义并添加
	b := strconv.AppendQuote([]byte("quote:"), `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))

	// 将字符串转义为ASCII并添加
	ba := strconv.AppendQuoteToASCII([]byte("quote (ascii):"), v)
	fmt.Println(string(ba))

	// 将rune转义并添加
	br := strconv.AppendQuoteRune([]byte("rune:"), vr)
	fmt.Println(string(br))

	// 将rune转义为ASCII并添加
	bra := strconv.AppendQuoteRuneToASCII([]byte("rune (ascii):"), vr)
	fmt.Println(string(bra))
}

func exampleUnquote() {

	// 与Quote相反
	// 将“带引号的字符串” s 转换为常规的字符串（不带引号和转义字符）
	// s 可以是“单引号”、“双引号”或“反引号”引起来的字符串（包括引号本身）
	// 如果 s 是单引号引起来的字符串，则返回该该字符串代表的字符
	s, err := strconv.Unquote("\"The string must be either double-quoted\"")
	fmt.Printf("%q, %v\n", s, err)

	// 将 转义后的字符串s 中的第一个字符“取消转义”并解码
	// 参数quote：字符串使用的“引号符”（用于对引号符“取消转义”）
	// 参数 quote 为“引号符”
	// value：解码后的字符；multibyte：value 是否为多字节字符；tail： 字符串s 除去 value 后的剩余部分；error：返回语法错误
	// 如果设置为单引号，则 s 中允许出现 \' 字符，不允许出现单独的 ' 字符
	// 如果设置为双引号，则 s 中允许出现 \" 字符，不允许出现单独的 " 字符
	// 如果设置为 0，则不允许出现 \' 或 \" 字符，可以出现单独的 ' 或 " 字符
	v, mb, t, err := strconv.UnquoteChar(`\"Fran & Freddie's Diner\"`, '"')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value:", string(v))
	fmt.Println("multibyte:", mb)
	fmt.Println("tail:", t)
}

func exampleNumError() {

	str := "Not a number"

	// 将一个不是数字的字符串解析为float类型
	if _, err := strconv.ParseFloat(str, 64); err != nil {

		// error 定义为 *strconv.NumError类型
		e := err.(*strconv.NumError)

		// 失败的函数（ParseBool、ParseInt、ParseUint、ParseFloat）
		fmt.Println("Func:", e.Func)

		// 输入的字符串
		fmt.Println("Num:", e.Num)

		// 失败的原因（ErrRange、ErrSyntax）
		fmt.Println("Err:", e.Err)

		fmt.Println(err)
	}
}
