package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

// 实现了几条覆盖素数有限域的标准椭圆曲线
func main() {

	// 返回一个实现了P-224的曲线
	elliptic.P224()
	// 返回一个实现了P-256的曲线
	elliptic.P256()
	// 返回一个实现了P-384的曲线
	elliptic.P384()
	// 返回一个实现了P-521的曲线
	curve := elliptic.P521()

	// 返回一个公钥/私钥对。priv是私钥，而(x,y)是公钥。密钥对是通过提供的随机数读取器来生成的，该io.Reader接口必须返回随机数据
	priv, x, y, err := elliptic.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(priv))
	fmt.Println(x)
	fmt.Println(y)

	// 将一个点编码为ANSI X9.62指定的格式
	d := elliptic.Marshal(curve, x, y)

	// 将一个Marshal编码后的点还原；如果出错，x会被设为nil
	elliptic.Unmarshal(curve, d)
}