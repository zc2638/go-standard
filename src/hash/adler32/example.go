package main

import (
	"encoding/hex"
	"fmt"
	"hash/adler32"
	"log"
)

// adler32包实现了Adler-32校验和算法
// Adler-32由两个每字节累积的和组成：s1是所有字节的累积，s2是所有s1的累积。两个累积值都取65521的余数。s1初始为1，s2初始为0。Afler-32校验和保存为s2*65536 + s1。（最高有效字节在前/大端在前）
func main() {

	// 声明内容
	var data = []byte("Hello World!")

	// 返回数据data的Adler-32校验和
	cipherText := adler32.Checksum(data)
	fmt.Println(cipherText)


	// 返回一个计算Adler-32校验和的hash.Hash32接口
	h := adler32.New()
	// 写入
	if _, err := h.Write(data); err!= nil {
		log.Fatal(err)
	}
	hb := h.Sum(nil)
	fmt.Println(hex.EncodeToString(hb))
}