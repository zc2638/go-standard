package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
)

// smtp包实现了简单邮件传输协议（SMTP）
// 本包实现了以下扩展：8BITMIME、AUTH、STARTTLS
func main() {

	// 邮件客户端发送邮件示例
	exampleDial()
	// 使用网络链接发送邮件示例
	exampleClient()
	// 简单的发送邮件示例
	exampleSendMail()
}

func exampleDial() {

	// 返回一个连接到地址为addr的SMTP服务器的*smtp.Client；addr必须包含端口号
	c, err := smtp.Dial("mail.example.com:25")
	if err != nil {
		log.Fatal(err)
	}

	// 发送给服务端一个HELO或EHLO命令
	// 本方法只有使用者需要控制使用的本地主机名时才应使用，否则程序会将本地主机名设为“localhost”，Hello方法只能在最开始调用
	if err := c.Hello("localhost"); err != nil {
		log.Fatal(err)
	}

	// 返回服务端是否支持某个扩展，扩展名是大小写不敏感的
	// 如果扩展被支持，方法还会返回一个包含指定给该扩展的各个参数的字符串
	fmt.Println(c.Extension("8BITMIME"))

	// 返回一个实现了PLAIN身份认证机制的Auth接口
	// 返回的接口使用给出的用户名和密码，通过TLS连接到主机认证，采用identity为身份管理和行动（通常应设identity为""，以便使用username为身份）
	auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")

	// 返回一个实现了CRAM-MD5身份认证机制的Auth接口
	// 返回的接口使用给出的用户名和密码，采用响应——回答机制与服务端进行身份认证
	smtp.CRAMMD5Auth("user@example.com", "password")

	// 使用提供的认证机制进行认证。失败的认证会关闭该连接。只有服务端支持AUTH时，本方法才有效。（但是不支持时，调用会成功）
	if err := c.Auth(auth); err != nil {
		log.Fatal(err)
	}

	// 发送STARTTLS命令，并将之后的所有数据往来加密
	// 只有服务器附加了STARTTLS扩展，这个方法才有效
	c.StartTLS(&tls.Config{})

	// 检查一个邮箱地址在其服务器是否合法，如果合法会返回nil
	// 但非nil的返回值并不代表不合法，因为许多服务器出于安全原因不支持这种查询
	c.Verify("user@example.com")

	// 发送发件人邮箱地址
	// 发送MAIL命令和邮箱地址from到服务器
	// 如果服务端支持8BITMIME扩展，本方法会添加BODY=8BITMIME参数
	// 方法初始化一次邮件传输，后应跟1到多个Rcpt方法的调用
	if err := c.Mail("sender@example.org"); err != nil {
		log.Fatal(err)
	}

	// 发送收件人邮箱地址
	// 发送RCPT命令和邮箱地址to到服务器
	// 调用Rcpt方法之前必须调用了Mail方法，之后可以再一次调用Rcpt方法，也可以调用Data方法
	if err := c.Rcpt("recipient@example.net"); err != nil {
		log.Fatal(err)
	}

	// 发送DATA指令到服务器并返回一个io.WriteCloser，用于写入邮件信息
	// 调用者必须在调用c的下一个方法之前关闭这个io.WriteCloser。方法必须在一次或多次Rcpt方法之后调用
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	// 关闭
	defer wc.Close()

	// 写入邮件内容
	if _, err := wc.Write([]byte("Hello World!")); err != nil {
		log.Fatal(err)
	}

	// 发送QUIT命令并关闭到服务端的连接
	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}


	// 关闭连接
	//c.Close()

	// 向服务端发送REST命令，中断当前的邮件传输
	//c.Reset()
}

func exampleClient() {

	addr := "mail.example.com:25"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("Dialing Error:", err)
	}

	//分解主机端口字符串
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		log.Fatal(err)
	}

	// 使用已经存在的连接conn和作为服务器名的host（用于身份认证）来创建一个*smtp.Client
	smtp.NewClient(conn, host)
}

func exampleSendMail() {

	// 邮件服务器地址
	addr := "mail.example.com:25"

	// 返回一个实现了PLAIN身份认证机制的Auth接口
	// 返回的接口使用给出的用户名和密码，通过TLS连接到主机认证，采用identity为身份管理和行动（通常应设identity为""，以便使用username为身份）
	auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")

	// 发件人
	from := "sender@example.org"

	// 收件人集合
	to := []string{"recipient@example.net"}

	// 邮件内容
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	// 发送邮件，封装了smtp.Client
	// 连接到addr指定的服务器；如果支持会开启TLS；如果支持会使用a认证身份；然后以from为邮件源地址发送邮件msg到目标地址to（可以是多个目标地址：群发）
	if err := smtp.SendMail(addr, auth, from, to, msg); err != nil {
		log.Fatal(err)
	}
}
