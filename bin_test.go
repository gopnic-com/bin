package bin

import (
	"testing"
)

type ty1 struct {
	A, B string
	C, D int64
}

type ty2 struct {
	A string
	T ty1
}

func Test(t *testing.T) {
	t1 := &ty1{"a", "b", 1, 2}
	b1, err := Serialize(t1)
	if err != nil {
		t.Error(err.Error())
	}
	if len(b1) != 56 {
		t.Error("invalid b1 length, expected 56")
	}
	t2 := new(ty1)
	err = Unserialize(b1, t2)
	if err != nil {
		t.Error("Unserialize error: " + err.Error())
	}
	if t2.A != "a" {
		t.Error("invalid unserialized data, expected a")
	}
	if t2.B != "b" {
		t.Error("invalid unserialized data, expected b")
	}
	if t2.C != 1 {
		t.Error("invalid unserialized data, expected 1")
	}
	if t2.D != 2 {
		t.Error("invalid unserialized data, expected 2")
	}
	t3 := &ty2{
		A: "a",
		T: ty1{"a", "b", 1, 2},
	}
	b2, err := Serialize(t3)
	if err != nil {
		t.Error("ty2 serialize failed: " + err.Error())
	}
	if len(b2) != 92 {
		t.Error("invalid b2 len, expected 92")
	}
	t4 := new(ty2)
	err = Unserialize(b2, t4)
	if err != nil {
		t.Error("t4 unserialize failed: " + err.Error())
	}
	if t4.A != "a" {
		t.Error("invalid unserialized data, expected a")
	}
	if t4.T.A != "a" {
		t.Error("invalid unserialized data, expected a")
	}
	if t4.T.B != "b" {
		t.Error("invalid unserialized data, expected b")
	}
	if t4.T.C != 1 {
		t.Error("invalid unserialized data, expected 1")
	}
	if t4.T.D != 2 {
		t.Error("invalid unserialized data, expected 2")
	}
}
