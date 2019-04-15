package main

import (
	"crypto"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

// 实现了SHA384和SHA512哈希算法
func main() {

	sha384Demo()
	sha512Demo()
	sha512_224Demo()
	sha512_256Demo()
}

func sha384Demo() {

	// 返回一个新的使用SHA384校验的hash.Hash
	h := sha512.New384()
	// 写入
	h.Write([]byte("Hello World"))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	m := h.Sum(nil)
	// 转base64字符串打印
	fmt.Println(base64.StdEncoding.EncodeToString(m))


	// 使用crypto导入
	hash := crypto.SHA384
	h2 := hash.New()
	h2.Write([]byte("Hello World"))
	m2 := h2.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m2))


	// 直接使用sha512.Sum384
	m3 := sha512.Sum384([]byte("Hello World"))
	fmt.Println(base64.StdEncoding.EncodeToString(m3[:]))
}

func sha512Demo() {

	// 返回一个新的使用SHA512校验的hash.Hash
	h := sha512.New()
	// 写入
	h.Write([]byte("Hello World"))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	m := h.Sum(nil)
	// 转base64字符串打印
	fmt.Println(base64.StdEncoding.EncodeToString(m))


	// 使用crypto导入
	hash := crypto.SHA512
	h2 := hash.New()
	h2.Write([]byte("Hello World"))
	m2 := h2.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m2))


	// 直接使用sha512.Sum512
	m3 := sha512.Sum512([]byte("Hello World"))
	fmt.Println(base64.StdEncoding.EncodeToString(m3[:]))
}

func sha512_224Demo() {

	// 返回一个新的使用SHA512/224校验的hash.Hash
	h := sha512.New512_224()
	// 写入
	h.Write([]byte("Hello World"))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	m := h.Sum(nil)
	// 转base64字符串打印
	fmt.Println(base64.StdEncoding.EncodeToString(m))


	// 使用crypto导入
	hash := crypto.SHA512_224
	h2 := hash.New()
	h2.Write([]byte("Hello World"))
	m2 := h2.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m2))


	// 直接使用sha512.Sum512_224
	m3 := sha512.Sum512_224([]byte("Hello World"))
	fmt.Println(base64.StdEncoding.EncodeToString(m3[:]))
}

func sha512_256Demo() {

	// 返回一个新的使用SHA512/256校验的hash.Hash
	h := sha512.New512_256()
	// 写入
	h.Write([]byte("Hello World"))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	m := h.Sum(nil)
	// 转base64字符串打印
	fmt.Println(base64.StdEncoding.EncodeToString(m))


	// 使用crypto导入
	hash := crypto.SHA512_256
	h2 := hash.New()
	h2.Write([]byte("Hello World"))
	m2 := h2.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m2))


	// 直接使用sha512.Sum512_256
	m3 := sha512.Sum512_256([]byte("Hello World"))
	fmt.Println(base64.StdEncoding.EncodeToString(m3[:]))
}