package main

import (
	"container/ring"
	"fmt"
)

// ring包实现环形链表
func main() {

	// 初始化n个元素的环形链表
	r := ring.New(5)
	s := ring.New(5)
	// 返回链表长度
	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = i  // 给元素赋值
		r = r.Next() // 获取下一个元素

		s.Value = i
		s = s.Next()
	}

	for j := 0; j < n; j++ {
		r = r.Prev()         // 获取上一个元素
		fmt.Println(r.Value) //
	}

	// 循环访问环形链表所有元素
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

	// 将前面n个元素移到后面，例：0,1,2,3,4,5 => 3,4,5,0,1,2
	r.Move(3)

	// 链表r与链表s是不同链表，则在r链表的后面链接s链表，否则删除相同部分
	r.Link(s)
	// 从下一个元素开始，移除链表连续n个元素
	r.Unlink(3)
}
