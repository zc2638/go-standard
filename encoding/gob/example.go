package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// gob包管理gob流——在编码器（发送器）和解码器（接受器）之间交换的binary值
// 一般用于传递远端程序调用（RPC）的参数和结果，如net/rpc包就有提供
// 本实现给每一个数据类型都编译生成一个编解码程序，当单个编码器用于传递数据流时，会分期偿还编译的消耗，是效率最高的
func main() {

	// 数据编码解码
	example()

	// 对interface编码解码
	exampleInterface()
}

func example() {

	var data = "Hello World!"

	var data2 = []string{"1", "2", "Hello Go!"}

	// 初始化buffer
	var buf bytes.Buffer

	// 返回一个将编码后数据写入w的*gob.Encoder
	enc := gob.NewEncoder(&buf)

	// 将e内容编码后发送，并且会保证所有的类型信息都先发送
	if err := enc.Encode(data); err != nil {
		log.Fatal("encode error:", err)
	}
	if err := enc.Encode(data2); err != nil {
		log.Fatal("encode error:", err)
	}

	var origin string
	var origin2 []string

	// 返回一个从r读取数据的*gob.Decoder，如果r不满足io.ByteReader接口，则会包装r为bufio.Reader
	dec := gob.NewDecoder(&buf)

	// 从输入流读取下一个值并将该值存入e。如果e是nil，将丢弃该值；否则e必须是可接收该值的类型的指针
	// 如果输入结束，方法会返回io.EOF并且不修改e（指向的值）
	if err := dec.Decode(&origin); err != nil {
		log.Fatal("decode error 1:", err)
	}
	if err := dec.Decode(&origin2); err != nil {
		log.Fatal("decode error 2:", err)
	}
	fmt.Println(origin, origin2)
}

func exampleInterface() {

	// 声明buffer
	var buf bytes.Buffer

	// 记录value下层具体值的类型和其名称
	// 该名称将用来识别发送或接受接口类型值时下层的具体类型
	// 本函数只应在初始化时调用，如果类型和名字的映射不是一一对应的，会panic
	gob.Register(Hello{})

	// 初始化一个encoder
	enc := gob.NewEncoder(&buf)
	// 写入编码数据
	interfaceEncode(enc, Hello{"World"})

	// 初始化一个decoder
	dec := gob.NewDecoder(&buf)
	// 读取数据并使用interface的方法
	interfaceDecode(dec).Say()
}

type Hello struct {
	Name string
}

func (h Hello) Say() { fmt.Println("Hello " + h.Name) }

type Pythagoras interface {
	Say()
}

func interfaceEncode(enc *gob.Encoder, p Pythagoras) {

	err := enc.Encode(&p)
	if err != nil {
		log.Fatal("encode:", err)
	}
}

func interfaceDecode(dec *gob.Decoder) (p Pythagoras) {

	err := dec.Decode(&p)
	if err != nil {
		log.Fatal("decode:", err)
	}
	return
}
