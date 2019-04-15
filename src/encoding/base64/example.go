package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

// 实现了base64编码
func main() {


	// 声明内容
	var origin = []byte("Hello World!")
	// 声明buffer
	var buf bytes.Buffer
	// 自定一个64字节的字符串
	var customEncode = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	// 使用给出的字符集生成一个*base64.Encoding，字符集必须是32字节的字符串
	e := base64.NewEncoding(customEncode)
	// 创建一个新的base64流编码器
	w := base64.NewEncoder(e, &buf)
	// 写入
	if _, err := w.Write(origin); err != nil {
		log.Fatal(err)
	}
	// 关闭
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("base64编码内容: ", string(buf.Bytes()))

	// 创建一个新的base64流解码器
	r := base64.NewDecoder(base64.StdEncoding, &buf)
	// 读取内容
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("base64解码内容: ", string(b))



	// 使用标准的base64编码字符集编码
	originEncode := base64.StdEncoding.EncodeToString(origin)
	fmt.Println("base64编码内容: ", originEncode)

	// 使用标准的base64编码字符集解码
	originBytes, err := base64.StdEncoding.DecodeString(originEncode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("base64解码内容: ", string(originBytes))



	// 获取数据进行base64编码后的最大长度
	var ne = base64.StdEncoding.EncodedLen(len(origin))
	// 声明[]byte
	var dst = make([]byte, ne)
	// 将src的数据编码后存入dst，最多写EncodedLen(len(src))字节数据到dst，并返回写入的字节数
	base64.StdEncoding.Encode(dst, origin)
	fmt.Println("base64编码内容: ", string(dst))

	// 获取base64编码的数据解码后的最大长度
	var nd = base64.StdEncoding.DecodedLen(len(dst))
	// 声明[]byte
	var originText = make([]byte, nd)
	if _, err := base64.StdEncoding.Decode(originText, dst); err != nil {
		log.Fatal(err)
	}
	fmt.Println("base64解码内容: ", string(originText))

	// 创建与enc相同的新编码，指定的填充字符除外，或者nopadding禁用填充。填充字符不能是'\r'或'\n'，不能包含在编码的字母表中，并且必须是等于或小于'\xff'的rune
	base64.StdEncoding.WithPadding(base64.StdPadding)
	// base64.StdEncoding 定义标准base64编码字符集
	// base64.URLEncoding 定义用于URL和文件名的，base64编码字符集
	// base64.RawStdEncoding 定义标准无填充字符的base64编码字符集
	// base64.RawURLEncoding 定义用于URL和文件名的，无填充字符的base64编码字符集
}