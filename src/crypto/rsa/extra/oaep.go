package extra

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
)

func EncryptOAEP(publicKey, originText, label []byte) ([]byte, error) {

	// 获取rsa.PublicKey
	pub, err := BuildRSAPublicKey(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	// 采用RSA-OAEP算法加密指定数据。数据不能超过((公共模数的长度)-2*( hash长度)+2)字节
	// label参数可能包含不加密的任意数据，但这给了信息重要的背景。例如，如果给定公钥用于解密两种类型的消息，然后是不同的标签值可用于确保用于一个目的的密文不能 被攻击者用于另一个目的。如果不需要，可以为空
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, originText, label)
}

func DecryptOAEP(privateKey, cipherText, label []byte) ([]byte, error) {

	// 获取rsa.PrivateKey
	pri, err := BuildRSAPrivateKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 使用PKCS#1 v1.5规定的填充方案和RSA算法解密密文。如果random不是nil，函数会注意规避时间侧信道攻击
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, pri, cipherText, label)
}

func BuildRSAPublicKey(publicKey []byte) (*rsa.PublicKey, error) {

	// 返回解码得到的pem.Block和剩余未解码的数据。如果未发现PEM数据，返回(nil, data)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	// 解析一个DER编码的公钥。这些公钥一般在以"BEGIN PUBLIC KEY"出现的PEM块中
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 指定为rsa.PublicKey结构
	pub := pubInterface.(*rsa.PublicKey)

	return pub, nil
}

func BuildRSAPrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {

	// 返回解码得到的pem.Block和剩余未解码的数据。如果未发现PEM数据，返回(nil, data)
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	// 解析一个未加密的PKCS#8私钥
	priInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 指定为rsa.PrivateKey结构
	pri := priInterface.(*rsa.PrivateKey)

	return pri, nil
}

func BuildRSAPKCS1PublicKey(publicKey []byte) (*rsa.PublicKey, error) {

	// 返回解码得到的pem.Block和剩余未解码的数据。如果未发现PEM数据，返回(nil, data)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	// 解析一个ASN.1 PKCS#1 DER编码的公钥。
	return x509.ParsePKCS1PublicKey(block.Bytes)
}

func BuildRSAPKCS1PrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {

	// 返回解码得到的pem.Block和剩余未解码的数据。如果未发现PEM数据，返回(nil, data)
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	// 解析一个ASN.1 PKCS#1 DER编码的私钥。
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}