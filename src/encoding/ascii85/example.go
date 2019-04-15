package main

import (
	"bytes"
	"encoding/ascii85"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

// ascii85包实现了ascii85数据编码（5个ascii字符表示4个字节），该编码用于btoa工具和Adobe的PostScript语言和PDF文档格式
func main() {

	// 声明内容
	var src = []byte("Hello World!")

	// 返回字节源数据编码后的最大字节数
	num := ascii85.MaxEncodedLen(len(src))

	// 初始化一个长度为num的[]byte
	var dst = make([]byte, num)

	// 将src编码成最多MaxEncodedLen(len(src))数据写入dst，返回实际写入的ascii字节数
	// 编码每4字节(5个ascii字符)一段进行一次，最后一个片段采用特殊的处理方式，因此不应将本函数用于处理大数据流的某一独立数据块
	ascii85.Encode(dst, src)

	// 转base64字符串输出
	fmt.Println("Ascii85编码内容: ", base64.StdEncoding.EncodeToString(dst))

	// 初始化一个用于接收原始数据的长度为num的[]byte
	var origin = make([]byte, num)

	// 将src解码后写入dst，返回写入dst的ascii字节数、从src解码的ascii字节数。
	// 如果src含有非法数据，函数将返回成功执行的数据（两个数字）和CorruptInputError。
	// 如果flush为true，则函数会认为src代表输入流的结尾，完全处理src，而不会等待另一个32字节的数据块
	_, _, err := ascii85.Decode(origin, dst, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ascii85解码内容: ", string(origin))


	// 声明一个buffer
	var buf bytes.Buffer

	// 创建一个将数据编码为ascii85流写入w的编码器。Ascii85编码算法操作32位块，写入结束后，必须调用Close方法将缓存中保留的不完整块刷新到w里
	w := ascii85.NewEncoder(&buf)
	// 写入
	if _, err := w.Write([]byte("Hello World!")); err != nil {
		log.Fatal(err)
	}
	// 关闭
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	// 转base64字符串输出
	fmt.Println("Ascii85编码内容: ", base64.StdEncoding.EncodeToString(buf.Bytes()))

	// 创建一个从r解码ascii85流的解码器
	r := ascii85.NewDecoder(&buf)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ascii85解码内容: ", string(b))
}