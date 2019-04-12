package main

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"io/ioutil"
	"log"
)

const FilePath = "testdata/test.bz2"

// bzip2包实现bzip2的解压缩
func main() {

	// 读取文件内容
	bf, err := ioutil.ReadFile(FilePath)
	if err != nil {
		log.Fatal(err)
	}

	// bzip2解压缩成reader
	r := bzip2.NewReader(bytes.NewReader(bf))

	// 读取内容
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("bzip2 content: ", string(b))
}