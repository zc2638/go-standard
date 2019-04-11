package main

import (
	"fmt"
	"unsafe"
)

// unsafe包提供了一些跳过go语言类型安全限制的操作
func main() {

	var hello = Hello{}

	// 返回类型v本身数据所占用的字节数
	// 返回值是“顶层”的数据占有的字节数
	// 例如，若v是一个切片，它会返回该切片描述符的大小，而非该切片底层引用的内存的大小
	s := unsafe.Sizeof(hello)
	fmt.Println(s)

	// 返回类型v所代表的结构体字段在结构体中的偏移量，它必须为结构体类型的字段的形式
	// 换句话说，它返回该结构起始处与该字段起始处之间的字节数
	f := unsafe.Offsetof(hello.b)
	fmt.Println(f)

	// 返回类型v的对齐方式（即类型v在内存中占用的字节数）
	// 若是结构体类型的字段的形式，它会返回字段f在该结构体中的对齐方式
	a := unsafe.Alignof(hello)
	fmt.Println(a)
}

type Hello struct {
	a bool
	b string
	c int
	d []float64
}
