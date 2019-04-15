package main

import (
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"log"
)

// 实现了RC4加密算法
func main() {

	// 声明密钥
	var key = []byte("example key")

	// 声明内容
	var text = []byte("Hello World!")

	// 创建并返回一个新的Cipher。参数key是RC4密钥
	c, err := rc4.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// 加密。将src的数据与秘钥生成的伪随机位流取XOR并写入dst。dst和src可指向同一内存地址；但如果指向不同则其底层内存不可重叠
	c.XORKeyStream(text, text)
	// byte转十六进制字符串
	textStr := hex.EncodeToString(text)
	fmt.Println(textStr)
}