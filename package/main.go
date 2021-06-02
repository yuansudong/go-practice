package main

import (
	"fmt"

	_ "github.com/yuansudong/go-practice/package/pa"
)

const (
	a string = "hello"
)

var (
	b string = GetStringB()
)

func GetStringB() string {
	fmt.Println("初始化main.b")
	return a + "B"
}

func init() {
	fmt.Println("初始化main.init")
}

func main() {
	fmt.Println("执行main:", b)
}

// 常量 --> 变量 --> init函数 --> main函数
