package main

import (
	"crypto"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// 实现了SHA224和SHA256哈希算法
func main() {

	sha224Demo()
	sha256Demo()
}

func sha224Demo() {

	// 返回一个新的使用SHA224校验的hash.Hash
	h := sha256.New224()
	// 写入
	h.Write([]byte("Hello World"))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	m := h.Sum(nil)
	// 转base64字符串打印
	fmt.Println(base64.StdEncoding.EncodeToString(m))


	// 使用crypto导入
	hash := crypto.SHA224
	h2 := hash.New()
	h2.Write([]byte("Hello World"))
	m2 := h2.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m2))


	// 直接使用sha256.Sum224
	m3 := sha256.Sum224([]byte("Hello World"))
	fmt.Println(base64.StdEncoding.EncodeToString(m3[:]))
}

func sha256Demo() {

	// 返回一个新的使用SHA256校验的hash.Hash
	h := sha256.New()
	// 写入
	h.Write([]byte("Hello World"))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	m := h.Sum(nil)
	// 转base64字符串打印
	fmt.Println(base64.StdEncoding.EncodeToString(m))


	// 使用crypto导入
	hash := crypto.SHA256
	h2 := hash.New()
	h2.Write([]byte("Hello World"))
	m2 := h2.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m2))


	// 直接使用sha256.Sum256
	m3 := sha256.Sum256([]byte("Hello World"))
	fmt.Println(base64.StdEncoding.EncodeToString(m3[:]))
}