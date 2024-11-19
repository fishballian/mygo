package mlib

import (
	"math"
	"math/rand"
	"testing"
)

func Test_treap(t *testing.T) {
	for i := 0; i < 10000; i++ {
		tree := &Treap{}
		m := make(map[int]bool)
		cmd := [][]int{}
		for j := 0; j < 20; j++ {
			v := rand.Int() % 10
			if rand.Int()%2 == 0 {
				tree.Put(v)
				m[v] = true
				cmd = append(cmd, []int{0, v})
			} else {
				tree.Delete(v)
				delete(m, v)
				cmd = append(cmd, []int{1, v})
			}
			v2 := rand.Int() % 10
			c, f := math.MaxInt, math.MinInt
			for k := range m {
				if k >= v2 {
					c = min(c, k)
				}
				if k <= v2 {
					f = max(f, k)
				}
			}
			c2, f2 := math.MaxInt, math.MinInt
			if n := tree.Ceil(v2); n != nil {
				c2 = n.val
			}
			if n := tree.Floor(v2); n != nil {
				f2 = n.val
			}
			if c != c2 {
				t.Fatal("ceil", c, c2, cmd, v2)
			}
			if f != f2 {
				t.Fatal("floor", f, f2, cmd, v2)
			}
		}

	}
}
