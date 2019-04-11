package main

import "hash/crc64"

// crc64包实现64位循环冗余校验或 CRC-64 校验和
func main() {

	// 声明内容
	var data = []byte("Hello World!")

	var p = []byte("test")

	// 返回一个代表poly指定的多项式的*Table
	// crc64.ISO  ISO 3309定义的ISO多项式，用于HDLC
	// crc64.ECMA ECMA 182定义的ECMA多项式
	t := crc64.MakeTable(crc64.ISO)

	// 返回数据data使用tab代表的多项式计算出的CRC-64校验和
	c := crc64.Checksum(data, t)

	// 返回将切片p的数据采用tab表示的多项式添加到crc之后计算出的新校验和
	crc64.Update(c, t, p)

	// 创建一个使用tab代表的多项式计算CRC-64校验和的hash.Hash64接口
	crc64.New(t)
}