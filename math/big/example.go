package main

import "math/big"

// big包实现了大数字的多精度计算
func main() {

	exampleInt()
}

func exampleInt() {

	// 创建一个值为x的*big.Int
	i := big.NewInt(10)

	// 返回x的int64表示，如果不能用int64表示，结果为undefined
	i.Int64()

	// 返回x的uint64表示，如果不能用uint64表示，结果为undefined
	i.Uint64()

	// 返回x的绝对值的大端在前的字节切片表示
	i.Bytes()

	// 返回字符串
	i.String()

	// 返回x的绝对值的字位数，0的字位数为0
	i.BitLen()

	// 提供了对x的数据不检查而快速的访问，返回构成x的绝对值的小端在前的word切片
	// 该切片与x的底层是同一个数组，本函数用于支持在包外实现缺少的低水平功能，否则不应被使用
	i.Bits()

	// 返回第i个字位的值，即返回(x>>i)&1。i必须不小于0
	i.Bit(1)

	// 将z设为x并返回z
	i.SetInt64(11)

	// 将z设为x并返回z
	i.SetUint64(11)

	// 将buf视为一个大端在前的无符号整数，将z设为该值，并返回z
	i.SetBytes([]byte("11"))

	// 将z设为s代表的值（base为基数）
	// 返回z并用一个bool来表明成功与否。如果失败，z的值是不确定的，但返回值为nil
	// 基数必须是0或者2到MaxBase之间的整数。如果基数为0，字符串的前缀决定实际的转换基数："0x"、"0X"表示十六进制；"0b"、"0B"表示二进制；"0"表示八进制；否则为十进制
	i.SetString("11", big.MaxBase)
}