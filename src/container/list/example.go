package main

import (
	"container/list"
	"fmt"
)

// list包实现双向链表
func main() {

	// 初始化链表
	l := list.New()
	l2 := list.New()

	// 向链表后面插入值为v的元素并返回一个新的元素
	e4 := l.PushBack(4)
	// 向链表后面插入另一个链表,l和l2可以相同，但是都不能为nil
	l.PushBackList(l2)
	// 向链表前面面插入值为v的元素并返回一个新的元素
	e1 := l.PushFront(1)
	// 向链表前面插入另一个链表,两个链表可以相同，但是都不能为nil
	l.PushFrontList(l2)
	// 向元素mark前插入值为v的元素并返回一个新的元素，如果mark不是链表的元素，则不改变链表，mark不能为nil
	l.InsertBefore(3, e4)
	// 向元素mark后插入值为v的元素并返回一个新的元素，如果mark不是链表的元素，则不改变链表，mark不能为nil
	l.InsertAfter(2, e1)

	// 如果元素存在链表中，从链表中移除元素，并返回元素值，元素值不能为nil
	l.Remove(e1)
	// 如果元素存在链表中，将元素移动到链表最前面，元素不能为nil
	l.MoveToFront(e4)
	// 如果元素存在链表中，将元素移动到链表最后面，元素不能为nil
	l.MoveToBack(e1)
	// 如果元素e和mark都存在链表中，将e移动到mark前面，两个元素都不能为nil
	l.MoveBefore(e1, e4)
	// 如果元素e和mark都存在链表中，将e移动到mark后面，两个元素都不能为nil
	l.MoveAfter(e1, e4)

	// 返回链表长度
	l.Len()

	// 遍历链表从前面开始打印内容
	fmt.Println("front: ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	// 遍历链表从后面开始打印内容
	fmt.Println("back: ")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}