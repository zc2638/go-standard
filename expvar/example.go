package main

import (
	"expvar"
	"fmt"
)

// 公共变量，例如可用做服务的操作计数器
func main() {

	// 创建float类型的公共变量
	expvar.NewFloat("testFloat")
	// 创建string类型的公共变量
	expvar.NewString("testStr")
	// 创建int类型的公共变量
	i := expvar.NewInt("testInt")
	// 设置变量值
	i.Set(16)
	// 原有变量加上一个变量
	i.Add(15)
	// 返回该变量类型值
	i.Value()
	// 返回该变量的string类型值
	i.String()

	// 创建map类型的公共变量
	m := expvar.NewMap("testMap")
	// 重置为一个初始的空map
	m.Init()
	// 设置map键值对
	m.Set("k1", i)
	// 为map指定键的值加上delta值，键值必须为int类型
	m.Add("k1", 5)
	// 为map指定键的值加上delta值，键值必须为float类型
	m.AddFloat("k2", 5.2)
	// 返回map字符串
	m.String()
	// 获取map指定键的expvar.Val值
	v := m.Get("k1")
	v.String()
	// 根据键删除指定键值对
	m.Delete("k2")
	// 为map下每组键值对调用指定方法
	m.Do(func(kv expvar.KeyValue) {
		fmt.Println(kv.Key, kv.Value.String())
	})

	// 获取公共变量Var值
	expvar.Get("testInt")
	// 为每组公共变量调用指定方法
	expvar.Do(func(kv expvar.KeyValue) {
		fmt.Println(kv.Key, kv.Value.String())
	})

	// 返回一个http.Handler
	expvar.Handler()
}

func init() {

	// 使用expvar.Fun结构返回值
	pbInt := func() interface{} {
		return 12
	}
	// 声明一个导出变量，必须在init函数里调用，如果变量名存在则调用log.Panic
	expvar.Publish("pbInt", expvar.Func(pbInt))
}
