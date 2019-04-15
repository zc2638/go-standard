package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const FilePath = "testdata/test.zip"

// 实现了zip档案文件的读写
func main() {
	// 写zip文件数据流
	buf := write()

	// 自动生成并写入文件
	if err := ioutil.WriteFile(FilePath, buf.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// 读取zip文件数据流
	read(FilePath)
}

func write() bytes.Buffer {

	// 声明buffer
	var buf bytes.Buffer

	// 初始化writer
	w := zip.NewWriter(&buf)

	// 设置压缩级别，不指定则默认
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	// 实例化一个结构体切片
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {

		// 根据文件名称，writer创建文件
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		// 创建的文件写入内容
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 关闭writer.使用defer更舒适
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	return buf
}

func read(path string) {

	// 根据文件路径，获取zip文件内容
	r, err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(err)
	}
	// 方法最后调用关闭
	defer r.Close()

	// 循环读取多个文件内容
	for _, f := range r.File {

		fmt.Printf("Contents of %s:\n", f.Name)
		// 打开文件
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		// 读取文件内容
		b, err := ioutil.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
		// 关闭文件
		rc.Close()
	}
}