package mlib

import (
	"testing"
)

func Test_treap(t *testing.T) {
	tr := &Treap{}
	tr.Put(1)
	tr.Put(4)
	tr.Put(3)
	n := tr.Ceil(1)
	if n == nil || n.val != 1 {
		t.Fatal("want 1 but", n)
	}
	n = tr.Ceil(2)
	if n == nil || n.val != 3 {
		t.Fatal("want 3 but", n)
	}
	n = tr.Ceil(3)
	if n == nil || n.val != 3 {
		t.Fatal("want 3 but", n)
	}
	n = tr.Ceil(4)
	if n == nil || n.val != 4 {
		t.Fatal("want 4 but", n)
	}
	n = tr.Ceil(5)
	if n != nil {
		t.Fatal("want nil but", n)
	}
	tr.Delete(3)
	n = tr.Ceil(3)
	if n == nil || n.val != 4 {
		t.Fatal("want 4 but", n)
	}
}
