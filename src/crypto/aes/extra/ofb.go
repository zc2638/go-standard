package extra

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"io/ioutil"
	"log"
)

func OFBEncrypt(originText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 根据 需加密内容[]byte长度,初始化一个新的byte数组，返回byte数组内存地址
	cipherText := make([]byte, aes.BlockSize+len(originText))

	// 返回一个输出反馈模式的、底层采用b生成key流的cipher.Stream，初始向量iv的长度必须等于b的块尺寸
	stream := cipher.NewOFB(block, iv)

	// 从加密器的key流和src中依次取出字节二者xor后写入dst，src和dst可指向同一内存地址
	// cipherText[:aes.BlockSize]为iv值，所以只写入cipherText后面部分
	stream.XORKeyStream(cipherText[aes.BlockSize:], originText)

	return cipherText, nil
}

func OFBDecrypt(cipherText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 只使用cipherText除去iv部分
	cipherText = cipherText[aes.BlockSize:]

	// 返回一个输出反馈模式的、底层采用b生成key流的cipher.Stream，初始向量iv的长度必须等于b的块尺寸
	stream := cipher.NewOFB(block, iv)

	// 从加密器的key流和src中依次取出字节二者xor后写入dst，src和dst可指向同一内存地址
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}

func OFBEncryptStreamReader(originText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 返回一个输出反馈模式的、底层采用b生成key流的cipher.Stream，初始向量iv的长度必须等于b的块尺寸
	stream := cipher.NewOFB(block, iv)

	// 初始化cipher.StreamReader。将一个cipher.Stream与一个io.Reader关联起来，Read方法会调用XORKeyStream方法来处理获取的所有切片
	reader := &cipher.StreamReader{
		S: stream,
		R: bytes.NewReader(originText),
	}

	return ioutil.ReadAll(reader)
}

func OFBDecryptStreamWriter(cipherText, key, iv []byte) ([]byte, error) {

	// 创建一个cipher.Block。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 返回一个输出反馈模式的、底层采用b生成key流的cipher.Stream，初始向量iv的长度必须等于b的块尺寸
	stream := cipher.NewOFB(block, iv)

	// 声明buffer
	var originText bytes.Buffer

	// 初始化cipher.StreamWriter。将一个cipher.Stream与一个io.Writer接口关联起来，Write方法会调用XORKeyStream方法来处理提供的所有切片
	// 如果Write方法返回的n小于提供的切片的长度，则表示StreamWriter不同步，必须丢弃。StreamWriter没有内建的缓存，不需要调用Close方法去清空缓存
	writer := &cipher.StreamWriter{
		S: stream,
		W: &originText,
	}

	// 把reader内容拷贝到writer, writer会调用write方法写入内容
	if _, err := io.Copy(writer, bytes.NewReader(cipherText)); err != nil {
		log.Fatal(err)
	}

	return originText.Bytes(), nil
}
