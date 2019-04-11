package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

// csv读写逗号分隔值（csv）的文件
func main() {

	var data = []string{"test", "Hello", "Go"}

	var buf bytes.Buffer

	// 初始化一个writer
	w := csv.NewWriter(&buf)
	// 写入
	if err := w.Write(data); err != nil {
		log.Fatal(err)
	}
	// 将缓存中的数据写入底层的io.Writer。要检查Flush时是否发生错误的话，应调用Error
	w.Flush()
	// 抓取错误信息
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf.Bytes()))


	// 初始化一个reader
	r := csv.NewReader(&buf)
	for {
		// 从r读取一条记录，返回值record是字符串的切片，每个字符串代表一个字段
		record, err := r.Read()
		if err == io.EOF { // 判断是否结尾
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}

	// 从r中读取所有剩余的记录，每个记录都是字段的切片
	// 成功的调用返回值err为nil而不是EOF,因为ReadAll方法定义为读取直到文件结尾，因此它不会将文件结尾视为应该报告的错误
	// *读取过的记录不会再次被读取
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records)
}
