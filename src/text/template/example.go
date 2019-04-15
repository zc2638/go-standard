package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

// template包实现了数据驱动的用于生成文本输出的模板
// 如果要生成HTML格式的输出，参见html/template包，该包提供了和本包相同的接口，但会自动将输出转化为安全的HTML格式输出，可以抵抗一些网络攻击。
// 通过将模板应用于一个数据结构（即该数据结构作为模板的参数）来执行，来获得输出。模板中的注释引用数据接口的元素（一般如结构体的字段或者字典的键）来控制执行过程和获取需要呈现的值。模板执行时会遍历结构并将指针表示为'.'（称之为"dot"）指向运行过程中数据结构的当前位置的值。
// 用作模板的输入文本必须是utf-8编码的文本。"Action"—数据运算和控制单位—由"{{"和"}}"界定；在Action之外的所有文本都不做修改的拷贝到输出中。Action内部不能有换行，但注释可以有换行。
// 经解析生成模板后，一个模板可以安全的并发执行
func main() {

	// 基础
	example()
	// 设置action分界字符串
	exampleDelims()
	// 自动解析
	exampleAuto()
}

func example() {

	// 声明buffer
	var buf bytes.Buffer
	var jsBuf bytes.Buffer

	// 声明内容
	var data = []byte("<div>Hello World!</div>")
	var p = []byte("<p>test</p>")
	var jsData = []byte("<script>alert('Hello World!')</script>")
	var url = []byte("http://www.baidu.com/test?a=1&b=2")

	// 向w中写入b的HTML转义等价表示
	template.HTMLEscape(&buf, data)
	fmt.Println(string(buf.Bytes()))

	// 返回s的HTML转义等价表示字符串
	str := template.HTMLEscapeString(string(data))
	fmt.Println(str)

	// 返回其所有参数文本表示的HTML转义等价表示字符串
	strs := template.HTMLEscaper(string(data), string(p))
	fmt.Println(string(strs))

	// 向w中写入b的JavaScript转义等价表示
	template.JSEscape(&jsBuf, jsData)
	fmt.Println(string(jsBuf.Bytes()))

	// 返回s的JavaScript转义等价表示字符串
	jsStr := template.JSEscapeString(string(jsData))
	fmt.Println(jsStr)

	// 返回其所有参数文本表示的JavaScript转义等价表示字符串
	jsStrs := template.JSEscaper(string(jsData), string(p))
	fmt.Println(jsStrs)

	// 返回其所有参数文本表示的可以嵌入URL查询的转义等价表示字符串
	urlStr := template.URLQueryEscaper(string(url))
	fmt.Println(urlStr)

	var testBuf bytes.Buffer
	var (
		test    = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	)
	// 声明数据集
	var dataMap = []string{"张三", "李四", "王五", "赵六", "黑七"}

	// 定义了函数名字符串到函数的映射，每个函数都必须有1到2个返回值，如果有2个则后一个必须是error接口类型；如果有2个返回值的方法返回的error非nil，模板执行会中断并返回给调用者该错误
	var funcs = template.FuncMap{"join": strings.Join}

	// 创建一个名为name的模板
	t := template.New("test")

	// 返回模板t的名字
	t.Name()

	// 向模板t的函数字典里加入参数funcMap内的键值对。如果funcMap某个键值对的值不是函数类型或者返回值不符合要求会panic
	// 但是，可以对t函数列表的成员进行重写。方法返回t以便进行链式调用
	t.Funcs(funcs)

	// 返回模板的一个副本，包括所有相关联的模板
	// 模板的底层表示树并未拷贝，而是拷贝了命名空间，因此拷贝调用Parse方法不会修改原模板的命名空间
	// Clone方法用于准备模板的公用部分，向拷贝中加入其他关联模板后再进行使用
	t.Clone()

	// 创建一个和t关联的名字为name的模板并返回它。这种可以传递的关联允许一个模板使用template action调用另一个模板
	t.New("hello")

	// 返回与t相关联的模板的切片，包括t自己
	t.Templates()

	// 将字符串text解析为模板。嵌套定义的模板会关联到最顶层的t
	// Parse可以多次调用，但只有第一次调用可以包含空格、注释和模板定义之外的文本
	// 如果后面的调用在解析后仍剩余文本会引发错误、返回nil且丢弃剩余文本；如果解析得到的模板已有相关联的同名模板，会覆盖掉原模板
	out, _ := t.Parse(test)

	// 解析filenames指定的文件里的模板定义并将解析结果与t关联。如果发生错误，会停止解析并返回nil，否则返回(t, nil)。至少要提供一个文件
	t.ParseFiles("testdata/htmlTemplateTest.html")

	// 解析匹配pattern的文件里的模板定义并将解析结果与t关联。如果发生错误，会停止解析并返回nil，否则返回(t, nil)。至少要存在一个匹配的文件
	t.ParseGlob("testdata/htmlTemplate/*.html")

	// 用于包装返回(*Template, error)的函数/方法调用，它会在err非nil时panic，一般用于变量初始化
	mt := template.Must(out.Clone())

	// 将解析好的模板应用到data上，并将输出写入wr
	// 如果执行时出现错误，会停止执行，但有可能已经写入wr部分数据。模板可以安全的并发执行
	_ = out.Execute(&testBuf, dataMap)
	fmt.Println(testBuf.String())

	mOut, _ := mt.Parse(overlay)
	_ = mOut.Execute(&testBuf, dataMap)
	fmt.Println(testBuf.String())
}

func exampleDelims() {

	var buf bytes.Buffer

	var text = "<<.Greeting>> {{.Name}}"

	data := struct {
		Greeting string
		Name     string
	}{
		Greeting: "Hello",
		Name:     "Joe",
	}

	// 创建模板
	t := template.New("tpl")

	// 用于设置action的分界字符串，应用于之后的Parse、ParseFiles、ParseGlob方法
	// 嵌套模板定义会继承这种分界符设置。空字符串分界符表示相应的默认分界符：{{或}}。返回值就是t，以便进行链式调用
	t = t.Delims("<<", ">>")
	t = template.Must(t.Parse(text))

	if err := t.Execute(&buf, data); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func exampleAuto() {

	var (
		text = `{{define "T"}}Hello, {{.}}!{{end}}`
		data = "<script>alert('you have been pwned')</script>\n"
	)
	t, err := template.New("test").Parse(text)
	if err != nil {
		log.Fatal(err)
	}

	// 类似Execute，但是使用名为name的t关联的模板产生输出
	if err := t.ExecuteTemplate(os.Stdout, "T", data); err != nil {
		log.Fatal(err)
	}

}