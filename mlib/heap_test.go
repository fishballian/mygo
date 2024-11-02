package mlib

import (
	"math/rand"
	"sort"
	"testing"
)

func TestHeap(t *testing.T) {
	a := []int{}
	h := &Heap{}
	for i := 0; i < 1000; i++ {
		r := rand.Int() % 10
		if len(a) == 0 || r < 6 {
			v := rand.Int() % 9
			a = append(a, v)
			sort.Ints(a)
			h.Push(v)
			//t.Log("push", v, a)
		} else {
			a0 := a[0]
			a = a[1:]
			//t.Log("pop", a0, a)
			top := h.Pop()
			if a0 != top {
				t.Fatal(top, a0, h, a)
			}
		}

		if len(a) == 0 {
			if h.Len() != 0 {
				t.Fatal(h.Len(), h)
			}
		} else if h.Len() == 0 || h.Top() != a[0] {
			t.Fatal(h.Top(), a[0], h, a)
		}
	}
}

func TestHeap2(t *testing.T) {
	h := &Heap{}
	h.Push(1)
	if h.Top() != 1 {
		t.Fatal(h.Top(), h)
	}
	h.Push(2)
	if h.Top() != 1 {
		t.Fatal(h.Top(), h)
	}
	h.Pop()
	if h.Len() != 1 {
		t.Fatal(h.Len(), h)
	}
	if h.Top() != 2 {
		t.Fatal(h.Top(), h)
	}
	h.Push(0)
	if h.Top() != 0 {
		t.Fatal(h.Top(), h)
	}
	h.Pop()
	if h.Top() != 2 {
		t.Fatal(h.Top(), h)
	}
	h.Pop()
	if h.Len() != 0 {
		t.Fatal(h.Len(), h)
	}
}
