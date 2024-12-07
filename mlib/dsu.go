package mlib

type Dsu struct {
	parent []int
}

func NewDsu(n int) *Dsu {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &Dsu{parent: parent}
}

func (d *Dsu) FindParent(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.FindParent(d.parent[x])
	}
	return d.parent[x]
}

func (d *Dsu) Merge(x, y int) {
	px, py := d.FindParent(x), d.FindParent(y)
	if px == py {
		return
	}
	d.parent[px] = py
}
