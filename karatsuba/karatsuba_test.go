package karatsuba_test

import (
	"testing"

	"github.com/mattcarp12/go-algorithms/karatsuba"
)

func TestStrCmp(t *testing.T) {
	t1 := "1234" < "123"
	if t1 {
		t.Error()
	}

	t2 := "1234" < "1235"
	if !t2 {
		t.Error()
	}

	t3 := "99" <= "99"
	if !t3 {
		t.Error()
	}

	t4 := karatsuba.YugeNum("9") < "10"
	if t4 {
		t.Error()
	}
}

func TestYugeNumToIntSlice(t *testing.T) {
	s := karatsuba.YugeNum("1234")
	is := s.ToIntSlice()
	if is[0] != 1 || is[1] != 2 || is[2] != 3 || is[3] != 4 {
		t.Error("Wrong!")
	}
}

func TestIntSliceToString(t *testing.T) {
	is := []int{1, 2, 3, 4}
	s := karatsuba.IntSliceToString(is)
	if s != "1234" {
		t.Error("Wrong!")
	}
}

func TestAddTwoPositive(t *testing.T) {
	y1 := karatsuba.YugeNum("123")
	y2 := karatsuba.YugeNum("123")
	y3 := karatsuba.Add(y1, y2)
	if y3 != "246" {
		t.Error("Wrong!")
	}
	t.Logf("%+v", y3)

	y1 = karatsuba.YugeNum("9999")
	y2 = karatsuba.YugeNum("99")
	y3 = karatsuba.Add(y1, y2)
	if y3 != "10098" {
		t.Error("Wrong!")
	}
	t.Logf("%+v", y3)
}

func TestE(t *testing.T) {
	y := karatsuba.YugeNum("1")
	y = y.E(2)
	if y != "100" {
		t.Error("Wrong!")
	}
}

func TestSplit(t *testing.T) {
	y := karatsuba.YugeNum("10")
	one, zero := y.Split(1)
	if one != "1" || zero != "0" {
		t.Error("Wrong!")
	}

	y = karatsuba.YugeNum("12345")
	top, bot := y.Split(3)
	if top != "123" || bot != "45" {
		t.Error("Wrong!")
	}
}

func TestMultiplyDigit(t *testing.T) {
	y := karatsuba.YugeNum("10")
	z := y.MultiplyDigit(2)
	if z != "20" {
		t.Error("Wrong!")
	}

	y = karatsuba.YugeNum("999")
	z = y.MultiplyDigit(9)
	if z != "8991" {
		t.Error("Wrong!")
	}
}

func TestSubtract(t *testing.T) {
	y1 := karatsuba.YugeNum("123")
	y2 := karatsuba.YugeNum("23")
	z := karatsuba.Sub(y1, y2)
	if z != "100" {
		t.Error("Wrong!")
	}

	y1 = karatsuba.YugeNum("100")
	y2 = karatsuba.YugeNum("55")
	z = karatsuba.Sub(y1, y2)
	if z != "45" {
		t.Error("Wrong!")
	}

}

func TestTrim(t *testing.T) {
	y := karatsuba.YugeNum("00123")
	y = y.Trim()
	if y != "123" {
		t.Error()
	}

	y = karatsuba.YugeNum("000")
	y = y.Trim()
	if y != "0" {
		t.Error()
	}
}

func TestKaratsuba(t *testing.T) {
	y1 := karatsuba.YugeNum("100")
	y2 := karatsuba.YugeNum("13")
	z := karatsuba.Karatsuba(y1, y2)
	t.Log(z)
	if z != "1300" {
		t.Error("Wrong!")
	}

	y1 = karatsuba.YugeNum("123")
	y2 = karatsuba.YugeNum("123")
	z = karatsuba.Karatsuba(y1, y2)
	t.Log(z)
	if z != "15129" {
		t.Error("Wrong!")
	}

	y1 = karatsuba.YugeNum("9999")
	y2 = karatsuba.YugeNum("9999")
	z = karatsuba.Karatsuba(y1, y2)
	t.Log(z)
	if z != "99980001" {
		t.Error("Wrong!")
	}

	y1 = karatsuba.YugeNum("999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999")
	y2 = karatsuba.YugeNum("999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999")
	z = karatsuba.Karatsuba(y1, y2)
	t.Log(z)
	if z != "999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999998000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001" {
		t.Error("Wrong!")
	}
}
