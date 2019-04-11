package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/textproto"
)

// textproto实现了对基于文本的请求/回复协议的一般性支持，包括HTTP、NNTP和SMTP
func main() {

	// 基础读写
	example()
}

func example() {

	// 声明buffer
	var buf bytes.Buffer

	// 根据bufio.Writer创建writer
	w := textproto.NewWriter(bufio.NewWriter(&buf))

	// 写入格式化的输出，结尾自动添加\r\n
	if err := w.PrintfLine("Hello %s!", "World"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	// 返回一个写入器，可以用来为 w 写一个点编码
	// 它需要在需要时插入前导点，将行结束符\n 转换为\r\n，并在 DotWriter 关闭时添加最后的\r\n 行
	// 调用者应在下一次调用 w 上的方法之前关闭 DotWriter
	d := w.DotWriter()

	// 写入
	if _, err := d.Write([]byte("abc\n.def\n..ghi\n.jkl\n.")); err != nil {
		log.Fatal(err)
	}
	// 关闭dotWriter，应在下一次调用 w 上的方法之前关闭
	if err := d.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	// 根据bufio.Reader创建一个reader
	r := textproto.NewReader(bufio.NewReader(&buf))


	// 从 r 读取一行，从返回的字符串中删除最后的\n或\r\n
	fmt.Println(r.ReadLine())

	// 返回一个新的 Reader，它使用从 r 读取的点编码块的解码文本满足 Reads
	// 返回的 Reader 只有在下一次调用 r 时才有效
	// 点编码是用于文本协议（如 SMTP）中的数据块的常用成帧
	// 数据由一系列行组成，每行以“\ r \ n”结尾
	// 序列本身结束于只包含一个点的一行：“。\ r \ n”。以点开头的行会用另外一个点进行转义，以避免看起来像序列的结尾
	dr := r.DotReader()

	// 读取内容
	b ,err := ioutil.ReadAll(dr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
