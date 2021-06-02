package main

import "fmt"

const (
	_  = iota // 用于占iota的0位.
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(KB, MB)
}
