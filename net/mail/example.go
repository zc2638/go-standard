package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"strings"
)

// mail包实现了邮件的解析
func main() {

	// 解析邮箱地址
	exampleParseAddress()

	// 读取邮件内容
	exampleReadMessage()
}

func exampleParseAddress() {

	var list = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"

	// 解析一串邮箱地址
	emails, err := mail.ParseAddressList(list)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range emails {
		fmt.Println(v.Name, v.Address)
	}


	// 解析单个邮箱地址
	address, err := mail.ParseAddress("Alice <alice@example.com>")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Name, address.Address)
}

func exampleReadMessage() {

	msg := `Date: Mon, 23 Jun 2015 11:40:36 -0400
From: Gopher <from@example.com>
To: Another Gopher <to@example.com>
Subject: Gophers at Gophercon

Message body
`

	// 解析邮件信息
	// 从reader读取一个邮件信息，会解析邮件头域，消息主体可以从r/msg.Body中读取
	m, err := mail.ReadMessage(strings.NewReader(msg))
	if err != nil {
		log.Fatal(err)
	}

	// 邮件头域
	header := m.Header

	// 回键key对应的第一个值，如果没有对应值，将返回空字符串
	fmt.Println("Date:", header.Get("Date"))
	fmt.Println("From:", header.Get("From"))
	fmt.Println("To:", header.Get("To"))
	fmt.Println("Subject:", header.Get("Subject"))


	// 将键key对应的值（字符串）作为邮箱地址列表解析并返回
	fmt.Println(header.AddressList("From"))

	// 读取body内容
	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}