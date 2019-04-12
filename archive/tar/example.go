package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// 文件类型
const (
	TypeReg           = tar.TypeReg           // 普通文件
	TypeLink          = tar.TypeLink          // 硬链接
	TypeSymlink       = tar.TypeSymlink       // 符号链接
	TypeChar          = tar.TypeChar          // 字符节点
	TypeBlock         = tar.TypeBlock         // 块节点
	TypeDir           = tar.TypeDir           // 目录
	TypeFifo          = tar.TypeFifo          // 先进先出队列节点
	TypeCont          = tar.TypeCont          // 保留位
	TypeXHeader       = tar.TypeXHeader       // 扩展头
	TypeXGlobalHeader = tar.TypeXGlobalHeader // 全局扩展头
	TypeGNULongName   = tar.TypeGNULongName   // 下一个文件记录有个长名字
	TypeGNULongLink   = tar.TypeGNULongLink   // 下一个文件记录指向一个具有长名字的文件
	TypeGNUSparse     = tar.TypeGNUSparse     // 稀疏文件
)

const FilePath = "testdata/test.tar"

// 实现了tar格式压缩文件的存取
func main() {

	// 写tar文件数据流
	buf := write()

	// 自动生成并写入文件
	if err := ioutil.WriteFile(FilePath, buf.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// 读tar文件数据流
	read(FilePath)
}

func write() bytes.Buffer {

	// 声明buffer
	var buf bytes.Buffer

	// 初始化writer
	var tw = tar.NewWriter(&buf)

	// 实例化一个结构体切片
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}

	for _, file := range files {

		// 根据结构体的内容实例化一个header
		hdr := &tar.Header{
			Name:       file.Name,             // 记录头域的文件名
			Mode:       0600,                  // 权限和模式位
			Uid:        0,                     // 所有者的用户ID
			Gid:        0,                     // 所有者的组ID
			Size:       int64(len(file.Body)), // 字节数（长度）
			ModTime:    time.Now(),            // 修改时间
			Typeflag:   TypeReg,               // 文件类型
			Linkname:   "",                    // 链接的目标名
			Uname:      "",                    // 所有者的用户名
			Gname:      "",                    // 所有者的组名
			Devmajor:   0,                     // 字符设备或块设备的major number
			Devminor:   0,                     // 字符设备或块设备的minor number
			AccessTime: time.Now(),            // 访问时间
			ChangeTime: time.Now(),            // 状态改变时间
		}
		// writer写入header
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		// writer写入内容
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}

	// 关闭writer
	if err := tw.Close(); err != nil {
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

	// 实例化buffer
	var readBuf = bytes.NewBuffer(bf)

	// 初始化一个reader去读取tar内容
	tr := tar.NewReader(readBuf)

	// 循环读取多个文件内容
	for {
		// 获取单个文件的header信息
		hdr, err := tr.Next()

		// 所有文件读取完毕
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)

		// 读取数据流信息
		b, err := ioutil.ReadAll(tr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
}