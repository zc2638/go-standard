package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

// rand包实现了伪随机数生成器
func main() {

	// 全局随机数
	example()

	// 自定随机数
	exampleRand()

	// 切片值随机排序
	exampleShuffle()
}

func example() {

	// 使用给定的seed来初始化生成器到一个确定的状态
	// 如未调用，默认资源的行为就好像调用了Seed(1)
	// 如果需要实现随机数，seed值建议为 时间戳
	// 如果seed给定固定值，结果永远相同
	rand.Seed(time.Now().UnixNano())

	// 返回一个非负的伪随机int值
	fmt.Println(rand.Int())

	// 返回一个取值范围在[0,n]的伪随机int值，如果n<=0会panic
	fmt.Println(rand.Intn(100))

	// 返回一个int32类型的非负的31位伪随机数
	fmt.Println(rand.Int31())

	// 返回一个取值范围在[0,n]的伪随机int32值，如果n<=0会panic
	fmt.Println(rand.Int31n(100))

	// 返回一个int64类型的非负的63位伪随机数
	fmt.Println(rand.Int63())

	// 返回一个取值范围在[0, n]的伪随机int64值，如果n<=0会panic
	fmt.Println(rand.Int63n(100))

	// 返回一个uint32类型的非负的32位伪随机数
	fmt.Println(rand.Uint32())

	// 返回一个uint64类型的非负的64位伪随机数
	fmt.Println(rand.Uint64())

	// 返回一个取值范围在[0, 1]的伪随机float32值
	fmt.Println(rand.Float32())

	// 返回一个取值范围在[0, 1]的伪随机float64值
	fmt.Println(rand.Float64())

	// 返回一个服从标准正态分布（标准差=1，期望=0）、取值范围在[-math.MaxFloat64, +math.MaxFloat64]的float64值
	fmt.Println(rand.NormFloat64())

	// 返回一个服从标准指数分布（率参数=1，率参数是期望的倒数）、取值范围在(0, +math.MaxFloat64]的float64值
	fmt.Println(rand.ExpFloat64())

	// 返回一个有n个元素的，[0,n)范围内整数的伪随机排列的切片
	fmt.Println(rand.Perm(15))
}

func exampleRand() {

	// 初始化一个Source，代表一个生成均匀分布在范围[0, 1<<63)的int64值的（伪随机的）资源
	// 需要随机，建议使用时间戳，详情查看example()
	source := rand.NewSource(99)

	// 初始化*rand.Rand
	r := rand.New(source)

	// 初始化一个过滤器
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()

	show := func(name string, v1, v2, v3 interface{}) {
		fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", name, v1, v2, v3)
	}

	show("Float32", r.Float32(), r.Float32(), r.Float32())
	show("Float64", r.Float64(), r.Float64(), r.Float64())
	show("ExpFloat64", r.ExpFloat64(), r.ExpFloat64(), r.ExpFloat64())
	show("NormFloat64", r.NormFloat64(), r.NormFloat64(), r.NormFloat64())
	show("Int31", r.Int31(), r.Int31(), r.Int31())
	show("Int63", r.Int63(), r.Int63(), r.Int63())
	show("Uint32", r.Uint32(), r.Uint32(), r.Uint32())
	show("Intn(10)", r.Intn(10), r.Intn(10), r.Intn(10))
	show("Int31n(10)", r.Int31n(10), r.Int31n(10), r.Int31n(10))
	show("Int63n(10)", r.Int63n(10), r.Int63n(10), r.Int63n(10))
	show("Perm", r.Perm(5), r.Perm(5), r.Perm(5))
}

func exampleShuffle() {

	numbers := []byte("12345")
	letters := []byte("ABCDE")

	// 随机交换数切片的位置, n应为切片的长度，n < 0 会panic
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
		letters[i], letters[j] = letters[j], letters[i]
	})
	for i := range numbers {
		fmt.Printf("%c: %c\n", letters[i], numbers[i])
	}
}