package extra

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
)

func Encrypt(publicKey, originText []byte) ([]byte, error) {

	// 获取rsa.PublicKey
	pub, err := BuildRSAPublicKey(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	// 使用PKCS#1 v1.5规定的填充方案和RSA算法加密msg。信息不能超过((公共模数的长度)-11)字节
	// 注意：使用本函数加密明文（而不是会话密钥）是危险的，请尽量在新协议中使用RSA OAEP
	return rsa.EncryptPKCS1v15(rand.Reader, pub, originText)
}

func Decrypt(privateKey, cipherText []byte) ([]byte, error) {

	// 获取rsa.PrivateKey
	pri, err := BuildRSAPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	// 使用PKCS#1 v1.5规定的填充方案和RSA算法解密密文。如果random不是nil，函数会注意规避时间侧信道攻击
	return rsa.DecryptPKCS1v15(rand.Reader, pri, cipherText)
}

func Sign(privateKey, originText []byte) ([]byte, error) {

	// 获取rsa.PrivateKey
	pri, err := BuildRSAPrivateKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 返回数据的 SHA256 校验和
	hashed := sha256.Sum256(originText)

	// 使用RSA PKCS#1 v1.5规定的RSASSA-PKCS1-V1_5-SIGN签名方案计算签名
	// 注意hashed必须是使用提供给本函数的hash参数对（要签名的）原始数据进行hash的结果
	return rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA256, hashed[:])
}

func Verify(publicKey, originText, signature []byte) error {

	// 获取rsa.PublicKey
	pub, err := BuildRSAPublicKey(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	// 返回数据的 SHA256 校验和
	hashed := sha256.Sum256(originText)

	// 验证 RSA PKCS＃1 v1.5 签名
	// hashed是使用提供的hash参数对（要签名的）原始数据进行hash的结果。合法的签名会返回nil，否则表示签名不合法
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}
