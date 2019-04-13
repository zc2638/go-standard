package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
	"net/mail"
	"strings"
)

// mime实现了MIME的部分规定
func main() {

	// 解析multipart内容，适用于HTTP和常见浏览器生成的multipart主体
	exampleMultipart()
	// 读/写multipart内容
	exampleMultipartRW()
	// quoted-printable编码/解码
	exampleQuotedprintable()
	//RFC2047编码字
	exampleWord()
}

func exampleWord() {

	// 初始化一个RFC2047编码字编码器
	e := mime.WordEncoder('q')
	// 编码
	// 如果s是没有特殊字符的ASCII码，返回原始字符串
	// 提供的字符集是s的IANA字符集名称不区分大小写
	str := e.Encode("utf-8", "¡Hola, señor!")
	fmt.Println(str)
	fmt.Println(e.Encode("utf-8", "Hello World!"))

	// 解码包含RFC 2047编码字的MIME头
	d := new(mime.WordDecoder)
	// 解码
	s, err := d.Decode(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}

func exampleMultipart() {

	// 声明邮件信息
	msg := &mail.Message{
		Header: map[string][]string{
			"Content-Type": {"multipart/mixed; boundary=foo"},
		},
		Body: strings.NewReader(
			"--foo\r\nFoo: one\r\n\r\nA section\r\n" +
				"--foo\r\nFoo: two\r\n\r\nAnd another\r\n" +
				"--foo--\r\n"),
	}

	// 解析Content-Type
	// 解析一个媒体类型值以及可能的参数
	// 媒体类型值一般应为Content-Type和Conten-Disposition头域的值
	// 成功的调用会返回小写字母、去空格的媒体类型和一个非空的map。返回的map映射小写字母的属性和对应的属性值
	mediaType, params, err := mime.ParseMediaType(msg.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}

	// 判断字符串是否有指定前缀
	if strings.HasPrefix(mediaType, "multipart/") {

		// 使用给出的MIME边界和reader创建一个新的reader
		// 例如：边界boundary值为foo，会使用foo解析传入的reader内容
		mr := multipart.NewReader(msg.Body, params["boundary"])
		for {
			// 返回multipart的下一个记录或者返回错误。如果没有更多记录会返回io.EOF
			p, err := mr.NextPart()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatal(err)
			}

			// 读取内容
			slurp, err := ioutil.ReadAll(p)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Part %q: %q\n", p.Header.Get("Foo"), slurp)
		}
	}
}

func exampleMultipartRW() {

	var buf bytes.Buffer

	// 初始化一个随机边界的writer
	w := multipart.NewWriter(&buf)

	// 查看边界
	fmt.Println(w.Boundary())

	// 设置边界
	// 必须在创建任何记录之前调用，boundary只能包含特定的ascii字符，并且长度应在1-69字节之间
	if err := w.SetBoundary("testing"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(w.Boundary())

	// 查看对应的HTTP multipart请求的Content-Type的值，多以multipart/form-data起始
	fmt.Println(w.FormDataContentType())

	// 使用给出的属性名调用CreatePart方法
	wt, err := w.CreateFormField("test")
	if err != nil {
		log.Fatal(err)
	}

	// 写入test
	if _, err := wt.Write([]byte("Hello World!")); err != nil {
		log.Fatal(err)
	}
	// 关闭
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}

	// 根据指定边界创建reader
	r := multipart.NewReader(&buf, "testing")
	for {
		// 返回multipart的下一个记录或者返回错误。如果没有更多记录会返回io.EOF
		p, err := r.NextPart()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}

		// 读取内容
		b, err := ioutil.ReadAll(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}
}

func exampleQuotedprintable() {

	// 声明buffer
	var buf bytes.Buffer

	// 初始化一个writer
	w := quotedprintable.NewWriter(&buf)

	// 写入，行长限制为76个字符。编码字节不必刷新，直到Writer关闭
	if _, err := w.Write([]byte("Hello, Gophers! = \t")); err != nil {
		log.Fatal(err)
	}

	// 关闭
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	// 初始化一个reader
	r := quotedprintable.NewReader(&buf)

	// 读取内容
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
