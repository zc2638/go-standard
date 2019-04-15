package main

import (
	"fmt"
	"log"
	"os/user"
)

// user包允许通过名称或ID查询用户帐户
func main() {

	// 根据用户名查询用户
	exampleLookup()
	// 获取用户账号
	exampleCurrent()
}

func exampleLookup() {

	// 根据用户名查询用户
	u, err := user.Lookup("root")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Lookup Name:", u.Name)

	// 根据用户ID查询用户
	ui, err :=  user.LookupId("501")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Lookup Uid:", ui.Uid)

	// 根据组ID查询组
	gi, err := user.LookupGroupId("20")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println( "Lookup Gid:", gi.Gid)

	// 根据组名称查询组
	g, err := user.LookupGroup("staff")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Lookup GroupName:", g.Name)
}

func exampleCurrent() {

	// 返回当前的用户帐户
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// 用户名
	fmt.Println("Name:", u.Name)
	// 登陆名
	fmt.Println("Username:", u.Username)
	// 用户文件夹路径
	fmt.Println("HomeDir:", u.HomeDir)
	// 用户ID
	fmt.Println("Uid:", u.Uid)
	// 用户组ID
	fmt.Println("Gid:", u.Gid)
	// 用户所属组ID列表
	fmt.Println(u.GroupIds())
}
