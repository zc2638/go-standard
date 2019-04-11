package main

import (
	"crypto"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

// 实现了SHA1哈希算法
func main() {

	// 返回一个新的使用SHA1校验的hash.Hash
	h := sha1.New()
	// 写入
	h.Write([]byte("Hello World"))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	m := h.Sum(nil)
	// 转base64字符串打印
	fmt.Println(base64.StdEncoding.EncodeToString(m))


	// 使用crypto导入
	hash := crypto.SHA1
	h2 := hash.New()
	h2.Write([]byte("Hello World"))
	m2 := h2.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m2))


	// 直接使用sha1.Sum
	m3 := sha1.Sum([]byte("Hello World"))
	fmt.Println(base64.StdEncoding.EncodeToString(m3[:]))
}