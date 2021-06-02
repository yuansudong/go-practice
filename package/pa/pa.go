package pa

import (
	"fmt"

	_ "github.com/yuansudong/go-practice/package/pb"
)

const (
	a string = "hello"
)

var (
	b string = GetStringB()
)

func GetStringB() string {
	fmt.Println("初始化pa.b")
	return a + "B"
}

func init() {
	fmt.Println("初始化pa.init函数")
}
