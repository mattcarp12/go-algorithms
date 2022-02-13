package karatsuba

type YugeNum string

func (y YugeNum) ToIntSlice() []int {
	var val []int
	for _, digit := range y {
		val = append(val, int(digit-'0'))
	}
	return val
}

func IntSliceToString(is []int) string {
	var barr []byte
	for _, digit := range is {
		barr = append(barr, byte(digit+'0'))
	}
	return string(barr)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func zeroPad(is []int, m int) []int {
	zs := make([]int, m)
	return append(zs, is...)
}

func reverseIntSlice(is []int) []int {
	var res []int
	for i := len(is) - 1; i >= 0; i-- {
		res = append(res, is[i])
	}
	return res
}

func (y1 YugeNum) lt(y2 YugeNum) bool {
	l1 := len(y1)
	l2 := len(y2)
	if l1 < l2 {
		for i := 0; i < l2-l1; i++ {
			y1 = "0" + y1
		}
	} else if l2 < l1 {
		for i := 0; i < l1-l2; i++ {
			y2 = "0" + y2
		}
	}
	return y1 < y2
}

func Add(y1, y2 YugeNum) YugeNum {

	// Add two positive integer strings
	var big, lil []int
	if y1.lt(y2) {
		lil = y1.ToIntSlice()
		big = y2.ToIntSlice()
	} else {
		lil = y2.ToIntSlice()
		big = y1.ToIntSlice()
	}

	// zero pad the smaller integer to make computation easier
	lil = zeroPad(lil, len(big)-len(lil))

	// Put result into new integer slice backwards, then reverse
	var sum []int
	carry := 0
	for i := len(big) - 1; i >= 0; i-- {
		res := big[i] + lil[i] + carry
		if res >= 10 {
			carry = 1
			res = res - 10
		} else {
			carry = 0
		}
		sum = append(sum, res)
	}
	if carry == 1 {
		sum = append(sum, carry)
	}

	// Reverse sum
	sum = reverseIntSlice(sum)

	// Return resulting YugeNum
	return YugeNum(IntSliceToString(sum))
}

func Sub(y1, y2 YugeNum) YugeNum {
	// For the purposes of this algorithm, we can
	// assume the first number is always greater than
	// the second, and the result will be positive

	is1 := y1.ToIntSlice()
	is2 := y2.ToIntSlice()

	// Zero pad the second number so they're the same length
	is2 = zeroPad(is2, len(is1)-len(is2))

	// Loop through, performing subtraction
	carry := 0
	diff := []int{}
	for i := len(is1) - 1; i >= 0; i-- {
		res := is1[i] - is2[i] - carry
		if res < 0 {
			carry = 1
			res = res + 10
		} else {
			carry = 0
		}
		diff = append(diff, res)
	}
	if carry != 0 {
		panic("Shouldn't happen")
	}

	// Reverse result, return YugeNum
	diff = reverseIntSlice(diff)
	return YugeNum(IntSliceToString(diff)).Trim()
}

func (y YugeNum) MultiplyDigit(d int) YugeNum {
	if d > 9 {
		panic("not a digit")
	}

	// Make string an int slice
	is := y.ToIntSlice()

	// Multiply each digit, handling carry
	res := []int{}
	carry := 0
	for i := len(is) - 1; i >= 0; i-- {
		mul := is[i]*d + carry
		res = append(res, mul%10)
		carry = mul / 10
	}
	if carry > 0 {
		res = append(res, carry%10)
		carry = carry / 10
		if carry > 0 {
			res = append(res, carry)
		}
	}

	// Reverse result, return YugeNum
	res = reverseIntSlice(res)
	return YugeNum(IntSliceToString(res))
}

func (y YugeNum) Split(m int) (YugeNum, YugeNum) {
	// split from right
	n := len(y) - m
	return y[0:n].Trim(), y[n:].Trim()
}

func (y YugeNum) E(pow int) YugeNum {
	for i := 0; i < pow; i++ {
		y += "0"
	}
	return y
}

func (y YugeNum) Trim() YugeNum {
	i := 0
	for ; i < len(y)-1; i++ {
		if y[i] != byte('0') {
			break
		}
	}
	return y[i:]
}

func Karatsuba(y1, y2 YugeNum) YugeNum {
	if y1.lt(YugeNum("10")) {
		return y2.MultiplyDigit(int(y1[0] - '0'))
	} else if y2.lt(YugeNum("10")) {
		return y1.MultiplyDigit(int(y2[0] - '0'))
	}

	// Split integers at half length of smaller of the two
	m := min(len(y1), len(y2))
	m2 := m / 2

	high1, low1 := y1.Split(m2)
	high2, low2 := y2.Split(m2)

	// Recursive calls for intermediate multiplications
	z0 := Karatsuba(low1, low2)
	z2 := Karatsuba(high1, high2)
	z1 := Sub(Karatsuba(Add(low1, high1), Add(low2, high2)), Add(z0, z2))

	return Add(Add(z2.E(m2*2), z1.E(m2)), z0)
}
