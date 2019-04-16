package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// os包提供了操作系统函数的不依赖平台的接口
// 设计为Unix风格的，虽然错误处理是go风格的；失败的调用会返回错误值而非错误码
// 通常错误值里包含更多信息。例如，如果某个使用一个文件名的调用（如Open、Stat）失败了，打印错误时会包含该文件名，错误类型将为*PathError，其内部可以解包获得更多信息
// os包的接口规定为在所有操作系统中都是一致的。非公用的属性可以从操作系统特定的syscall包获取
func main() {

	// 获取系统基本信息
	example()
	// 创建文件夹
	exampleDir()
	// 文件操作
	exampleFile()
	// 文件改变
	exampleChange()
	// 文件操作模式
	exampleFileMode()
	// 环境变量
	exampleEnv()
	// 进程操作
	exampleProcess()
	// 获取命令行参数
	exampleArgs()
}

const (
	DirPath = "testdata/"
	FilePath = "testdata/test_os.txt"
	NewFilePath = "testdata/test_os_new.txt"
)

func example() {

	// 返回内核提供的主机名
	fmt.Println(os.Hostname())

	// 返回底层的系统内存页的尺寸
	fmt.Println("agesize:", os.Getpagesize())

	// 返回调用者的用户ID
	fmt.Println("uid:", os.Getuid())

	// 返回调用者的有效用户ID
	fmt.Println("euid:", os.Geteuid())

	// 返回调用者的组ID
	fmt.Println("gid:", os.Getgid())

	// 返回调用者的有效组ID
	fmt.Println("egid:",os.Getegid())

	// 返回调用者所属的所有用户组的组ID
	fmt.Println(os.Getgroups())

	// 返回调用者所在进程的进程ID
	fmt.Println("pid:", os.Getpid())

	// 返回调用者所在进程的父进程的进程ID
	fmt.Println("ppid:",os.Getppid())

	// 返回表示环境变量的格式为"key=value"的字符串的切片拷贝
	fmt.Println(os.Environ())

	// 返回对应当前工作目录的根路径。如果当前目录可以经过多条路径抵达（因为硬链接），Getwd会返回其中一个
	fmt.Println(os.Getwd())

	// 返回用于保管临时文件的默认目录
	fmt.Println(os.TempDir())

	// 返回用于用户的缓存数据的默认根目录。用户应该在这个目录中创建自己的特定于应用程序的子目录，并使用它
	fmt.Println(os.UserCacheDir())

	// 用户文件夹路径
	fmt.Println(os.UserHomeDir())
}

func exampleDir() {

	// 使用指定的权限和名称创建一个目录。如果出错，会返回*PathError底层类型的错误
	if err := os.Mkdir(DirPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// 使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回nil，否则返回错误
	// 权限位perm会应用在每一个被本函数创建的目录上
	// 如果path指定了一个已经存在的目录，MkdirAll不做任何操作并返回nil
	if err := os.MkdirAll(DirPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func exampleFile() {

	// 以读取的方式打开文件，文件不存在会报错
	//os.Open(FilePath)

	// 以读写方式打开文件，并且清空原始内容，如果文件不存在以0666操作模式创建文件
	//os.Create(FilePath)

	// 改变文件的大小，它不会改变I/O的当前位置。 如果截断文件，多出的部分就会被丢弃。如果出错，错误底层类型是*PathError
	//os.Truncate(FilePath, 1024)

	// 删除指定的文件或目录。如果出错，会返回*PathError底层类型的错误
	//os.Remove(FilePath)

	// 删除path指定的文件，或目录及它包含的任何下级对象。它会尝试删除所有东西，除非遇到错误并返回。如果path指定的对象不存在，会返回nil而不返回错误
	//os.RemoveAll(FilePath)

	// 创建硬链接。如果出错，会返回* LinkError底层类型的错误
	//os.Link(FilePath, NewFilePath)

	// 创建软连接。如果出错，会返回* LinkError底层类型的错误
	//os.Symlink(FilePath, NewFilePath)

	// 获取指定的软链接文件指向的文件的路径。如果出错，会返回*PathError底层类型的错误
	//os.Readlink(NewFilePath)

	// 以读写方式打开文件，并且内容写入方式为添加，如果文件不存在以0755操作模式创建文件
	f, err := os.OpenFile(FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// 关闭文件
	defer f.Close()

	// 写入内容（向文件内添加内容）
	if _, err := f.Write([]byte("appended some data\n")); err != nil {
		log.Fatal(err)
	}
}

func exampleChange() {

	// 判断文件是否存在
	if _, err := os.Stat(FilePath); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file does not exist")
		} else {
			log.Fatal(err)
		}
	}

	// 修改指定文件的mode操作模式
	// 如果name指定的文件是一个符号链接，它会修改该链接的目的地文件的mode。如果出错，会返回*PathError底层类型的错误
	if err := os.Chmod(FilePath, 0644); err != nil {
		log.Fatal(err)
	}

	// 修改指定文件的用户id和组id
	// 如果name指定的文件是一个符号链接，它会修改该链接的目的地文件的用户id和组id。如果出错，会返回*PathError底层类型的错误
	if err := os.Chown(FilePath, 501, 20); err != nil {
		log.Fatal(err)
	}

	// 修改指定文件的用户id和组id
	// 如果name指定的文件是一个符号链接，它会修改该符号链接自身的用户id和组id。如果出错，会返回*PathError底层类型的错误
	if err := os.Lchown(FilePath, 501, 20); err != nil {
		log.Fatal(err)
	}

	mtime := time.Date(2019, time.February, 1, 3, 4, 5, 0, time.UTC)
	atime := time.Date(2019, time.March, 2, 4, 5, 6, 0, time.UTC)

	// 修改指定文件的访问时间和修改时间，类似Unix的utime()或utimes()函数
	// 底层的文件系统可能会截断/舍入时间单位到更低的精确度。如果出错，会返回*PathError底层类型的错误
	if err := os.Chtimes(FilePath, atime, mtime); err != nil {
		log.Fatal(err)
	}

	// 将当前工作目录修改为dir指定的目录。如果出错，会返回*PathError底层类型的错误
	//if err := os.Chdir("/"); err != nil {
	//	log.Fatal(err)
	//}
}
func exampleFileMode() {

	// 返回一个描述指定文件的FileInfo
	// 如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接的信息，本函数不会试图跳转该链接。如果出错，返回的错误值为*PathError类型
	fi, err := os.Lstat(FilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("permissions: %#o\n", fi.Mode().Perm()) // 0400, 0777, etc.
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")
	case mode.IsDir():
		fmt.Println("directory")
	case mode&os.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode&os.ModeNamedPipe != 0:
		fmt.Println("named pipe")
	}
}
func exampleEnv() {

	// 设置名为key的环境变量（临时）。如果出错会返回该错误
	if err := os.Setenv("NAME", "Gopher"); err != nil {
		log.Fatal(err)
	}

	// 检索并返回名为key的环境变量的值。如果不存在该环境变量会返回空字符串
	fmt.Println(os.Getenv("NAME"))

	// 检索并返回名为key的环境变量的值 和 是否存在的bool值。如果不存在布尔值为false
	fmt.Println(os.LookupEnv("NAME"))

	// 删除所有环境变量
	//os.Clearenv()

	// 使用指定函数替换s中的${var}或$var。例如，os.ExpandEnv(s)等价于os.Expand(s, os.Getenv)
	fmt.Println(os.Expand("Hello $NAME!", func(s string) string { return "Gopher" }))

	// 替换s中的${var}或$var为名为var 的环境变量的值。引用未定义环境变量会被替换为空字符串
	fmt.Println(os.ExpandEnv("Hello ${NAME}!"))

	// 取消设置单个环境变量
	if err := os.Unsetenv("NAME"); err != nil {
		log.Fatal(err)
	}
}

func exampleProcess() {

	// 初始化一个 保管将被StartProcess函数用于一个新进程的属性
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}, //其他变量如果不清楚可以不设定
	}

	// 使用提供的属性、程序名、命令行参数开始一个新进程
	// StartProcess函数是一个低水平的接口
	// os/exec包提供了高水平的接口，应该尽量使用该包。如果出错，错误的底层类型会是*PathError
	p, err := os.StartProcess("/usr/bin/vim", []string{"/usr/bin/vim", "temp.txt"}, attr)
	if err != nil {
		log.Fatal(err)
	}

	// 根据进程id查找一个运行中的进程
	p, err = os.FindProcess(p.Pid)
	if err != nil {
		log.Fatal(err)
	}

	// 立刻退出进程
	if err := p.Kill(); err != nil {
		log.Fatal(err)
	}
}

func exampleArgs() {

	// 获取命令行参数
	// 第一个参数是命令本身，所以从第二个开始截取
	args := os.Args[1:]
	fmt.Println(args)
}