package mytypes

import "strings"

type MyInt int
type MyString string
type MyBuilder strings.Builder

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (s MyString) Len() int {
	return len(s)
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}
