package main

import (
	"bytes"
	"compress/lzw"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const FilePath = "testdata/test.lzw"

func main() {

	// 写lzw数据流
	buf := write()
	// 自动生成并写入文件
	if err := ioutil.WriteFile(FilePath, buf.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
	// 读lzw数据流
	read(FilePath)
}

func write() bytes.Buffer {

	var input = "Hello World!"

	var buf bytes.Buffer

	// lzw: Lempel-Ziv-Welch数据压缩格式
	// 初始化writer
	// lsb表示最低有效位，在gif文件格式中使用。
	// msb表示最高有效位，在tiff和pdf中所用
	// litWidth编码位数，范围[2,8]且通常为8。输入字节必须小于1<<litwidth。
	lw := lzw.NewWriter(&buf, lzw.LSB, 8)

	// 写入
	if _, err := lw.Write([]byte(input)); err != nil {
		log.Fatal(err)
	}

	// 关闭
	if err := lw.Close(); err != nil {
		log.Fatal(err)
	}
	return buf
}

func read(path string) {

	// 读取文件内容
	bf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// 初始化reader
	lr := lzw.NewReader(bytes.NewReader(bf), lzw.LSB, 8)
	defer lr.Close()

	// 读取内容
	b, err := ioutil.ReadAll(lr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}