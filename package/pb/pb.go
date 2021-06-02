package pb

import "fmt"

const (
	a  string = "hello"
	HA string = "HA"
)

var (
	b string = GetStringB()
)

func GetStringB() string {
	fmt.Println("初始化pb.b")
	return a + "B"
}

func init() {
	fmt.Println("初始化pb.init函数")
}
