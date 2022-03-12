package my_string

import (
	"testing"
)

func TestStringToRune(t *testing.T) {
	s := "客"
	t.Log(len(s))
	c := []rune(s)
	t.Log(len(c))

	s = "中华人民共和国"
	for _, d := range s {
		t.Logf("%[1]c %[1]x %[1]d", d)
	}
}
