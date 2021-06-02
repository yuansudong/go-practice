package string

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	builder := bytes.NewBuffer(make([]byte, 1))
	for i := 0; i < 10000; i++ {
		builder.WriteString(strconv.FormatInt(int64(i), 10))
	}
	fmt.Println(builder.String())
}
