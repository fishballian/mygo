package mlib

type FwTree struct {
	a []int
}

func NewFwTree(c int) *FwTree {
	return &FwTree{a: make([]int, c)}
}

func (f *FwTree) Add(p, add int) {
	for p < len(f.a) {
		f.a[p] += add
		p += p & -p
	}
}

func (f *FwTree) Get(p int) int {
	s := 0
	for p > 0 {
		s += f.a[p]
		p -= p & -p
	}
	return s
}
