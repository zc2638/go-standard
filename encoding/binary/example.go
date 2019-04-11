package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
)

// binary包实现了简单的数字与byte的转换以及变长值的编解
// 数字翻译为定长值来读写，一个定长值，要么是固定长度的数字类型（int8, uint8, int16, float32, complex64, ...）或者只包含定长值的结构体或者数组
// 本包相对于效率更注重简单。如果需要高效的序列化，特别是数据结构较复杂的，请参见更高级的解决方法，例如encoding/gob包，或者采用协议缓存
// 简述：可用于 将uint，float，complex等类型与[]byte类型互相转换
func main() {

	// 编码/解码
	example()
	// 多个数字一起编码/解码
	multiExample()

	ByteOrder()
	uvarint()
}

func example() {

	// 指定binary写入时的字节序
	// binary.BigEndian 大端字节序的实现
	// binary.LittleEndian 小端字节序的实现
	var order = binary.LittleEndian
	// 声明内容
	var data = math.Pi
	// 初始化一个buffer
	buf := new(bytes.Buffer)

	// 将data的binary编码格式写入w，data必须是定长值、定长值的切片、定长值的指针
	// order指定写入数据的字节序，写入结构体时，名字中有'_'的字段会置为0
	if err := binary.Write(buf, order, data); err != nil {
		log.Fatal("binary.Write failed:", err)
	}
	fmt.Printf("%x\n", buf.Bytes())


	var origin float64
	// 从r中读取binary编码的数据并赋给data，data必须是一个指向定长值的指针或者定长值的切片
	// 从r读取的字节使用order指定的字节序解码并写入data的字段里当写入结构体是，名字中有'_'的字段会被跳过，这些字段可用于填充（内存空间）
	if err := binary.Read(buf, order, &origin); err != nil {
		log.Fatal("binary.Read failed:", err)
	}
	fmt.Println(origin)
}

func multiExample() {

	// 初始化buffer
	buf := new(bytes.Buffer)
	// 声明interface{}切片
	var data = []interface{}{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	// 循环切片按顺序编码写入buf
	for _, v := range data {
		if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
			log.Fatal("binary.Write failed:", err)
		}
	}
	fmt.Printf("%x\n", buf.Bytes())

	var origin = struct {
		A uint16
		B int8
		C uint8
	}{}

	// 解码
	if err := binary.Read(buf, binary.LittleEndian, &origin); err != nil {
		log.Fatal("binary.Read failed:", err)
	}
	fmt.Println(origin)
}

func ByteOrder() {

	// 声明一个长度为4的[]byte
	b := make([]byte, 4)
	// 将16比特的无符号整数 转为 字节序列
	binary.LittleEndian.PutUint16(b[:2], '1')
	binary.LittleEndian.PutUint16(b[2:], 0x07d0)
	fmt.Printf("% x\n", b)

	// 将字节序列 转为 16比特的无符号整数
	x1 := binary.LittleEndian.Uint16(b[:2])
	x2 := binary.LittleEndian.Uint16(b[2:])
	fmt.Printf("%#04x %#04x\n", x1, x2)
}

func uvarint() {

	var n uint64 = 256

	// 根据 变长编码N位整数的最大字节数 初始化一个[]byte
	buf := make([]byte, binary.MaxVarintLen64)
	// 将一个uint64数字编码写入buf并返回写入的长度，如果buf太小，则会panic
	binary.PutUvarint(buf, n)

	// 从reader读取一个编码后的无符号整数，并返回该整数
	i , err := binary.ReadUvarint(bytes.NewReader(buf))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)

	// binary.PutVarint binary.ReadVarint 同上
}