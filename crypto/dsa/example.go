package main

import (
	"crypto/dsa"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

// 实现FIPS 186-3定义的数字签名算法(DSA算法)
func main() {

	// 声明一个DSA参数中的质数可以接受的字位长度的枚举
	var size = dsa.L1024N160

	// 声明新的dsa.PrivateKey
	var pri dsa.PrivateKey

	// 获取密钥域的参数的地址
	params := &pri.Parameters

	// 随机设置合法的参数到params。即使机器很快，函数也可能会花费很多时间来生成参数
	err := dsa.GenerateParameters(params, rand.Reader, size)
	if err != nil {
		log.Fatal(err)
	}

	if params.P.BitLen() != 1024 {
		log.Fatalf("%d: params.BitLen got:%d want:%d\n", int(size), params.P.BitLen(), 1024)
	}

	if params.Q.BitLen() != 160 {
		log.Fatalf("%d: q.BitLen got:%d want:%d\n", int(size), params.Q.BitLen(), 160)
	}

	// 声明一个高精度的数值
	one := new(big.Int).SetInt64(1)
	// 声明一个params.P减去one值的高精度数值
	pm1 := new(big.Int).Sub(params.P, one)

	// 如果y != 0将z设为x/y，将m设为 x%y 并返回(z, m)；如果y == 0会panic
	quo, rem := new(big.Int).DivMod(pm1, params.Q, new(big.Int))
	if rem.Sign() != 0 {
		log.Fatalf("%d: p-1 mod q != 0\n", int(size))
	}

	// 将z设为x**y mod |m|并返回z；如果y <= 0，返回1；如果m == nil 或 m == 0，z设为x**y
	x := new(big.Int).Exp(params.G, quo, params.P)
	if x.Cmp(one) == 0 {
		log.Fatalf("%d: invalid generator\n", int(size))
	}

	// 生成一对公私钥
	err = dsa.GenerateKey(&pri, rand.Reader)
	if err != nil {
		log.Fatalf("error generating key: %s\n", err)
	}

	// 声明签名内容
	var origin = []byte("Hello World!")

	// 使用私钥对任意长度的hash值（必须是较大信息的hash结果）进行签名，返回签名结果（一对大整数）
	r, s, err := dsa.Sign(rand.Reader, &pri, origin)
	if err != nil {
		log.Fatalf("%d: error signing: %s\n", int(size), err)
	}
	fmt.Println("DSA签名内容:", r, s)

	// 使用公钥验证hash值和两个大整数r、s构成的签名，并返回签名是否合法
	if !dsa.Verify(&pri.PublicKey, origin, r, s) {
		log.Fatalf("%d: Verify failed\n", int(size))
	}
	fmt.Println("DSA签名验证成功")
}
