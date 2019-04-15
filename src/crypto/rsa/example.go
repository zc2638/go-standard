package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/zc2638/go-standard/src/crypto/rsa/extra"
	"io/ioutil"
	"log"
	"os"
)

const (
	PublicPemFile = "testdata/rsa_public.pem"
	PublicPKCS1PemFile = "testdata/rsa_public_pkcs1.pem"
	PrivatePemFile = "testdata/rsa_private.pem"
	PrivatePKCS1PemFile = "testdata/rsa_private_pkcs1.pem"
)

// 实现了PKCS#1规定的RSA加密算法
func main() {

	// 生成RSA密钥对
	generateRSAKey()
	// 使用PKCS#1v1.5规定的填充方案和RSA公钥加密/私钥解密
	rsaPKCS1v15Demo()
	// RSA-OAEP算法公钥加密/私钥解密
	rsaOAEPDemo()
	// RSA私钥签名/公钥验证
	rsaSignatureDemo()
	// RSA-PASS私钥签名/公钥验证
	rsaSignPassDemo()


}

func generateRSAKey() {

	// 声明位数
	var bits = 1024

	// 使用随机数据生成器random生成一对具有指定位数的RSA密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	// 使用指定的位数生成一对多质数的RSA密钥。(质数定义为在大于1的自然数中，除了1和它本身以外不再有其他因数)
	// 虽然公钥可以和二质数情况下的公钥兼容（事实上，不能区分两种公钥），私钥却不行。
	// 因此有可能无法生成特定格式的多质数的密钥对，或不能将生成的密钥用在其他（语言的）代码里
	//privateKey, err := rsa.GenerateMultiPrimeKey(rand.Reader, 5, bits)

	if err != nil {
		log.Fatal(err)
	}

	// 将rsa私钥序列化为ASN.1 PKCS#1 DER编码
	derPrivate := x509.MarshalPKCS1PrivateKey(privateKey)
	// 初始化一个PEM编码的结构
	priBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derPrivate,
	}
	// 创建文件，如果文件存在内容重置为空
	file, err := os.Create(PrivatePKCS1PemFile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 将Block的pem编码写入文件
	err = pem.Encode(file, priBlock)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RSA-PKCS1私钥生成成功")


	// 将rsa私钥序列化为PKCS#8 DER编码
	derPrivate8, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 初始化一个PEM编码的结构
	priBlock8 := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derPrivate8,
	}

	// 创建文件，如果文件存在内容重置为空
	file2, err := os.Create(PrivatePemFile)
	defer file2.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 将Block的pem编码写入文件
	err = pem.Encode(file2, priBlock8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RSA-PKCS8私钥生成成功")

	// 获取公钥
	publicKey := &privateKey.PublicKey

	// 将公钥序列化为PKIX格式DER编码
	derPublic, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatal(err)
	}
	// 初始化一个PEM编码的结构
	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPublic,
	}

	// 创建文件，如果文件存在内容重置为空
	file3, err := os.Create(PublicPemFile)
	defer file3.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 将Block的pem编码写入文件
	err = pem.Encode(file3, pubBlock)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RSA-PKIX公钥生成成功")


	derPublic1 := x509.MarshalPKCS1PublicKey(publicKey)
	// 初始化一个PEM编码的结构
	pubBlock1 := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPublic1,
	}
	// 创建文件，如果文件存在内容重置为空
	file4, err := os.Create(PublicPKCS1PemFile)
	defer file4.Close()
	if err != nil {
		log.Fatal(err)
	}
	// 将Block的pem编码写入文件
	err = pem.Encode(file4, pubBlock1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RSA-PKCS1公钥生成成功")
}

func getRSAKey() (privateKey []byte, publicKey []byte, err error) {

	// 读取publicKey内容
	publicKey, err = ioutil.ReadFile(PublicPemFile)
	if err != nil {
		return
	}
	// 读取privateKey内容
	privateKey, err = ioutil.ReadFile(PrivatePemFile)
	return
}

func rsaPKCS1v15Demo() {

	// 获取公钥私钥
	privateKey, publicKey, err := getRSAKey()
	if err != nil {
		log.Fatal(err)
	}

	// 声明内容
	var origin = []byte("Hello World!")

	// rsa公钥加密
	cipherText, err := extra.Encrypt(publicKey, origin)
	if err != nil {
		log.Fatal(err)
	}

	// 转base64字符串
	cipherTextStr := base64.StdEncoding.EncodeToString(cipherText)
	fmt.Println("RSA-PKCS1v15加密内容: ", cipherTextStr)

	// rsa私钥解密
	originText, err := extra.Decrypt(privateKey, cipherText)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RSA-PKCS1v15解密内容: ", string(originText))
}

func rsaOAEPDemo() {

	// 获取公钥私钥
	privateKey, publicKey, err := getRSAKey()
	if err != nil {
		log.Fatal(err)
	}

	// 声明内容
	var origin = []byte("Hello World!")

	// 声明label
	var label = []byte("test")

	// rsa公钥加密
	cipherText, err := extra.EncryptOAEP(publicKey, origin, label)
	if err != nil {
		log.Fatal(err)
	}

	// 转base64字符串
	cipherTextStr := base64.StdEncoding.EncodeToString(cipherText)
	fmt.Println("RSA-OAEP加密内容: ", cipherTextStr)

	// rsa私钥解密
	originText, err := extra.DecryptOAEP(privateKey, cipherText, label)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RSA-OAEP解密内容: ", string(originText))
}

func rsaSignatureDemo() {

	// 获取公钥私钥
	privateKey, publicKey, err := getRSAKey()
	if err != nil {
		log.Fatal(err)
	}

	// 声明内容
	var origin = []byte("Hello World!")

	// rsa私钥签名
	signature, err := extra.Sign(privateKey, origin)
	if err != nil {
		log.Fatal(err)
	}

	// 转base64字符串
	signatureStr := base64.StdEncoding.EncodeToString(signature)
	fmt.Println("RSA签名内容: ", signatureStr)

	//rsa公钥验签
	err = extra.Verify(publicKey, origin, signature)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RSA签名验证成功")
}

func rsaSignPassDemo() {

	// 获取公钥私钥
	privateKey, publicKey, err := getRSAKey()
	if err != nil {
		log.Fatal(err)
	}

	// 声明内容
	var origin = []byte("Hello World!")

	// 初始化一个 PSS签名 的参数
	var SignOpts = rsa.PSSOptions{SaltLength: 8}

	// 初始化一个 PSS认证 的参数
	var VerifyOpts = rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto}
	// 当PSS签名参数SaltLength 为 rsa.PSSSaltLengthAuto, PSS认证参数SaltLength 必须为 rsa.PSSSaltLengthAuto
	// 当PSS签名参数SaltLength 为 rsa.PSSSaltLengthEqualsHash, PSS认证参数SaltLength 为 rsa.PSSSaltLengthAuto或rsa.PSSSaltLengthEqualsHash
	// 当PSS签名参数SaltLength 为 指定值时 如8, PSS认证参数SaltLength 为rsa.PSSSaltLengthAuto或8
	//

	// rsa私钥签名
	signature, err := extra.SignPass(privateKey, origin, &SignOpts)
	if err != nil {
		log.Fatal(err)
	}

	// 转base64字符串
	signatureStr := base64.StdEncoding.EncodeToString(signature)
	fmt.Println("RSA-PASS签名内容: ", signatureStr)

	//rsa公钥验签
	err = extra.VerifyPass(publicKey, origin, signature, &VerifyOpts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RSA-PASS签名验证成功")
}
