package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const FilePath = "testdata/test.deflate"
const DictFilePath = "testdata/testDict.deflate"

// flate包实现了deflate压缩数据格式。gzip包和zlib包实现了对基于deflate的文件格式的访问
func main() {

	// 写deflate数据流
	buf := write()
	// 自动生成并写入文件
	if err := ioutil.WriteFile(FilePath, buf.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
	// 读deflate数据流
	read(FilePath)

	const dict = `<?xml version="1.0"?><book></book><data></data><meta name="" content="`
	// 预设字典写deflate数据流
	dictBuf := writeDict(dict)
	// 自动生成并写入文件
	if err := ioutil.WriteFile(DictFilePath, dictBuf.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
	// 根据预设字典读deflate数据流
	readDict(dict, DictFilePath)
}

func write() bytes.Buffer {

	inputs := []string{
		"Don't communicate by sharing memory, share memory by communicating.\n",
		"Concurrency is not parallelism.\n",
		"The bigger the interface, the weaker the abstraction.\n",
		"Documentation is for users.\n",
	}
	resetInput := "This is a reset test message!"

	var buf bytes.Buffer

	// 初始化writer，可设置压缩类型
	fw, err := flate.NewWriter(&buf, flate.DefaultCompression)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range inputs {
		// 写入
		if _, err := fw.Write([]byte(v)); err != nil {
			log.Fatal(err)
		}
		// 挂起写入数据，重置缓冲区
		if err := fw.Flush(); err != nil {
			log.Fatal(err)
		}
	}

	// 重置buffer内容为空
	fw.Reset(&buf)
	fw.Write([]byte(resetInput))

	// 关闭writer
	if err := fw.Close(); err != nil {
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
	r := flate.NewReader(bytes.NewReader(bf))
	if err := r.Close(); err != nil {
		log.Fatal(err)
	}

	// 读取内容
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func writeDict(dict string) bytes.Buffer {

	const data = `<?xml version="1.0"?>
<book>
	<meta name="title" content="The Go Programming Language"/>
	<meta name="authors" content="Alan Donovan and Brian Kernighan"/>
	<meta name="published" content="2015-10-26"/>
	<meta name="isbn" content="978-0134190440"/>
	<data>...</data>
</book>`

	var buf bytes.Buffer

	// 初始化一个预设字典的writer
	fw, err := flate.NewWriterDict(&buf, flate.DefaultCompression, []byte(dict))
	defer fw.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 写入
	if _, err := fw.Write([]byte(data)); err != nil {
		log.Fatal(err)
	}
	// 关闭
	if err := fw.Close(); err != nil {
		log.Fatal(err)
	}
	return buf
}

func readDict(dict string, path string) {

	bf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// 可以读取的时候替换字典内容，但是必须保证长度相同，注释测试原字典解压
	hashDict := []byte(dict)
	for i := range hashDict {
		hashDict[i] = '#'
	}

	// 初始化一个预设字典的reader
	fr := flate.NewReaderDict(bytes.NewReader(bf), hashDict)
	defer fr.Close()

	b, err := ioutil.ReadAll(fr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
