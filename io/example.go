package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// io包提供了对I/O原语的基本接口
// 本包的基本任务是包装这些原语已有的实现（如os包里的原语），使之成为共享的公共接口，这些公共接口抽象出了泛用的函数并附加了一些相关的原语的操作
// 因为这些接口和原语是对底层实现完全不同的低水平操作的包装，除非得到其它方面的通知，客户端不应假设它们是并发执行安全的
func main() {

	// 读取reader内容向指定大小buffer填充
	exampleReadFull()
	// 向writer中写入字符串
	exampleWriteString()
	// writer拷贝reader内容
	exampleCopy()
	// writer拷贝reader指定字节数的内容
	exampleCopayN()
	// writer拷贝reader内容，提供指定大小的缓冲区
	exampleCopyBuffer()
	// 从reader读取至少n字节的内容填充指定大小的buffer
	exampleReadAtLeast()
	// 从reader中读取n个字节并返回一个新的reader
	exampleLimitReader()
	// 将几个reader内容串联起来并返回一个新的reader
	exampleMultiReader()
	// 返回一个将其从r读取的数据写入w的Reader接口。所有通过该接口对r的读取都会执行对应的对w的写入
	exampleTeeReader()
	// 返回一个从reader第off个字节开始往后n个的内容的新reader
	exampleSectionReader()
	// 创建一个可以同时写入多个writer的writer接口
	exampleMultiWriter()
	// 管道实例
	examplePipe()
}

func exampleReadFull() {

	// 根据字符串创建一个reader
	r := strings.NewReader("some io.Reader stream to be read\n")

	// 声明4个字节的buffer
	buf := make([]byte, 4)

	// 从reader精确地读取len(buf)字节数据填充进buffer
	// 函数返回写入的字节数和错误（如果没有读取足够的字节）
	// 只有没有读取到字节时才可能返回EOF；如果读取了有但不够的字节时遇到了EOF，函数会返回ErrUnexpectedEOF
	// 只有返回值err为nil时，返回值n才会等于len(buf)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)
}

func exampleWriteString() {

	var buf bytes.Buffer
	// 将字符串s的内容写入w中。如果w已经实现了WriteString方法，函数会直接调用该方法
	if _, err := io.WriteString(&buf, "Hello World"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func exampleCopy() {

	// 声明buffer
	var buf bytes.Buffer
	// 根据字符串创建reader
	r := strings.NewReader("some io.Reader stream to be read")

	// 将src的数据拷贝到dst，直到在src上到达EOF或发生错误
	// 返回拷贝的字节数和遇到的第一个错误
	if _, err := io.Copy(&buf, r); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func exampleCopayN() {

	// 声明buffer
	var buf bytes.Buffer
	// 根据字符串创建reader
	r := strings.NewReader("some io.Reader stream to be read")

	// 将src的数据拷贝到dst，直到在src上到达EOF或发生错误
	// 返回拷贝的字节数和遇到的第一个错误
	if _, err := io.CopyN(&buf, r, 10); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func exampleCopyBuffer() {

	var buf bytes.Buffer

	r1 := strings.NewReader("copy reader")
	b := make([]byte, 8)

	// 与 Copy 相同，只是它通过提供的缓冲区（如果需要的话）进行分级，而不是分配临时的缓冲区
	// 如果buf为nil，则分配一个；否则如果它的长度为零，CopyBuffer将会panic
	if _, err := io.CopyBuffer(&buf, r1, b); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func exampleReadAtLeast() {

	r := strings.NewReader("some io.Reader stream to be read\n")
	buf := make([]byte, 10)

	// 从r至少读取min字节数据填充进buf
	// 返回写入的字节数和错误（如果没有读取足够的字节）
	// 只有没有读取到字节时才可能返回EOF；如果读取了有但不够的字节时遇到了EOF，会返回ErrUnexpectedEOF
	// ;如果min比buf的长度还大，函数会返回ErrShortBuffer。只有返回值err为nil时，返回值n才会不小于min
	if _, err := io.ReadAtLeast(r, buf, 6); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)
}

func exampleLimitReader() {

	r := strings.NewReader("some io.Reader stream to be read\n")

	// 返回一个从r中读取了n个字节以EOF停止的Reader
	// 返回值接口的底层为*io.LimitedReader类型
	lr := io.LimitReader(r, 4)
	// 拷贝内容到输出
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}


func exampleMultiReader() {

	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")

	// 返回一个将提供的Reader在逻辑上串联起来的Reader接口
	// 他们依次被读取，当所有的输入流都读取完毕，Read才会返回EOF
	// 如果readers中任一个返回了非nil非EOF的错误，Read方法会返回该错误
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func exampleTeeReader() {

	r := strings.NewReader("some io.Reader stream to be read\n")
	var buf bytes.Buffer

	// 返回一个将其从r读取的数据写入w的Reader接口
	// 所有通过该接口对r的读取都会执行对应的对w的写入
	// 没有内部的缓冲：写入必须在读取完成前完成，写入时遇到的任何错误都会作为读取错误返回
	tee := io.TeeReader(r, &buf)

	if _, err := io.Copy(os.Stdout, tee); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func exampleSectionReader() {

	r := strings.NewReader("some io.Reader stream to be read\n")

	// 返回一个从r中的偏移量off处为起始，读取n个字节后以EOF停止的SectionReader
	s := io.NewSectionReader(r, 5, 17)

	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
	fmt.Println()


	// 声明一个长度为6的[]byte
	be := make([]byte, 6)

	// 从s的第off个字节开始读取内容写入p
	if _, err := s.ReadAt(be, 10); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(be))


	// 声明一个长度为6的[]byte
	//bs := make([]byte, 6)

	// 将s内容填充为从s第10个开始往后的内容
	// io.SeekStart 从前面开始往后offset个
	// io.SeekEnd   从结尾开始往前-offset个
	if _, err := s.Seek(10, io.SeekStart); err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

func exampleMultiWriter() {

	r := strings.NewReader("some io.Reader stream to be read\n")

	// 声明两个buffer
	var buf1, buf2 bytes.Buffer

	// 创建一个Writer接口，会将提供给其的数据写入所有创建时提供的Writer接口
	w := io.MultiWriter(&buf1, &buf2)

	// 拷贝
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
	fmt.Print(buf1.String())
	fmt.Print(buf2.String())
}

func examplePipe() {

	// 创建一个同步的内存中的管道
	// 它可以用于连接期望io.Reader的代码和期望io.Writer的代码
	// 一端的读取对应另一端的写入，直接在两端拷贝数据，没有内部缓冲
	// 可以安全的并行调用Read和Write或者Read/Write与Close方法
	// Close方法会在最后一次阻塞中的I/O操作结束后完成
	// 并行调用Read或并行调用Write也是安全的：每一个独立的调用会依次进行
	r, w := io.Pipe()

	go func() {
		w.Write([]byte("some text to be read\n"))
		w.Close()
	}()

	// 初始化一个buffer
	buf := new(bytes.Buffer)
	// 从reader读取内容
	buf.ReadFrom(r)
	fmt.Print(buf.String())
}
