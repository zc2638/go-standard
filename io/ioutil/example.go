package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const Dir = "testdata/testDir"
const File = "testdata/test"

// ioutil包实现了一些 I/O 实用程序功能
func main() {

	// 声明内容
	var data = []byte("Go is a general-purpose language designed with systems programming in mind.")

	// 创建一个reader
	var r = bytes.NewReader(data)

	// 从r读取数据直到EOF或遇到error，返回读取的数据和遇到的错误
	// 成功的调用返回的err为nil而非EOF
	// 因为本函数定义为读取r直到EOF，它不会将读取返回的EOF视为应报告的错误
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	// 向filename指定的文件中写入数据。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件
	_ = ioutil.WriteFile(File, data, os.ModePerm)

	// 从filename指定的文件中读取数据并返回文件的内容
	// 成功的调用返回的err为nil而非EOF
	// 因为本函数定义为读取整个文件，它不会将读取返回的EOF视为应报告的错误
	bf, err := ioutil.ReadFile(File)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bf))

	// 返回dirname指定的目录的目录信息的有序列表
	files, err := ioutil.ReadDir(Dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {

		// 文件名称(不含扩展名)
		fmt.Println(file.Name())
		// 修改时间
		fmt.Println(file.ModTime())
		// 文件模式
		fmt.Println(file.Mode())
		// 判断是否是一个目录，等价于Mode().IsDir()
		fmt.Println(file.IsDir())
		// 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
		fmt.Println(file.Size())
		// 基础数据源（可以返回nil）
		fmt.Println(file.Sys())
	}

	// 用一个无操作的Close方法包装r返回一个ReadCloser接口
	rc := ioutil.NopCloser(r)
	if err := rc.Close(); err != nil {
		log.Fatal(err)
	}

	// 创建temp文件
	exampleTempFile()
	// 创建temp目录
	exampleTempDir()
}

func exampleTempDir() {

	content := []byte("temporary file's content")
	// 在dir目录里创建一个新的、使用prfix作为前缀的临时文件夹，并返回文件夹的路径
	// 如果dir是空字符串，TempDir使用默认用于临时文件的目录（参见os.TempDir函数）
	// 不同程序同时调用该函数会创建不同的临时目录，调用本函数的程序有责任在不需要临时文件夹时摧毁它
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}
	// 移除文件夹
	defer os.RemoveAll(dir)

	// 将新的文件夹路径放入一个单一路径里
	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
}

func exampleTempFile() {

	data := []byte("temporary file's content")

	// 在dir目录下创建一个新的、使用prefix为前缀的临时文件，以读写模式打开该文件并返回os.File指针
	// 如果dir是空字符串，TempFile使用默认用于临时文件的目录（参见os.TempDir函数）
	// 不同程序同时调用该函数会创建不同的临时文件，调用本函数的程序有责任在不需要临时文件时摧毁它
	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	// 移除文件
	defer os.Remove(tmpFile.Name())
	// 写入
	if _, err := tmpFile.Write(data); err != nil {
		log.Fatal(err)
	}
	// 关闭
	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}
}
