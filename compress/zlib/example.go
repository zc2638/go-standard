package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const FilePath = "testdata/test.zlib"

// zlib包实现了对zlib格式压缩数据的读写
func main() {

	// 写zlib数据流
	buf := write()
	// 自动生成并写入文件
	if err := ioutil.WriteFile(FilePath, buf.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
	// 读zlib数据流
	read(FilePath)
}

func write() bytes.Buffer {

	input := "Hello World!"

	// 声明buffer
	var buf bytes.Buffer

	// 初始化writer
	zw := zlib.NewWriter(&buf)

	// 写入
	if _, err := zw.Write([]byte(input)); err != nil {
		log.Fatal(err)
	}

	// 关闭
	if err := zw.Close(); err != nil {
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
	zr, err := zlib.NewReader(bytes.NewReader(bf))
	defer zr.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 读取内容
	b, err := ioutil.ReadAll(zr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}