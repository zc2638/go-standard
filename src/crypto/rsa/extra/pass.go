package extra

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"log"
)

func SignPass(privateKey, originText []byte, opts *rsa.PSSOptions) ([]byte, error) {

	// 获取rsa.PrivateKey
	pri, err := BuildRSAPrivateKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 声明MD5 Hash类型
	hash := crypto.MD5
	// 根据Hash类型创建hash
	h := hash.New()
	// 写入内容
	h.Write(originText)
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	hashed := h.Sum(nil)

	// 采用RSASSA-PSS方案计算签名
	// 注意hashed必须是使用提供给本函数的hash参数对（要签名的）原始数据进行hash的结果
	// opts参数可以为nil，此时会使用默认参数
	return rsa.SignPSS(rand.Reader, pri, hash, hashed, opts)
}

func VerifyPass(publicKey, originText, signature []byte, opts *rsa.PSSOptions) error {

	// 获取rsa.PublicKey
	pub, err := BuildRSAPublicKey(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	// 声明MD5 Hash类型
	hash := crypto.MD5
	// 根据Hash类型创建hash
	h := hash.New()
	// 写入内容
	h.Write(originText)
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	hashed := h.Sum(nil)

	// 认证一个PSS签名
	// hashed是使用提供给本函数的hash参数对（要签名的）原始数据进行hash的结果。合法的签名会返回nil，否则表示签名不合法
	// opts参数可以为nil，此时会使用默认参数
	return rsa.VerifyPSS(pub, hash, hashed, signature, opts)
}