package mytypes

import "strings"

type MyInt int
type MyString string
type MyBuilder struct {
	Contents strings.Builder
	}
type MyUpper struct {
	Content strings.Builder
	}

func (mu MyUpper) ToUpper() string {
	return strings.ToUpper(mu.Content.String())
}	

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (s MyString) Len() int {
	return len(s)
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

func (x *MyInt) Double() {
	*x *= 2
}
