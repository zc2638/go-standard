package extra

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

func GCMEncrypt(originText, key, nonce []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// 函数用迦洛瓦计数器模式包装提供的128位Block接口，并返回cipher.AEAD
	g, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 返回加密结果。认证附加的additionalData，将加密结果添加到dst生成新的加密结果，nonce的长度必须是NonceSize()字节，且对给定的key和时间都是独一无二的
	cipherText := g.Seal(nil, nonce, originText, nil)

	return cipherText, nil
}

func GCMDecrypt(cipherText, key, nonce []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 函数用迦洛瓦计数器模式包装提供的128位Block接口，并返回cipher.AEAD
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 返回解密结果。认证附加的additionalData，将解密结果添加到dst生成新的加密结果，nonce的长度必须是NonceSize()字节，nonce和data都必须和加密时使用的相同
	return aesgcm.Open(nil, nonce, cipherText, nil)
}
