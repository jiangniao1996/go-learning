package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantBit(t *testing.T) {
	a := 1 //0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	//t.Log(Writable)
}
