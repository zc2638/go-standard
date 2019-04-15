package extra

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"errors"
)

func CFBEncrypt(originText, key, iv []byte, triple bool) ([]byte, error) {

	var block cipher.Block
	var err error
	if triple {
		// 创建一个cipher.Block。参数key为24字节密钥
		block, err = des.NewTripleDESCipher(key)
	} else {
		// 创建一个cipher.Block。参数key为8字节密钥
		block, err = des.NewCipher(key)
	}
	if err != nil {
		return nil, err
	}

	// 根据 需加密内容[]byte长度,初始化一个新的byte数组，返回byte数组内存地址
	cipherText := make([]byte, aes.BlockSize+len(originText))

	// 返回一个密码反馈模式的、底层用block加密的cipher.Stream，初始向量iv的长度必须等于block的块尺寸
	stream := cipher.NewCFBEncrypter(block, iv)

	// 从加密器的key流和src中依次取出字节二者xor后写入dst，src和dst可指向同一内存地址
	// cipherText[:aes.BlockSize]为iv值，所以只写入cipherText后面部分
	stream.XORKeyStream(cipherText[aes.BlockSize:], originText)

	return cipherText, nil
}

func CFBDecrypt(cipherText, key, iv []byte, triple bool) ([]byte, error) {

	var block cipher.Block
	var err error
	if triple {
		// 创建一个cipher.Block。参数key为24字节密钥
		block, err = des.NewTripleDESCipher(key)
	} else {
		// 创建一个cipher.Block。参数key为8字节密钥
		block, err = des.NewCipher(key)
	}
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("cipherText too short")
	}

	// 只使用cipherText除去iv部分
	cipherText = cipherText[aes.BlockSize:]

	// 返回一个密码反馈模式的、底层用block解密的cipher.Stream，初始向量iv必须和加密时使用的iv相同
	stream := cipher.NewCFBDecrypter(block, iv)

	// 从加密器的key流和src中依次取出字节二者xor后写入dst，src和dst可指向同一内存地址
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}