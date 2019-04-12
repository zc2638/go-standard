package main

import (
	"encoding/asn1"
	"fmt"
	"log"
)

// asn1包实现了DER编码的ASN.1数据结构的解析
func main() {

	hello := Hello{
		Num: 12,
		Str: "Go",
	}

	// ASN1序列化
	/*
		结构体序列化标签可选
		ia5:           使字符串序列化为ASN.1 IA5String类型
		omitempty:     使空切片被跳过
		printable:     使字符串序列化为ASN.1 PrintableString类型
		utf8:          使字符串序列化为ASN.1 UTF8字符串
	 */
	b, err := asn1.Marshal(hello)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b)

	var h Hello

	// ASN1反序列化，返回未操作的字节切片和一个error
	/*
		结构体反序列化标签可选
		application    指明使用了APPLICATION标签
		default:x      设置一个可选整数字段的默认值
		explicit       给一个隐式的标签设置一个额外的显式标签
		optional       标记字段为ASN.1 OPTIONAL的
		set            表示期望一个SET而不是SEQUENCE类型
		tag:x          指定ASN.1标签码，隐含ASN.1 CONTEXT SPECIFIC
	 */
	if _, err := asn1.Unmarshal(b, &h); err != nil {
		log.Fatal(err)
	}
	fmt.Println(h)
}

type Hello struct {
	Num int
	Str string
}