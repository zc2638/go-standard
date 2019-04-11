package main

import (
	"fmt"
	"sort"
)

// sort包提供了排序切片和用户自定义数据集的函数
func main() {

	// 对任意切片排序，函数实现版
	example()
	// 对任意切片排序，函数实现版，推荐使用
	exampleStable()
	// 对任意 函数实现版 切片类型 反向排序
	exampleReverse()
	// []string排序
	exampleString()
	// []int排序
	exampleInt()
	// []float64排序
	exampleFloat()
	// 对任意切片排序
	exampleSlice()
	// 对任意切片排序，推荐使用
	exampleSliceStable()
}

type People struct {
	Name string
	Age  int
}

type Peoples []People

func (p Peoples) Len() int           { return len(p) }
func (p Peoples) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Peoples) Less(i, j int) bool { return p[i].Age < p[j].Age } // 按age递增排序

func example() {

	// 声明一个简单的结构体
	peoples := Peoples{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	// 任意切片排序，但是切片必须包含Len(),Swap(),Less()方法
	sort.Sort(peoples)
	fmt.Println(peoples)

	// 检查是否已经被排序
	fmt.Println(sort.IsSorted(peoples))
}

func exampleStable() {

	// 声明一个简单的结构体
	peoples := Peoples{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	// 任意切片排序，但是切片必须包含Len(),Swap(),Less()方法
	sort.Stable(peoples)
	fmt.Println(peoples)

	// 检查是否已经被排序
	fmt.Println(sort.IsSorted(peoples))
}

func exampleReverse() {

	// 声明一个简单的结构体
	peoples := Peoples{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	// 包装一个Interface接口并返回一个新的Interface接口，对该接口排序可生成递减序列
	p := sort.Reverse(peoples)
	// 排序
	sort.Sort(p)
	fmt.Println(peoples)
}

func exampleString() {

	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}

	// 排序为递增顺序
	sort.Strings(s)
	fmt.Println(s)

	s = []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	// 排序为递增顺序
	// 将[]string指定为sort.StringSlice类型并排序，类型内部实现了排序方法
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

	// 检查是否已排序为递增顺序
	fmt.Println(sort.StringsAreSorted(s))

	// 返回x在a中应该存在的位置，无论a中是否存在a中
	// 在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)
	index := sort.SearchStrings(s, "Hello")
	fmt.Println(index)
}

func exampleInt() {

	s := []int{5, 2, 6, 3, 1, 4}

	// 排序为递增顺序
	sort.Ints(s)
	fmt.Println(s)


	s = []int{5, 2, 6, 3, 1, 4}
	// 排序为递增顺序
	// 将[]int指定为sort.IntSlice类型并排序，类型内部实现了排序方法
	sort.Sort(sort.IntSlice(s))
	fmt.Println(s)

	// 检查是否已排序为递增顺序
	fmt.Println(sort.IntsAreSorted(s))

	// 返回x在a中应该存在的位置，无论a中是否存在a中
	// 在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)
	index := sort.SearchInts(s, 5)
	fmt.Println(index)
}

func exampleFloat() {

	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6}

	// 排序为递增顺序
	sort.Float64s(s)
	fmt.Println(s)

	s = []float64{5.2, -1.3, 0.7, -3.8, 2.6}
	// 将[]float64指定为sort.Float64Slice类型并排序，类型内部实现了排序方法
	sort.Sort(sort.Float64Slice(s))
	fmt.Println(s)

	// 检查是否已排序为递增序列
	fmt.Println(sort.Float64sAreSorted(s))

	// 返回x在a中应该存在的位置，无论a中是否存在a中
	// 在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)
	index := sort.SearchFloat64s(s, 2)
	fmt.Println(index)
}

func exampleSlice() {

	// 声明一个简单的结构体
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	// 对切片自定义进行排序
	// 该接口不能保证是稳定的。对于稳定类型，使用SliceStable
	// 如果提供的接口不是切片，则函数将终止
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("By age:", people)
}

func exampleSliceStable() {

	// 声明一个简单的结构体
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Elizabeth", 75},
		{"Alice", 75},
		{"Bob", 75},
		{"Alice", 75},
		{"Bob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}

	// 对提供的切片进行排序，同时保持相等元素的原始顺序
	// 如果提供的接口不是切片，则函数将终止
	sort.SliceStable(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	sort.SliceStable(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age,name:", people)
}