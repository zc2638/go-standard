package main

import (
	"fmt"
	"hash/crc32"
	"log"
)

// crc32包实现了32位循环冗余校验（CRC-32）的校验和算法
func main() {

	// 声明内容
	var data = []byte("Hello World!")

	var p = []byte("test")

	// 创建一个使用IEEE多项式计算CRC-32校验和的hash.Hash32接口
	h := crc32.NewIEEE()
	// 写入
	if _, err := h.Write(data); err != nil {
		log.Fatal(err)
	}
	// 返回uint32值
	s := h.Sum32()
	fmt.Println(s)

	// 返回一个代表poly指定的多项式的Table
	// crc32.IEEE        最常用的CRC-32多项式；用于以太网、v.42、fddi、gzip、zip、png、mpeg-2……
	// crc32.Castagnoli  卡斯塔尼奥利多项式，用在iSCSI；有比IEEE更好的错误探测特性
	// crc32.Koopman     库普曼多项式；错误探测特性也比IEEE好
	t := crc32.MakeTable(crc32.IEEE)

	// 返回数据data使用tab代表的多项式计算出的CRC-32校验和
	crc32.Checksum(data, t)

	// 返回数据data使用IEEE多项式计算出的CRC-32校验和
	crc32.ChecksumIEEE(data)

	// 返回将切片p的数据采用tab表示的多项式添加到crc之后计算出的新校验和
	crc32.Update(s, t, p)

	// 创建一个使用tab代表的多项式计算CRC-32校验和的hash.Hash32接口
	crc32.New(t)
}