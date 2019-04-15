package main

import (
	"bytes"
	"encoding/base32"
	"fmt"
	"io/ioutil"
	"log"
)

// 实现了base32编码
func main() {

	// 声明内容
	var origin = []byte("Hello World!")
	// 声明buffer
	var buf bytes.Buffer
	// 自定一个32字节的字符串
	var customEncode = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	// 使用给出的字符集生成一个*base32.Encoding，字符集必须是32字节的字符串
	e := base32.NewEncoding(customEncode)
	// 创建一个新的base32流编码器
	w := base32.NewEncoder(e, &buf)
	// 写入
	if _, err := w.Write(origin); err != nil {
		log.Fatal(err)
	}
	// 关闭
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("base32编码内容: ", string(buf.Bytes()))

	// 创建一个新的base32流解码器
	r := base32.NewDecoder(base32.StdEncoding, &buf)
	// 读取内容
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("base32解码内容: ", string(b))



	// 使用标准的base32编码字符集编码
	originEncode := base32.StdEncoding.EncodeToString(origin)
	fmt.Println("base32编码内容: ", originEncode)

	// 使用标准的base32编码字符集解码
	originBytes, err := base32.StdEncoding.DecodeString(originEncode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("base32解码内容: ", string(originBytes))



	// 获取数据进行base32编码后的最大长度
	var ne = base32.StdEncoding.EncodedLen(len(origin))
	// 声明[]byte
	var dst = make([]byte, ne)
	// 将src的数据编码后存入dst，最多写EncodedLen(len(src))字节数据到dst，并返回写入的字节数
	base32.StdEncoding.Encode(dst, origin)
	fmt.Println("base32编码内容: ", string(dst))

	// 获取base32编码的数据解码后的最大长度
	var nd = base32.StdEncoding.DecodedLen(len(dst))
	// 声明[]byte
	var originText = make([]byte, nd)
	if _, err := base32.StdEncoding.Decode(originText, dst); err != nil {
		log.Fatal(err)
	}
	fmt.Println("base32解码内容: ", string(originText))

	// base32.HexEncoding 定义用于"扩展Hex字符集"，用于DNS
}
