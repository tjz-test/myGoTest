package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

//func getUserInfo() (string, int) {
//	return "123", 12
//}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to yanzi travel")
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	// fmt.Println("asfd")
	// var a = "aaa"
	// fmt.Printf("%v", a)
	// a := 10
	// fmt.Printf("%v", a)
	//var _, age = getUserInfo()
	//fmt.Println(age)
	//常量定义 ：定义后必须赋值，常量值不可改变
	// const pi = 3.1415926
	// fmt.Println(pi)

	// const a, b, c, d = 1, 2, "d", 3
	// fmt.Println(a, b, c, d)

	// const (
	// 	n = 1
	// 	m
	// )
	// fmt.Println(m, n)

	//iota 常量计数器 从零开始 可插队 多个可定义在一行

	//变量区分大小写

	//基本类型 整数 浮点 布尔 字符串
	//复合 数组 切片 结构体 函数 MAP 通道（channel） 接口

	//整形 有符号 int 8.16.32.64 无符号 uint 8.16.32.64

	//浮点 float 32 4 64 8
	//m1 := 1.2
	//m2 := 2.2
	//fmt.Printf("%f", m1+m2)
	////fmt.Printf("%f", decimal.NewFormFloat(m1).add(decimal.NewFormFloat(m2)))
	//str := `adv
	//adv
	//adf
	//adf`
	//fmt.Println(str)
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
