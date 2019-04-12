package main

import (
	"os"
	"text/tabwriter"
)

// tabwriter包实现了写入过滤器（tabwriter.Writer），可以将输入的缩进修正为正确的对齐文本
func main() {

	// 创建一个过滤器
	w := new(tabwriter.Writer)

	// 初始化，第一个参数指定输出目标
	// minwidth 最小单元长度
	// tabwidth tab字符的宽度
	// padding  计算单元宽度时会额外加上它
	// padchar  用于填充的ASCII字符，如果是'\t'，则Writer会假设tabwidth作为输出中tab的宽度，且单元必然左对齐
	// flags    格式化控制
	// 等同于 tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	w.Write([]byte("a\tb\tc\td\t."))
	w.Write([]byte("123\t12345\t1234567\t123456789\t.\n"))

	// 在最后一次调用Write后，必须调用Flush方法以清空缓存，并将格式化对齐后的文本写入生成时提供的output中
	w.Flush()

	// 最小单元长度为5，宽度额外增加1，空格填充，右对齐
	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)
	w.Write([]byte("a\tb\tc\td\t."))
	w.Write([]byte("123\t12345\t1234567\t123456789\t.\n"))
	w.Flush()
}