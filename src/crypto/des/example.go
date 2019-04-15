package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/zc2638/go-standard/src/crypto/des/extra"
	"log"
)

// 实现了DES标准和TDEA算法
func main() {

	// DES-CBC加密/解密
	CBC()
	// DES-CFB加密/解密
	CFB()
	// DES-CTR加密/解密
	CTR()
	// DES-OFB加密/解密
	OFB()
	// DES-OFB加密/解密，使用cipher的StreamReader加密、cipher的StreamWriter解密
	OFBStream()

	CBCTriple()
	CFBTriple()
	CTRTriple()
	OFBTriple()
	OFBStreamTriple()
}

func CBC() {

	// 声明一个8字节的key
	var key = []byte("test key")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-cbc encode test text")
	// 声明一个8字节的iv
	var iv = []byte("test ivs")

	// 加密
	cipherText, err := extra.CBCEncrypt(origin, key, iv, false)
	if err != nil {
		log.Fatal(err)
	}

	// byte转base64字符串
	cipherTextStr := base64.StdEncoding.EncodeToString(cipherText)
	fmt.Println("DES-CBC加密内容: ", cipherTextStr)

	// 解密
	originText, err := extra.CBCDecrypt(cipherText, key, iv, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DES-CBC解密内容: ", string(originText))
}

func CFB() {

	// 声明一个8字节的key
	var key = []byte("12345678")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-cfb encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := extra.CFBEncrypt(origin, key, iv, false)
	if err != nil {
		log.Fatal(err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	fmt.Println("DES-CFB加密内容: ", cipherTextStr)

	// 解密
	originText, err := extra.CFBDecrypt(cipherText, key, iv, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DES-CFB解密内容: ", string(originText))
}

func CTR() {

	// 声明一个8字节的key
	var key = []byte("12345678")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-ctr encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := extra.CTREncrypt(origin, key, iv, false)
	if err != nil {
		log.Fatal(err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	fmt.Println("DES-CTR加密内容: ", cipherTextStr)

	// 解密
	originText, err := extra.CTRDecrypt(cipherText, key, iv, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DES-CTR解密内容: ", string(originText))
}

func OFB() {

	// 声明一个8字节的key
	var key = []byte("12345678")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-ofb encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := extra.OFBEncrypt(origin, key, iv, false)
	if err != nil {
		log.Fatal(err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	fmt.Println("DES-OFB加密内容: ", cipherTextStr)

	// 解密
	originText, err := extra.OFBDecrypt(cipherText, key, iv, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DES-OFB解密内容: ", string(originText))
}

func OFBStream() {

	// 声明一个8字节的key
	var key = []byte("12345678")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-ofb-stream encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// StreamReader方式加密
	cipherText, err := extra.OFBEncryptStreamReader(origin, key, iv, false)
	if err != nil {
		log.Fatal(err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	fmt.Println("DES-OFB-Stream方式加密内容: ", cipherTextStr)

	// StreamWriter方式解密
	originText, err := extra.OFBDecryptStreamWriter(cipherText, key, iv, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DES-OFB-Stream方式解密内容: ", string(originText))
}

func CBCTriple() {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-cbc-triple encode test text")
	// 声明一个8字节的iv
	var iv = []byte("test ivs")

	// 加密
	cipherText, err := extra.CBCEncrypt(origin, key, iv, true)
	if err != nil {
		log.Fatal(err)
	}

	// byte转base64字符串
	cipherTextStr := base64.StdEncoding.EncodeToString(cipherText)
	fmt.Println("DES-CBC-Triple加密内容: ", cipherTextStr)

	// 解密
	originText, err := extra.CBCDecrypt(cipherText, key, iv, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DES-CBC-Triple解密内容: ", string(originText))
}

func CFBTriple() {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-cfb-triple encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := extra.CFBEncrypt(origin, key, iv, true)
	if err != nil {
		log.Fatal(err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	fmt.Println("DES-CFB-Triple加密内容: ", cipherTextStr)

	// 解密
	originText, err := extra.CFBDecrypt(cipherText, key, iv, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DES-CFB-Triple解密内容: ", string(originText))
}

func CTRTriple() {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-ctr-triple encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := extra.CTREncrypt(origin, key, iv, true)
	if err != nil {
		log.Fatal(err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	fmt.Println("DES-CTR-Triple加密内容: ", cipherTextStr)

	// 解密
	originText, err := extra.CTRDecrypt(cipherText, key, iv, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DES-CTR-Triple解密内容: ", string(originText))
}

func OFBTriple() {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-ofb-triple encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := extra.OFBEncrypt(origin, key, iv, true)
	if err != nil {
		log.Fatal(err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	fmt.Println("DES-OFB-Triple加密内容: ", cipherTextStr)

	// 解密
	originText, err := extra.OFBDecrypt(cipherText, key, iv, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DES-OFB-Triple解密内容: ", string(originText))
}

func OFBStreamTriple() {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-ofb-triple-stream encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// StreamReader方式加密
	cipherText, err := extra.OFBEncryptStreamReader(origin, key, iv, true)
	if err != nil {
		log.Fatal(err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	fmt.Println("DES-OFB-Triple-Stream方式加密内容: ", cipherTextStr)

	// StreamWriter方式解密
	originText, err := extra.OFBDecryptStreamWriter(cipherText, key, iv, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DES-OFB-Triple-Stream方式解密内容: ", string(originText))
}