package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoing(t *testing.T) {
	a := 1
	aPrt := &a
	// aPrt += 1
	t.Log(a, aPrt)
	t.Logf("%T %T", a, aPrt)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
