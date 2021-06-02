package string

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var data []string = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "night", "ten",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "night", "ten",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "night", "ten",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "night", "ten",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "night", "ten",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "night", "ten",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "night", "ten",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "night", "ten",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "night", "ten",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "night", "ten",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "night", "ten",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "night", "ten",
}

func BenchmarkPlusString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PlusString()
	}
}
func PlusString() string {
	s := ""
	for i := 0; i < len(data); i++ {
		s += data[i]
	}
	return s
}

func BenchmarkBuilderString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BuilderString()
	}
}
func BuilderString() string {
	b := new(strings.Builder)
	for i := 0; i < len(data); i++ {
		b.WriteString(data[i])
	}
	return b.String()
}

func BenchmarkSprintString(b *testing.B) {
	iData := []interface{}{}
	for _, item := range data {
		iData = append(iData, item)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SprintString(iData)
	}
}
func SprintString(idata []interface{}) string {

	return fmt.Sprint(idata...)
}

func BenchmarkJoinString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JoinString()
	}
}
func JoinString() string {
	return strings.Join(data, "")
}

func BenchmarkBufferString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BufferString()
	}
}
func BufferString() string {
	builder := bytes.NewBuffer(make([]byte, 4096))
	for i := 0; i < len(data); i++ {
		builder.WriteString(data[i])
	}
	return builder.String()
}

func TestReplaceXXX(t *testing.T) {
	fmt.Println(strings.Replace("aaa", "a", "b", 1))
	fmt.Println(strings.ReplaceAll("aaa", "a", "b"))
}

func TestSplitXXX(t *testing.T) {
	fmt.Println(strings.Split("a,b,c", ","))
	fmt.Println(strings.SplitN("a,b,c", ",", 1))
	fmt.Println(strings.SplitAfter("a,b,c,", ","))
	fmt.Println(strings.SplitAfterN("a,b,c,", ",", 1))
}

func TestHasPrefixAndSuffix(t *testing.T) {

	if strings.HasPrefix("ac", "a") {
		fmt.Println("ac有前缀a")
	}
	if strings.HasSuffix("ac", "d") {
		fmt.Println("ac没有后缀d")
	}

}

func TestToXXX(t *testing.T) {
	//str
	str := "aTbacG hello"
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToTitle(str))
	fmt.Println(strings.Title(str))
}

func TestMap(t *testing.T) {

	fmt.Println(strings.Map(func(r rune) rune {
		if r == 'o' {
			return 'e'
		}
		return r
	}, "hello,world"))
}

func TestTrimXXX(t *testing.T) {
	fmt.Println(strings.Trim("\n\n\n hello,world \n\n\n", "\n"))
	fmt.Println(strings.TrimFunc("\n\n\n hello,world \n\n\n", func(r rune) bool {
		return r == '\n' || r == ' '
	}))
	fmt.Println(strings.TrimLeft("\n\n\n hello,world \n\n\n", "\n"))
	fmt.Println(strings.TrimLeftFunc("\n\n\n hello,world \n\n\n", func(r rune) bool {
		return r == '\n' || r == ' '
	}))
	fmt.Println(strings.TrimRight("\n\n\n hello,world \n\n\n", "\n"))
	fmt.Println(strings.TrimRightFunc("\n\n\n hello,world \n\n\n", func(r rune) bool {
		return r == '\n' || r == ' '
	}))
	fmt.Println(strings.TrimSpace(" \t\n Hello, world \n\t\r\n"))
}

func TestContains(t *testing.T) {
	a := "adsjgdysafdb"
	b := "b"
	if strings.Contains(a, b) {
		fmt.Println("a是包含b的")
	}
	if strings.ContainsAny(a, "☺") {
		fmt.Println("a是包含☺的")
	}
	if strings.ContainsRune(a, 'a') {
		fmt.Println("a是包含a的")
	}

}

func TestField(t *testing.T) {
	txt := `
	dasdsa			dskajdbksa 			

	dsakjdbas

	hello 
	world
	`
	fmt.Println(strings.Join(strings.Fields(txt), " "))
}
