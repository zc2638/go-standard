package main

import (
	"crypto/hmac"
	"crypto/sha256"
)

// 实现了U.S. Federal Information Processing Standards Publication 198规定的HMAC（加密哈希信息认证码）
func main() {

	// 声明密钥
	var key = []byte("test hmac key")

	// 声明随意内容
	var content = []byte("Hello World!")

	// 返回一个采用hash.Hash作为底层hash接口、key作为密钥的HMAC算法的hash接口
	h := hmac.New(sha256.New, key)
	// 写入
	h.Write(content)
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	expectedMAC := h.Sum(nil)

	// 比较两个MAC是否相同
	hmac.Equal(expectedMAC, expectedMAC)
}