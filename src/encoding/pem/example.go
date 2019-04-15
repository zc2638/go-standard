package main

import (
	"bytes"
	"encoding/pem"
	"fmt"
	"log"
)

// pem包实现了PEM数据编码（源自保密增强邮件协议）。目前PEM编码主要用于TLS密钥和证书
func main() {

	// 声明buffer
	var buf bytes.Buffer

	// 初始化一个PEM编码的结构
	block := &pem.Block{
		Type:    "PUBLIC KEY",   // 类型
		Headers: map[string]string{
			"author": "zc",
			"name": "Gopher",
		},            // 可选的头项
		Bytes:   []byte("test"), // 内容解码后的数据，一般是DER编码的ASN.1结构
	}

	// 将Block的pem编码写入writer
	if err := pem.Encode(&buf, block); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf.Bytes()))

	// pem解码返回一个PEM结构 和 未解码数据
	dk, _ := pem.Decode(buf.Bytes())
	fmt.Println(dk)
}