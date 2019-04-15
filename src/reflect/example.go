package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

// reflect包实现了运行时反射，允许程序操作任意类型的对象
// 典型用法是用静态类型interface{}保存一个值，通过调用TypeOf获取其动态类型信息，该函数返回一个Type类型值
// 调用ValueOf函数返回一个Value类型值，该值代表运行时的数据
// Zero接受一个Type类型参数并返回一个代表该类型零值的Value类型值
func main() {

	example()
	exampleType()
	exampleValue()
	exampleMake()
}

type Hello struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Arr  []string `json:""`
}

func (h *Hello) Say() {
	fmt.Println("Hello World!")
}

func (h Hello) Says() {
	fmt.Println("Hello World!")
}

func exampleType() {

	hello := Hello{}

	// 返回接口中保存的值的类型，TypeOf(nil)会返回nil
	t := reflect.TypeOf(hello)

	// 返回该类型的元素类型，如果该类型的Kind不是Array、Chan、Map、Ptr或Slice，会panic
	reflect.TypeOf(&hello).Elem()

	// 返回该接口的具体类型
	t.Kind()
	fmt.Println("reflect.Struct:", t.Kind() == reflect.Struct)

	// 返回该类型在自身包内的类型名，如果是未命名类型会返回""
	fmt.Println("Name:", t.Name())

	// 返回类型的包路径，即明确指定包的import路径，如"encoding/base64"
	// 如果类型为内建类型(string, error)或未命名类型(*T, struct{}, []int)，会返回""
	fmt.Println("PkgPath:", t.PkgPath())

	// 返回类型的字符串表示。该字符串可能会使用短包名（如用base64代替"encoding/base64"）
	// 也不保证每个类型的字符串表示不同。如果要比较两个类型是否相等，请直接用Type类型比较
	fmt.Println("String:", t.String())

	// 返回要保存一个该类型的值需要多少字节；类似unsafe.Sizeof
	fmt.Println("Size:", t.Size())

	// 返回当从内存中申请一个该类型值时，会对齐的字节数
	fmt.Println("Align:", t.Align())

	// 返回当该类型作为结构体的字段时，会对齐的字节数
	fmt.Println("FieldAlign:", t.FieldAlign())

	// 返回该类型的方法集中方法的数目
	// 匿名字段的方法会被计算；主体类型的方法会屏蔽匿名字段的同名方法；
	// 匿名字段导致的歧义方法会滤除
	fmt.Println("NumMethod:", t.NumMethod())

	// 返回该类型方法集中的第i个方法，i不在[0, NumMethod())范围内时，将导致panic
	// 对非接口类型T或*T，返回值的Type字段和Func字段描述方法的未绑定函数状态
	// 对接口类型，返回值的Type字段描述方法的签名，Func字段为nil
	fmt.Println("Method:", t.Method(0))

	// 根据方法名返回该类型方法集中的方法，使用一个布尔值说明是否发现该方法
	// 对非接口类型T或*T，返回值的Type字段和Func字段描述方法的未绑定函数状态
	// 对接口类型，返回值的Type字段描述方法的签名，Func字段为nil
	fmt.Println(t.MethodByName("Says"))

	// 返回struct类型的字段数（匿名字段算作一个字段），如非结构体类型将panic
	fmt.Println("NumField:", t.NumField())

	// 返回struct类型的第i个字段的类型，如非结构体或者i不在[0, NumField())内将会panic
	fmt.Println("Field:", t.Field(0))

	// 获取字段下的tag标签
	// 使用Tag.Get时不存在的tag会返回空字符串
	fmt.Println("Field Tag(json):", t.Field(0).Tag.Get("jsons"))

	// 查找指定名称的tag，返回 tag内容 和 布尔值（用于判断是否存在tag）
	// 当布尔值为false时，这个tag不存在
	// 当布尔值为true时，tag内容为空字符串时，说明tag存在但是设为空值
	fmt.Println(t.Field(2).Tag.Lookup("json"))

	// 返回索引序列指定的嵌套字段的类型，
	// 等价于用索引中每个值链式调用本方法，如非结构体将会panic
	fmt.Println("FieldByIndex:", t.FieldByIndex([]int{1}))

	// 返回该类型名为name的字段（会查找匿名字段及其子字段），
	// 布尔值说明是否找到，如非结构体将panic
	fmt.Println(t.FieldByName("Name"))

	// 返回该类型第一个字段名满足函数match的字段，布尔值说明是否找到，如非结构体将会panic
	fmt.Println(t.FieldByNameFunc(func(s string) bool { return s == "Name" }))

	// 返回该类型的字位数。如果该类型的Kind不是Int、Uint、Float或Complex，会panic
	//t.Bits()

	// 返回array类型的长度，如非数组类型将panic
	//t.Len()

	// 返回map类型的键的类型。如非映射类型将panic
	//t.Key()

	// 返回一个channel类型的方向，如非通道类型将会panic
	//t.ChanDir()

	// 如果函数类型的最后一个输入参数是"..."形式的参数，IsVariadic返回true
	// 如果这样，t.In(t.NumIn() - 1)返回参数的隐式的实际类型（声明类型的切片）
	// 如非函数类型将panic
	//t.IsVariadic()

	// 返回func类型的参数个数，如果不是函数，将会panic
	//t.NumIn()

	// 返回func类型的第i个参数的类型，如非函数或者i不在[0, NumIn())内将会panic
	//t.In(0)

	// 返回func类型的返回值个数，如果不是函数，将会panic
	//t.NumOut()

	// 返回func类型的第i个返回值的类型，如非函数或者i不在[0, NumOut())内将会panic
	//t.Out(0)

	// 获取*io.Writer的Type
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	// 获取*os.File的Type
	fileType := reflect.TypeOf((*os.File)(nil))

	// 判断*os.File是否实现了*io.Writer接口
	fmt.Println(fileType.Implements(writerType))

	// 如果该类型的值可以直接赋值给u代表的类型，返回true
	fmt.Println(fileType.AssignableTo(writerType))

	// 如该类型的值可以转换为u代表的类型，返回true
	fmt.Println(fileType.ConvertibleTo(writerType))
}

func exampleValue() {

	hello := Hello{}

	// 返回一个初始化为i接口保管的具体值的Value，ValueOf(nil)返回Value零值
	v := reflect.ValueOf(hello)

	// 返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装
	// 如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值
	reflect.ValueOf(&hello).Elem()

	// 返回v持有的值的类型的Type表示
	// 详细查看exampleType()
	v.Type()

	// 返回v持有的值的类型，如果v是Value零值，返回值为Invalid
	v.Kind()
	fmt.Println("reflect.String:", v.Kind() == reflect.String)

	// 返回v是否持有一个值
	// 如果v是Value零值会返回false，此时v除了IsValid、String、Kind之外的方法都会导致panic
	// 绝大多数函数和方法都永远不返回Value零值。如果某个函数/方法返回了非法的Value，它的文档必须显式的说明具体情况
	fmt.Println(v.IsValid())

	// 判断v持有的值是否为nil
	// v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic
	// 注意IsNil并不总是等价于go语言中值与nil的常规比较。例如：如果v是通过使用某个值为nil的接口调用ValueOf函数创建的，v.IsNil()返回true，但是如果v是Value零值，会panic
	//v.IsNil()

	// 返回v持有的布尔值，如果v的Kind不是Bool会panic
	//v.Bool()

	// 返回v持有的有符号整数（表示为int64），如果v的Kind不是Int、Int8、Int16、Int32、Int64会panic
	//v.Int()

	// 如果v持有值的类型不能无溢出的表示x，会返回true。如果v的Kind不是Int、Int8、Int16、Int32、Int64会panic
	//v.OverflowInt(0)

	// 返回v持有的无符号整数（表示为uint64），如v的Kind不是Uint、Uintptr、Uint8、Uint16、Uint32、Uint64会panic
	//v.Uint()

	// 如果v持有值的类型不能无溢出的表示x，会返回true。如果v的Kind不是Uint、Uintptr、Uint8、Uint16、Uint32、Uint64会panic
	//v.OverflowUint(0)

	// 返回v持有的浮点数（表示为float64），如果v的Kind不是Float32、Float64会panic
	//v.Float()

	// 如果v持有值的类型不能无溢出的表示x，会返回true。如果v的Kind不是Float32、Float64会panic
	//v.OverflowFloat(0)

	// 返回v持有的复数（表示为complex64），如果v的Kind不是Complex64、Complex128会panic
	//v.Complex()

	// 如果v持有值的类型不能无溢出的表示x，会返回true。如果v的Kind不是Complex64、Complex128会panic
	//v.OverflowComplex(0)

	// 将v持有的值作为一个指针返回
	// 本方法返回值不是unsafe.Pointer类型，以避免程序员不显式导入unsafe包却得到unsafe.Pointer类型表示的指针
	// 如果v的Kind不是Chan、Func、Map、Ptr、Slice或UnsafePointer会panic
	// 如果v的Kind是Func，返回值是底层代码的指针，但并不足以用于区分不同的函数；只能保证当且仅当v持有函数类型零值nil时，返回值为0
	// 如果v的Kind是Slice，返回值是指向切片第一个元素的指针。如果持有的切片为nil，返回值为0；如果持有的切片没有元素但不是nil，返回值不会是0
	//v.Pointer()

	// 返回v持有的[]byte类型值。如果v持有的值的类型不是[]byte会panic
	//v.Bytes()

	// 返回v持有的值的字符串表示
	// 因为go的String方法的惯例，Value的String方法比较特别
	// 和其他获取v持有值的方法不同：v的Kind是String时，返回该字符串；v的Kind不是String时也不会panic而是返回格式为"<T value>"的字符串，其中T是v持有值的类型
	//v.String()

	// 返回v持有的接口类型值的数据。如果v的Kind不是Interface会panic
	//v.InterfaceData()

	// 返回v[i:j]（v持有的切片的子切片的Value封装）；如果v的Kind不是Array、Slice或String会panic。如果v是一个不可寻址的数组，或者索引出界，也会panic
	//v.Slice(0, 1)

	// 是Slice的3参数版本，返回v[i:j:k] ；如果v的Kind不是Array、Slice或String会panic。如果v是一个不可寻址的数组，或者索引出界，也会panic
	//v.Slice3(0, 1, 2)

	// 返回v持有值的容量，如果v的Kind不是Array、Chan、Slice会panic
	//v.Cap()

	// 返回v持有值的长度，如果v的Kind不是Array、Chan、Slice、Map、String会panic
	//v.Len()

	// 返回v持有值的第i个元素。如果v的Kind不是Array、Chan、Slice、String，或者i出界，会panic
	//v.Index(0)

	// 返回v持有值里key持有值为键对应的值的Value封装，如果v的Kind不是Map，会panic
	// 如果未找到对应值或者v持有值是nil映射，会返回Value零值
	// key的持有值必须可以直接赋值给v持有值类型的键类型
	//v.MapIndex(reflect.ValueOf("test"))

	// 返回一个包含v持有值中所有键的Value封装的切片，该切片未排序。如果v的Kind不是Map会panic。如果v持有值是nil，返回空切片（非nil）
	//v.MapKeys()

	// 返回v持有的结构体类型值的字段数，如果v的Kind不是Struct会panic
	//v.NumField()

	// 返回结构体的第i个字段（的Value封装）。如果v的Kind不是Struct或i出界会panic
	//v.Field(0)

	// 返回索引序列指定的嵌套字段的Value表示，等价于用索引中的值链式调用本方法，如v的Kind非Struct将会panic
	//v.FieldByIndex([]int{1})

	// 返回该类型名为name的字段（的Value封装）（会查找匿名字段及其子字段），如果v的Kind不是Struct会panic；如果未找到会返回Value零值
	//v.FieldByName("Name")

	// 返回该类型第一个字段名满足match的字段（的Value封装）（会查找匿名字段及其子字段），如果v的Kind不是Struct会panic；如果未找到会返回Value零值
	//v.FieldByNameFunc(func(s string) bool { return s == "Name" })

	// 从v持有的通道接收并返回一个值（的Value封装）。如果v的Kind不是Chan会panic。方法会阻塞直到获取到值
	// 如果返回值x对应于某个发送到v持有的通道的值，ok为true；如果因为通道关闭而返回，x为Value零值而ok为false
	//v.Recv()

	// 尝试从v持有的通道接收一个值，但不会阻塞。如果v的Kind不是Chan会panic
	// 如果方法成功接收到一个值，会返回该值（的Value封装）和true；如果不能无阻塞的接收到值，返回Value零值和false；如果因为通道关闭而返回，返回值x是持有通道元素类型的零值的Value和false
	//v.TryRecv()

	// 向v持有的通道发送x持有的值。如果v的Kind不是Chan，或者x的持有值不能直接赋值给v持有通道的元素类型，会panic
	//v.Send(reflect.ValueOf(5))

	// 尝试向v持有的通道发送x持有的值，但不会阻塞。如果v的Kind不是Chan会panic
	// 如果成功发送会返回true，否则返回false。x的持有值必须可以直接赋值给v持有通道的元素类型
	//v.TrySend(reflect.ValueOf(5))

	// 关闭v持有的通道，如果v的Kind不是Chan会panic
	//v.Close()

	// 使用输入的参数in调用v持有的函数
	// 例如，如果len(in) == 3，v.Call(in)代表调用v(in[0], in[1], in[2])（其中Value值表示其持有值）
	// 如果v的Kind不是Func会panic。它返回函数所有输出结果的Value封装的切片
	// 和go代码一样，每一个输入实参的持有值都必须可以直接赋值给函数对应输入参数的类型
	// 如果v持有值是可变参数函数，Call方法会自行创建一个代表可变参数的切片，将对应可变参数的值都拷贝到里面
	//v.Call([]reflect.Value{reflect.ValueOf("Go"), reflect.ValueOf(10)})

	// 调用v持有的可变参数函数，会将切片类型的in[len(in)-1]（的成员）分配给v的最后的可变参数
	// 例如，如果len(in) == 3，v.Call(in)代表调用v(in[0], in[1], in[2])（其中Value值表示其持有值，可变参数函数的可变参数位置提供一个切片并跟三个点号代表"解切片"）
	// 如果v的Kind不是Func或者v的持有值不是可变参数函数，会panic。它返回函数所有输出结果的Value封装的切片
	// 和go代码一样，每一个输入实参的持有值都必须可以直接赋值给函数对应输入参数的类型
	//v.CallSlice([]reflect.Value{reflect.ValueOf(9), reflect.ValueOf(10), reflect.ValueOf(11)})

	// 返回v持有值的方法集的方法数目
	fmt.Println(v.NumMethod())

	// 返回v持有值类型的第i个方法的已绑定（到v的持有值的）状态的函数形式的Value封装
	// 返回值调用Call方法时不应包含接收者；返回值持有的函数总是使用v的持有者作为接收者（即第一个参数）
	// 如果i出界，或者v的持有值是接口类型的零值（nil），会panic
	//v.Method(0)

	// 返回v的名为name的方法的已绑定（到v的持有值的）状态的函数形式的Value封装
	// 返回值调用Call方法时不应包含接收者；返回值持有的函数总是使用v的持有者作为接收者（即第一个参数）
	// 如果未找到该方法，会返回一个Value零值
	//v.MethodByName("Says")

	// 判断是否是指针，返回是否可以获取v持有值的指针
	// 可以获取指针的值被称为可寻址的。如果一个值是切片或可寻址数组的元素、可寻址结构体的字段、或从指针解引用得到的，该值即为可寻址的
	fmt.Println(v.CanAddr())

	// 返回一个持有指向v持有者的指针的Value封装
	// 如果v.CanAddr()返回false，调用本方法会panic。Addr一般用于获取结构体字段的指针或者切片的元素（的Value封装）以便调用需要指针类型接收者的方法
	//v.Addr()

	// 返回指向v持有数据的地址的指针（表示为uintptr）以用作高级用途，如果v不可寻址会panic
	//v.UnsafeAddr()

	// 判断是否interface，如果返回true，v可以不导致panic的调用Interface方法
	fmt.Println(v.CanInterface())

	// 返回v当前持有的值（表示为/保管在interface{}类型），如果v是通过访问非导出结构体字段获取的，会导致panic
	fmt.Println(v.Interface())

	// 判断是否可以设置值
	// 如果v持有的值可以被修改，会返回true。只有一个Value持有值可以被寻址同时又不是来自非导出字段时，它才可以被修改
	// 如果返回false，调用Set或任何限定类型的设置函数（如SetBool、SetInt64）都会panic
	fmt.Println(v.CanSet())

	// 设置bool值。如果v的Kind不是Bool或者v.CanSet()返回false，会panic
	//v.SetBool(true)

	// 设置int值。如果v的Kind不是Int、Int8、Int16、Int32、Int64之一或者v.CanSet()返回false，会panic
	//v.SetInt(0)

	// 设置uint值。如果v的Kind不是Uint、Uintptr、Uint8、Uint16、Uint32、Uint64或者v.CanSet()返回false，会panic
	//v.SetUint(0)

	// 设置float值。如果v的Kind不是Float32、Float64或者v.CanSet()返回false，会panic
	//v.SetFloat(0)

	// 设置complex值。如果v的Kind不是Complex64、Complex128或者v.CanSet()返回false，会panic
	//v.SetComplex(0)

	// 设置[]byte值。如果v持有值不是[]byte类型或者v.CanSet()返回false，会panic
	//v.SetBytes([]byte("Hello World!"))

	// 设置string值。如果v的Kind不是String或者v.CanSet()返回false，会panic
	//v.SetString("Hello World!")

	// 设置指针值。如果v的Kind不是UnsafePointer或者v.CanSet()返回false，会panic
	//v.SetPointer(&Hello{"Go", 12, []string{"Hello", "World"}})

	// 设定值的容量。如果v的Kind不是Slice或者n出界（小于长度或超出容量），将导致panic
	//v.SetCap(10)

	// 设定值的长度。如果v的Kind不是Slice或者n出界（小于零或超出容量），将导致panic
	//v.SetLen(10)

	// 用来给v的映射类型持有值添加/修改键值对，如果val是Value零值，则是删除键值对。如果v的Kind不是Map，或者v的持有值是nil，将会panic
	// key的持有值必须可以直接赋值给v持有值类型的键类型。val的持有值必须可以直接赋值给v持有值类型的值类型
	//v.SetMapIndex(reflect.ValueOf("test"), reflect.ValueOf("Hello World!"))

	// 将v的持有值修改为x的持有值。如果v.CanSet()返回false，会panic。x的持有值必须能直接赋给v持有值的类型
	//v.Set(reflect.ValueOf("Hello"))
}

func exampleMake() {

	var c chan int
	var fn func(int) int
	var m map[string]string
	var sl = []int{1, 2, 3}

	// 创建一个元素类型为typ、有buffer个缓存的channel类型的Value值
	reflect.MakeChan(reflect.TypeOf(c), 1)

	// 创建一个具有给定类型、包装函数fn的函数的Value封装。
	// 当被调用时，该函数会将提供给它的参数转化为Value切片
	// 执行results := fn(args)
	// 将results中每一个result依次排列作为返回值
	reflect.MakeFunc(reflect.TypeOf(fn), func(args []reflect.Value) (results []reflect.Value) { return []reflect.Value{args[0]} })

	// 创建一个特定map类型的Value值
	reflect.MakeMap(reflect.TypeOf(m))

	// 创建一个具有n个初始空间的map类型的Value值
	reflect.MakeMapWithSize(reflect.TypeOf(m), 10)

	// 创建一个新申请的元素类型为typ，长度len容量cap的slice类型的Value值
	reflect.MakeSlice(reflect.ValueOf(sl).Type(), reflect.ValueOf(sl).Len(), reflect.ValueOf(sl).Cap())

	// 返回元素类型为t、方向为dir的channel类型的Type
	// 运行时GC强制将通道的元素类型的大小限定为64kb
	// 如果t的尺寸大于或等于该限制，本函数将会panic
	reflect.ChanOf(reflect.BothDir, reflect.TypeOf(c))

	// 返回给定参数和结果类型的func类型的Type
	// 例如，如果 k 表示 int 并且 e 表示字符串，则 FuncOf([]Type{k}, []Type{e}, false) 表示 func(int) string
	// 可变参数variadic 控制函数是否可变。如果len(in)-1 不代表切片并且可变参数为true，则 FuncOf 会panic
	reflect.FuncOf([]reflect.Type{reflect.TypeOf(15)}, []reflect.Type{reflect.TypeOf("Hello")}, false)

	// 使用给定的键和元素类型返回map类型的Type
	// 例如，如果 k 表示 int 并且 e 表示字符串，则 MapOf(k, e) 表示 map[int]string
	// 如果键类型不是有效的映射键类型（也就是说，如果它不执行 Go 的==运算符），则 MapOf会panic
	reflect.MapOf(reflect.TypeOf(15), reflect.TypeOf("Hello"))

	// 使用给定的计数和元素类型返回数组类型的Type
	// 例如，如果 t 表示 int ，则 ArrayOf(5, t) 表示 [5]int
	// 如果生成的类型会比可用的地址空间大，ArrayOf 会panic
	reflect.ArrayOf(5, reflect.TypeOf(15))

	// 返回元素类型为 t 的slice类型的Type
	// 例如，如果 t 表示 int ， SliceOf(t) 表示 []int
	reflect.SliceOf(reflect.ValueOf(sl).Type())

	// 返回包含字段的结构类型的Type
	// 偏移量和索引字段被忽略和计算，就像编译器一样
	// StructOf 目前不会为嵌入字段生成封装器方法。未来版本中可能会取消此限制
	reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"age"`,
		},
	})
}

func example() {

	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"age"`,
		},
	})

	// 返回一个Value类型值，该值持有一个指向类型为typ的新申请的零值的指针，返回值的Type为PtrTo(typ)
	reflect.New(typ)

	// 返回类型t的指针的类型
	reflect.PtrTo(typ)

	// 返回一个持有类型typ的零值的Value
	// 注意持有零值的Value和Value零值是两回事
	// Value零值表示不持有任何值，例如Zero(TypeOf(42))返回一个Kind为Int、值为0的Value
	// Zero的返回值不能设置也不会寻址
	reflect.Zero(typ)

	// 返回一个函数，它交换提供的 slice 中的元素
	// 如果提供的接口不是切片，Swapper会panic
	sl := []int{1, 5, 9, 3, 7}
	reflect.Swapper(sl)

	// 将src中的值拷贝到dst，直到src被耗尽或者dst被装满，要求这二者都是slice或array，且元素类型相同
	var dst []byte
	var src = []byte("Hello World!")
	reflect.Copy(reflect.ValueOf(dst), reflect.ValueOf(src))

	// 用来判断两个值是否深度一致：除了类型相同；在可以时（主要是基本类型）会使用==；但还会比较array、slice的成员，map的键值对，结构体字段进行深入比对
	// map的键值对，对键只使用==，但值会继续往深层比对，可以正确处理循环的类型
	// 函数类型只有都会nil时才相等；空切片不等于nil切片；还会考虑array、slice的长度、map键值对数
	reflect.DeepEqual("1", 1)
}
