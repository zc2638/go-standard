package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
)

// xml包实现了XML 1.0解析
func main() {

	example()
	example2()
}

type Hello struct {
	XMLName   xml.Name `xml:"xml"`        // 指定xml头标签,不指定默认为结构体名称
	Id        int      `xml:"id,attr"`    // 指定xml头标签上的元素
	FullName  string   `xml:"fullName"`   // 指定标签名
	FirstName string   `xml:"name>first"` // 指定上级标签
	LastName  string   `xml:"name>last"`  // 指定上级标签
	Sex       int      `xml:"sex"`        // 指定标签名
	Comment   string   `xml:",comment"`   // 声明注释
}

func example() {

	// 声明结构体
	hello := Hello{
		Id:        12,
		FullName:  "zc",
		FirstName: "z",
		LastName:  "c",
		Sex:       1,
		Comment:   "test notes",
	}

	// xml编码
	b, err := xml.Marshal(hello)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	// 声明一个空的结构体
	origin := Hello{}
	if err := xml.Unmarshal(b, &origin); err != nil {
		log.Fatal(err)
	}
	fmt.Println(origin)

	// 自定义前缀和index编码
	bi, err := xml.MarshalIndent(hello, "\r", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bi))
}

func example2() {

	// 声明buffer
	var buf bytes.Buffer
	// 声明结构体
	hello := Hello{
		Id:        12,
		FullName:  "zc",
		FirstName: "z",
		LastName:  "c",
		Sex:       1,
		Comment:   "test notes",
	}

	// 创建一个写入w的*Encoder
	e := xml.NewEncoder(&buf)

	// 自定义前缀和index
	e.Indent("\r", "  ")

	// 自定义xml头编码
	err := e.EncodeElement(hello, xml.StartElement{
		Name: xml.Name{
			Space: "test",   // xmlns
			Local: "newXml", // xml头名称
		},
		Attr: nil, // 设置xml头元素
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf.Bytes()))
}
