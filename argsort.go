package ggwo

import "sort"

type argsort struct {
	s    []float64
	inds []int
}

func (a argsort) Len() int {
	return len(a.s)
}

func (a argsort) Less(i, j int) bool {
	return a.s[i] < a.s[j]
}

func (a argsort) Swap(i, j int) {
	a.s[i], a.s[j] = a.s[j], a.s[i]
	a.inds[i], a.inds[j] = a.inds[j], a.inds[i]
}

func argsortFunc(dst []float64) []int {
	n := len(dst)
	inds := make([]int, n)
	for i := range dst {
		inds[i] = i
	}

	a := argsort{s: dst, inds: inds}
	sort.Stable(a)

	return inds
}
