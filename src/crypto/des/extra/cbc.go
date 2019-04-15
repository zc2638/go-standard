package extra

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func CBCEncrypt(originText, key, iv []byte, triple bool) ([]byte, error) {

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

	// 返回加密字节块的大小
	blockSize := block.BlockSize()

	// PKCS5填充需加密内容
	originText = PKCS5Padding(originText, blockSize)

	// 返回一个密码分组链接模式的、底层用Block加密的cipher.BlockMode，初始向量iv的长度必须等于Block的块尺寸(Block块尺寸等于密钥尺寸)
	blockMode := cipher.NewCBCEncrypter(block, iv)

	// 根据 需加密内容[]byte长度,初始化一个新的byte数组，返回byte数组内存地址
	cipherText := make([]byte, len(originText))

	// 加密或解密连续的数据块，将加密内容存储到dst中，src需加密内容的长度必须是块大小的整数倍，src和dst可指向同一内存地址
	blockMode.CryptBlocks(cipherText, originText)

	return cipherText, nil
}

func CBCDecrypt(cipherText, key, iv []byte, triple bool) ([]byte, error) {

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

	// 返回一个密码分组链接模式的、底层用b解密的cipher.BlockMode，初始向量iv必须和加密时使用的iv相同
	blockMode := cipher.NewCBCDecrypter(block, iv)

	// 根据 密文[]byte长度,初始化一个新的byte数组，返回byte数组内存地址
	originText := make([]byte, len(cipherText))

	// 加密或解密连续的数据块，将解密内容存储到dst中，src需加密内容的长度必须是块大小的整数倍，src和dst可指向同一内存地址
	blockMode.CryptBlocks(originText, cipherText)

	// PKCS5反填充解密内容
	originText = PKCS5UnPadding(originText)
	return originText, nil
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {

	// 获取 需加密字符串长度 和 加密字节块长度 的余数
	less := len(cipherText) % blockSize
	// 求填充长度
	padding := blockSize - less
	// 将填充长度 重复 长度padding次，返回填充内容
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 将填充内容 添加到 需加密内容
	return append(cipherText, padText...)
}

func PKCS5UnPadding(originText []byte) []byte {

	// 获取 解密内容长度
	length := len(originText)
	// 获取反填充长度(只获取最后个byte当做长度，因为填充的时候是重复按照长度填充的)
	unPadding := int(originText[length-1])
	// 截取解密内容中的原文内容
	return originText[:(length - unPadding)]
}