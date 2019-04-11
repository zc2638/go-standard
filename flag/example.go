package main

import (
	"flag"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"
)

// 命令行参数解析
func main() {

	// 返返回解析之后剩下的非flag参数(不包括命令名)
	flag.Args()
	// 返回解析之后剩下的第i个参数，从0开始索引
	flag.Arg(2)
	// 返回解析flag之后剩余参数的个数
	flag.NArg()
	// 返回解析时进行了设置的flag的数量
	flag.NFlag()
	// 从Args中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()
	// 返回是否flag.Parse已经被调用过
	flag.Parsed()

	// 创建一个新的指定名称，指定错误处理策略的FlagSet。
	fg := flag.NewFlagSet("test", flag.PanicOnError)
	// 解析flag出现错误时会被调用
	fg.Usage = func() {
		log.Fatal("异常")
	}
	// fg.Int等同于直接调用flag.Int，其它类型也相同

	// 用指定名称、默认值、使用信息，注册一个int类型flag。返回一个保存了该flag的值的指针
	flag.Int("i", 2, "int flag")
	// 用指定名称、默认值、使用信息，注册一个int类型flag，并将flag的值保存到p指向的变量
	var i int
	flag.IntVar(&i, "i", 2, "int flag")

	var i64 int64
	flag.Int64("i64", 64, "int64 flag")
	flag.Int64Var(&i64, "i64", 64, "int64 flag")

	var f float64
	flag.Float64("f", 6.4, "float64 flag")
	flag.Float64Var(&f, "f", 6.4, "float64 flag")

	var s string
	flag.String("s", "str", "string flag")
	flag.StringVar(&s, "s", "str", "string flag")

	var b bool
	flag.Bool("b", true, "bool flag")
	flag.BoolVar(&b, "b", true, "bool flag")

	var u uint
	flag.Uint("u", 5, "uint flag")
	flag.UintVar(&u, "u", 5, "uint flag")

	var u64 uint64
	flag.Uint64("u64", 5, "uint64 flag")
	flag.Uint64Var(&u64, "u64", 5, "uint64 flag")

	var d time.Duration
	flag.Duration("d", time.Second, "duration flag")
	flag.DurationVar(&d, "d", time.Second, "duration flag")

	// 设置已注册的flag的值
	err := flag.Set("s", "test")
	if err != nil {
		log.Fatal(err)
	}

	// 返回已经已注册flag的Flag结构体指针；如果flag不存在的话，返回nil
	fl := flag.Lookup("s")
	fmt.Println(fl.Usage)          // 使用说明
	fmt.Println(fl.Value.String()) // 要设置的值
	fmt.Println(fl.Name)           // 在命令行中的名称
	fmt.Println(fl.DefValue)       // 默认值

	// 返回引用名称和用法。根据flag的Flag结构，从用法中提取反向引用的名称，如果没有名称为值可能的类型，如果是布尔值则为空字符串
	flag.UnquoteUsage(fl)

	// 用指定名称、使用信息，注册一个flag。该flag的类型的值由第一个参数表示，该参数应实现了Value接口。例如，用户可以创建一个flag，可以用Value接口的Set方法将逗号分隔的字符串转化为字符串切片。
	var ni newVal = 2
	flag.Var(ni, "new", "it is a new flag")

	// 按照字典顺序遍历标签，并且对每个标签调用fn。 这个函数只遍历解析时进行了设置的标签
	flag.Visit(func(f *flag.Flag) {
		fmt.Println(f.Name)
	})
	// 按照字典顺序遍历标签，并且对每个标签调用fn。 这个函数会遍历所有标签，不管解析时有无进行设置
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Println(f.Name)
	})
}

type newVal int
func (v newVal) String() string {
	val := reflect.ValueOf(v)
	return strconv.Itoa(int(val.Int()))
}
func (v newVal) Set(s string) error {
	i, err := strconv.Atoi(s)
	val := reflect.ValueOf(v)
	val.SetInt(int64(i))
	return err
}
