package main

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const FilePath = "testdata/test.gzip"

// gzip包实现了gzip格式压缩文件的读写
func main() {

	// 写gzip数据流
	buf := write()
	// 自动生成并写入文件
	if err := ioutil.WriteFile(FilePath, buf.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
	// 读gzip数据流
	read(FilePath)
}

func write() bytes.Buffer {

	var files = []struct {
		name    string
		comment string
		modTime time.Time
		data    string
	}{
		{"file-1.txt", "file-header-1", time.Now(), "Hello Gophers - 1"},
		{"file-2.txt", "file-header-2", time.Now(), "Hello Gophers - 2"},
	}

	// 声明buffer
	var buf bytes.Buffer

	// 初始化writer
	gw := gzip.NewWriter(&buf)

	// 初始化writer，可设置压缩级别
	gw, err := gzip.NewWriterLevel(&buf, flate.BestCompression)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		gw.Name = file.name       // 设置文件名称
		gw.Comment = file.comment // 设置说明
		gw.ModTime = file.modTime // 设置修改时间
		gw.Extra = []byte("")     // 设置额外内容
		// 写入
		if _, err := gw.Write([]byte(file.data)); err != nil {
			log.Fatal(err)
		}

		// 关闭
		if err := gw.Close(); err != nil {
			log.Fatal(err)
		}

		// 重置buffer内容
		gw.Reset(&buf)
	}

	return buf
}

func read(path string) {

	// 读取文件内容
	bf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewReader(bf)
	// 初始化reader
	gr, err := gzip.NewReader(buf)
	defer gr.Close()
	if err != nil {
		log.Fatal(err)
	}

	for {
		// 设置数据流文件内容分隔，false分开每个单独文件循环读取reader内容。true所有文件内容一次性读取
		gr.Multistream(false)

		b, err := ioutil.ReadAll(gr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n", gr.Name, gr.Comment, gr.ModTime.UTC())
		fmt.Println(string(b))

		err = gr.Reset(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
