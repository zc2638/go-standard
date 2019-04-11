package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

// 实现了用于加解密的更安全的随机数生成器
func main() {

	// 声明一个指定长度的byte切片地址
	b := make([]byte, 10)
	// 随机填充b。本函数是一个使用io.ReadFull调用Reader.Read的辅助性函数。当且仅当err == nil时，返回值n == len(b)
	if _, err := rand.Read(b); err != nil {
		log.Fatal(err)
	}

	// 返回一个统一的随机值，最大不会超过设置的max值，设置的max值必须大于0
	// rand.Reader是一个全局、共享的密码用强随机生成器
	i, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i.Bytes())

	// 返回一个具有指定字位数的数字，该数字具有很高可能性是质数。如果从rand读取时出错，或者bits<2会返回错误
	p, err := rand.Prime(rand.Reader, 11)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.Bytes())
}
