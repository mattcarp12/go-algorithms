package mergesort_test

import (
	"testing"

	"github.com/mattcarp12/go-algorithms/sorting/mergesort"
)

func sorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func TestMerge(t *testing.T) {
	a1, a2 := []int{2}, []int{1}
	res := mergesort.Merge(a1, a2)
	if !sorted(res) {
		t.Error("Wrong!")
	}

	a1, a2 = []int{1}, []int{}
	res = mergesort.Merge(a1, a2)
	if !sorted(res) {
		t.Error("Wrong!")
	}

	a1, a2 = []int{1, 2, 3}, []int{20}
	res = mergesort.Merge(a1, a2)
	if !sorted(res) {
		t.Error("Wrong!")
	}

	a1, a2 = []int{1}, []int{21, 22, 23}
	res = mergesort.Merge(a1, a2)
	if !sorted(res) {
		t.Error("Wrong!")
	}

	a1, a2 = []int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8}
	res = mergesort.Merge(a1, a2)
	if !sorted(res) {
		t.Error("Wrong!")
	}
}

func TestMergesort(t *testing.T) {
	arr := []int{2, 1}
	res := mergesort.Mergesort(arr)
	if !sorted(res) {
		t.Error("Wrong!")
	}

	arr = []int{}
	res = mergesort.Mergesort(arr)
	if !sorted(res) {
		t.Error("Wrong!")
	}

	arr = []int{1}
	res = mergesort.Mergesort(arr)
	if !sorted(res) {
		t.Error("Wrong!")
	}

	arr = []int{9, 8, 7, 7, 6, 5, 6, 4, 5, 4, 33, 1}
	res = mergesort.Mergesort(arr)
	if !sorted(res) {
		t.Error("Wrong!")
	}
}
