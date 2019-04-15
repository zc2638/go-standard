package main

import (
	"fmt"
	"hash/fnv"
	"log"
)

// fnv包实现了FNV-1和FNV-1a（非加密hash函数）
func main() {

	var data = []byte("Hello World!")

	// 返回一个新的32位FNV-1的hash.Hash32接口
	h := fnv.New32()
	// 写入
	if _, err := h.Write(data); err != nil {
		log.Fatal(err)
	}
	fmt.Println(h.Sum32())

	// 返回一个新的32位FNV-1a的hash.Hash32接口
	fnv.New32a()

	// 返回一个新的64位FNV-1的hash.Hash64接口
	fnv.New64()

	// 返回一个新的64位FNV-1a的hash.Hash64接口
	fnv.New64a()

	//返回一个新的128位FNV-1的hash.Hash64接口
	fnv.New128()

	// 返回一个新的128位FNV-1a的hash.Hash64接口
	fnv.New128a()
}