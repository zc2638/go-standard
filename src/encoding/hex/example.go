package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

// hex包实现了16进制字符表示的编解码
func main() {

	// 编码/解码
	example()
	example2()

	// 编码/解码(string)
	exampleString()

	// hex dump编码
	exampleDump()
}

func example() {

	// 声明内容
	var src = []byte("Hello Gopher!")

	// 声明一个 明文数据编码后的编码数据的长度 的[]byte
	dst := make([]byte, hex.EncodedLen(len(src)))

	// 将src的数据解码为EncodedLen(len(src))字节，返回实际写入dst的字节数
	hex.Encode(dst, src)
	fmt.Printf("%s\n", dst)

	// 声明一个 编码数据解码后的明文数据的长度 的[]byte
	origin := make([]byte, hex.DecodedLen(len(dst)))
	// 将src解码为DecodedLen(len(src))字节，返回实际写入dst的字节数；如遇到非法字符，返回描述错误的error
	_, err := hex.Decode(origin, dst)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", origin)
}

func example2() {

	// 声明内容
	var src = []byte("Hello Gopher!")
	// 声明buffer
	var buf bytes.Buffer

	// 创建一个用于编码的writer
	w := hex.NewEncoder(&buf)
	if _, err := w.Write(src); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf.Bytes()))

	// 创建一个用于解码的reader
	r := hex.NewDecoder(&buf)
	// 读取内容
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func exampleString() {

	// 声明内容
	var src = []byte("Hello Gopher!")

	// 将数据src编码为字符串
	str := hex.EncodeToString(src)
	fmt.Println(str)

	// 将十六进制字符串解码为原数据
	origin, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(origin))
}

func exampleDump() {

	// 声明内容
	content := []byte("Go is an open source programming language.")

	// 返回回给定数据的hex dump格式的字符串，这个字符串与控制台下`hexdump -C`对该数据的输出是一致的
	str := hex.Dump(content)
	fmt.Printf("%s\n", str)



	// 声明buffer
	var buf bytes.Buffer
	// 返回一个io.WriteCloser接口，将写入的数据的hex dump格式写入w，具体格式为'hexdump -C'
	dumper := hex.Dumper(&buf)
	// 关闭
	defer dumper.Close()
	// 写入
	if _, err := dumper.Write(content); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf.Bytes())
}