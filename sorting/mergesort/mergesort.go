package mergesort

func Mergesort(arr []int) []int {
	// Base case
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	// Split into two halves
	m := len(arr) / 2
	a1, a2 := arr[:m], arr[m:]

	// Recursive step
	s1, s2 := Mergesort(a1), Mergesort(a2)

	// Merge sorted arrays
	return Merge(s1, s2)
}

func Merge(s1, s2 []int) []int {
	l1, l2 := len(s1), len(s2)

	var res []int
	var p1, p2 int
	for i := 0; i < l1+l2; i++ {
		if p1 == l1 {
			res = append(res, s2[p2])
			p2++
		} else if p2 == l2 {
			res = append(res, s1[p1])
			p1++
		} else {
			if s1[p1] < s2[p2] {
				res = append(res, s1[p1])
				p1++
			} else {
				res = append(res, s2[p2])
				p2++
			}
		}
	}
	return res
}
