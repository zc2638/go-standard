package main

import (
	"fmt"
	"unicode"
)

// unicode包提供数据和函数来测试Unicode代码点的一些属性
func main() {

	// 判断示例
	exampleIs()
	// 对应示例
	exampleSimpleFold()
	// 转换示例
	exampleTo()
}

func exampleIs() {

	// constant with mixed type runes
	const mixed = "\b5Ὂg̀9! ℃ᾭG"
	for _, c := range mixed {

		fmt.Printf("For %q:\n", c)

		// 判断一个字符是否是控制字符，主要是策略C的字符和一些其他的字符如代理字符
		if unicode.IsControl(c) {
			fmt.Println("\tis control rune")
		}

		// 判断一个r字符是否是十进制数字字符
		if unicode.IsDigit(c) {
			fmt.Println("\tis digit rune")
		}

		// 判断一个字符是否是unicode图形。包括字母、标记、数字、符号、标点、空白，参见L、M、N、P、S、Zs
		if unicode.IsGraphic(c) {
			fmt.Println("\tis graphic rune")
		}

		// 判断一个字符是否是字母
		if unicode.IsLetter(c) {
			fmt.Println("\tis letter rune")
		}

		// 判断字符是否是小写字母
		if unicode.IsLower(c) {
			fmt.Println("\tis lower case rune")
		}

		// 判断字符是否是大写字母
		if unicode.IsUpper(c) {
			fmt.Println("\tis upper case rune")
		}

		// 判断一个字符是否是标记字符
		if unicode.IsMark(c) {
			fmt.Println("\tis mark rune")
		}

		// 判断一个字符是否是数字字符
		if unicode.IsNumber(c) {
			fmt.Println("\tis number rune")
		}

		// 判断一个字符是否是go的可打印字符
		// 本函数基本和IsGraphic一致，只是ASCII空白字符U+0020会返回假
		if unicode.IsPrint(c) {
			fmt.Println("\tis printable rune")
		}

		// 判断一个字符是否是unicode标点字符
		if unicode.IsPunct(c) {
			fmt.Println("\tis punct rune")
		}

		// 判断一个字符是否是空白字符
		// 在Latin-1字符空间中，空白字符为：'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).其它的空白字符请参见策略Z和属性Pattern_White_Space
		if unicode.IsSpace(c) {
			fmt.Println("\tis space rune")
		}

		// 判断一个字符是否是unicode符号字符
		if unicode.IsSymbol(c) {
			fmt.Println("\tis symbol rune")
		}

		// 判断字符是否是标题字母
		if unicode.IsTitle(c) {
			fmt.Println("\tis title case rune")
		}
	}
}

func exampleSimpleFold() {

	// 迭代在unicode标准字符映射中互相对应的unicode码值
	// 在与r对应的码值中（包括r自身），会返回最小的那个大于r的字符（如果有）；否则返回映射中最小的字符
	fmt.Printf("%#U\n", unicode.SimpleFold('A'))      // 'a'
	fmt.Printf("%#U\n", unicode.SimpleFold('a'))      // 'A'
	fmt.Printf("%#U\n", unicode.SimpleFold('K'))      // 'k'
	fmt.Printf("%#U\n", unicode.SimpleFold('k'))      // '\u212A' (Kelvin symbol, K)
	fmt.Printf("%#U\n", unicode.SimpleFold('\u212A')) // 'K'
	fmt.Printf("%#U\n", unicode.SimpleFold('1'))      // '1'
}

func exampleTo() {

	const lcG = 'g'

	// 转大写
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, lcG))
	fmt.Println(unicode.ToUpper(lcG))
	// 转小写
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, lcG))
	fmt.Println(unicode.ToLower(lcG))
	// 转标题
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, lcG))
	fmt.Println(unicode.ToTitle(lcG))
}