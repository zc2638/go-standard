package main

import "crypto/subtle"

// 实现了在加密代码中常用的功能，但需要仔细考虑才能正确使用
// 比如[]byte中含有验证用户身份的数据（密文哈希、token等）的时候使用
func main() {

	var x, y int32 = 6, 8
	var a, b byte = 1, 2
	var c, d = 1, 2
	var e, f = []byte("22"), []byte("33")

	// 如果x == y返回1，否则返回0
	subtle.ConstantTimeByteEq(a, b)

	// 如果x == y返回1，否则返回0
	subtle.ConstantTimeEq(x, y)

	// 如果x <= y返回1，否则返回0；如果x或y为负数，或者大于2**31-1，函数行为是未定义的
	subtle.ConstantTimeLessOrEq(c, d)

	// 如果x、y的长度和内容都相同返回1；否则返回0。消耗的时间正比于切片长度而与内容无关
	subtle.ConstantTimeCompare(e, f)

	// 如果v == 1,则将y的内容拷贝到x；如果v == 0，x不作修改；其他情况的行为是未定义并应避免的
	subtle.ConstantTimeCopy(1, e, f)

	// 如果v == 1，返回x；如果v == 0，返回y；其他情况的行为是未定义并应避免的
	subtle.ConstantTimeSelect(0, c, d)
}
