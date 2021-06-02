package main

import "fmt"

func main() {

	i8 := int8(8)
	i16 := int16(16)
	i32 := int32(32)
	i64 := int64(64)
	f32 := float32(16.2)
	ui8 := uint8(i8)
	ui16 := uint16(i16)
	ui32 := uint32(i32)
	ui64 := uint64(i64)
	f64 := float32(f32)
	fmt.Println(ui8, ui16, ui32, ui64, f64)
}
