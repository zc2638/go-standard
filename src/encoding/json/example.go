package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// json包实现了json对象的编解码
func main() {

	// 声明内容
	var data = []byte(`{"message": "Hello Gopher!"}`)
	// 验证内容是否符合json格式
	json.Valid(data)

	example()
	example2()
}

type Hello struct {
	Name string
	Sex  int
}

func example() {

	// 声明一个结构体
	hello := Hello{"Gopher", 1}
	// 进行json编码
	b, err := json.Marshal(hello)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	// 声明一个空的结构体
	originHello := Hello{}
	// 将json内容解码到指定结构中
	if err := json.Unmarshal(b, &originHello); err != nil {
		log.Fatal(err)
	}
	fmt.Println(originHello)
}

func example2() {

	// 声明buffer
	var buf bytes.Buffer
	// 声明一个结构体
	hello := Hello{"<div>World</div>", 2}

	// 创建一个将数据写入w的*Encoder
	e := json.NewEncoder(&buf)

	// 是否转义html标签, 例如：&, <, > 为\u0026, \u003c, \u003e
	e.SetEscapeHTML(false)

	// 设置缩进(可以随意设置缩进内容)
	e.SetIndent("", "")

	// 进行json编码
	if err := e.Encode(hello); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf.Bytes()))

	// 声明一个空的结构体
	originHello := Hello{}

	// 创建一个从r读取并解码json对象的*Decoder
	d := json.NewDecoder(&buf)

	// 返回保存在json.Decoder缓存里数据的读取器，该返回值在下次调用Decode方法之前有效
	d.Buffered()

	// 设置此项，如果发现不能识别的字符串将导致解码失败
	d.DisallowUnknownFields()

	// 设置此项，当接收端是interface{}接口时将json数字解码为Number类型而不是float64类型
	d.UseNumber()

	// 返回是否存在其他元素
	d.More()

	// 将json内容解码到指定结构中
	if err := d.Decode(&originHello); err != nil {
		log.Fatal(err)
	}
	fmt.Println(originHello)

}