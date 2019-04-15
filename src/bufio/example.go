package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const FilePath = "testdata/bufio"

// bufio包实现了有缓冲的I/O
func main() {
	// 缓冲区写入
	buf := write()
	fmt.Printf("%s\n", buf.Bytes())

	// 写入文件读取测试使用
	if err := ioutil.WriteFile(FilePath, buf.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// 缓冲区读取
	read(FilePath)

	// 扫描器demo
	scannerDemo()
	// 扫描器自定义demo
	scannerCustomDemo()
	// 扫描器切割最后个为空时避免报错的demo
	scannerSplitWithCommaDemo()


}

func write() bytes.Buffer {

	// 声明buffer
	var buf bytes.Buffer

	// 初始化writer
	var w = bufio.NewWriter(&buf)

	// 创建一个自定义长度的缓冲区，必须大于默认值16，小于默认值则以默认值大小生成。
	//var w = bufio.NewWriterSize(&buf, 32)

	// 写入字符串类型
	w.WriteString("Hello, ")
	// 写入单个rune类型
	w.WriteRune('W')
	// 写入单个byte类型
	w.WriteByte('o')
	// 写入byte类型
	w.Write([]byte("rld!"))
	// 重置当前缓冲区
	w.Flush()

	return buf
}

func read(path string) {

	var buf bytes.Buffer

	// 读取文件内容
	bf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// 初始化reader
	var r = bufio.NewReader(bytes.NewReader(bf))

	// 读取内容直到出现 W 停止并输出(丢出缓冲区)，其余内容仍在缓冲区
	// r.ReadBytes('W')效果如下
	if s, err := r.ReadString('W'); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(s)
	}

	// 返回可以从当前缓冲区读取的字节数。
	fmt.Println(r.Buffered())

	// 将缓冲区内容写入buffer
	r.WriteTo(&buf)
	fmt.Printf("%s\n", buf.Bytes())
}

func scannerDemo() {

	// 声明字符串
	input := "foo bar   baz"
	// 初始化扫描器模块
	scanner := bufio.NewScanner(strings.NewReader(input))

	// split调用一个split函数，函数内容可以根据格式自己定义。
	// scanwords是一个自带的split函数，用于返回以空格分隔的字符串，删除了周围的空格。它绝不返回空字符串。
	scanner.Split(bufio.ScanWords)

	// 循环读取
	for scanner.Scan() {
		//fmt.Printf("%s\n", scanner.Bytes())
		fmt.Println(scanner.Text())
	}
	// 获取异常
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func scannerCustomDemo() {
	// An artificial input source.
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 自定义split函数
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// 自带的split函数。
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			// 字符串转int
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	scanner.Split(split)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Invalid input: ", err)
	}
}

func scannerSplitWithCommaDemo() {

	const input = "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		// 最后一个可能是空字符串，error传入ErrFinalToken告诉bufio是最后个，使其不报错。
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Invalid input:", err)
	}
}