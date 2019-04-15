package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
)

// 实现了椭圆曲线数字签名算法
func main() {

	// 生成一对公钥/私钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	// 声明签名内容
	msg := "hello, world"

	// 返回数据的SHA256校验和
	hash := sha256.Sum256([]byte(msg))

	// 使用私钥对任意长度的hash值（必须是较大信息的hash结果）进行签名，返回签名结果（一对大整数）。私钥的安全性取决于密码读取器的熵度（随机程度）
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("signature: (0x%x, 0x%x)\n", r, s)

	// 使用公钥验证hash值和两个大整数r、s构成的签名，并返回签名是否合法
	valid := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
	fmt.Println("signature verified:", valid)
}