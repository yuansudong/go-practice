package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string  `xorm:"x_name" json:"j_name"`
	Age    int     `xorm:"x_age" json:"j_age"`
	Height float32 `xorm:"x_height" json:"j_height"`
}

func (s *Student) GetName_V1() string {
	return s.Name
}

func (s *Student) GetAge_V1() int {
	return s.Age
}
func (s *Student) GetHeight_V1() float32 {
	return s.Height
}

func (s Student) GetName_V2() string {
	return s.Name
}
func (s Student) GetAge_V2() int {
	return s.Age
}
func (s Student) GetHeight_V2() float32 {
	return s.Height
}

func main() {
	mStu := &Student{
		Age:    12,
		Name:   "Steven",
		Height: 20.69,
	}
	ModifyValue(mStu)
	fmt.Println(mStu)
}

func RangeType(i interface{}) {
	v := reflect.Indirect(reflect.ValueOf(i))
	t := v.Type()

	for index := 0; index < t.NumField(); index++ {
		f := t.Field(index)
		fmt.Println(f.Index, f.Anonymous, f.Name, f.Offset, f.PkgPath, f.Tag)
	}
}

func RangeMethod(i interface{}) {
	t := reflect.TypeOf(i)
	for index := 0; index < t.NumMethod(); index++ {
		f := t.Method(index)
		fmt.Println(f.Index, f.Name, f.Type, f.PkgPath)
	}
}

// RangeValue 遍历值
func RangeValue(i interface{}) {
	t := reflect.Indirect(reflect.ValueOf(i))
	for index := 0; index < t.NumField(); index++ {
		f := t.Field(index)
		fmt.Println(f.Type().Name(), f.Interface())
	}
}

func ModifyValue(i interface{}) {
	t := reflect.Indirect(reflect.ValueOf(i))
	t.FieldByName("Name").SetString("hex")
	t.FieldByName("Age").SetInt(28)
	t.FieldByName("Height").SetFloat(28.8)
}

func RangeCallFunc(i interface{}) {
	t := reflect.ValueOf(i)
	for index := 0; index < t.NumMethod(); index++ {
		f := t.Method(index)
		fmt.Println(f.Call([]reflect.Value{})[0].Interface())
	}
}

func NewStruct(arr []string) *Student {
	reflect.New(reflect.TypeOf(Student{}))
	return nil
}
