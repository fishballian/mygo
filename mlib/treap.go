package mlib

import "math/rand"

type Node struct {
	Child    [2]*Node
	Val      int
	Priority int
}

func (o *Node) cmp(v int) int {
	if o.Val > v {
		return 0
	} else if o.Val < v {
		return 1
	}
	return -1
}

func (o *Node) rotate(d int) *Node {
	x := o.Child[d^1]
	o.Child[d^1] = x.Child[d]
	x.Child[d] = o
	return x
}

type Treap struct {
	root *Node
}

func (t *Treap) Put(v int) {
	t.root = t._put(t.root, v)
}

func (t *Treap) _put(o *Node, v int) *Node {
	if o == nil {
		return &Node{Val: v, Priority: rand.Int()}
	}
	d := o.cmp(v)
	if d >= 0 {
		o.Child[d] = t._put(o.Child[d], v)
		if o.Child[d].Priority > o.Priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *Treap) Delete(v int) {
	t.root = t._delete(t.root, v)
}

func (t *Treap) _delete(o *Node, v int) *Node {
	if o == nil {
		return nil
	}
	if d := o.cmp(v); d >= 0 {
		o.Child[d] = t._delete(o.Child[d], v)
		return o
	}
	if o.Child[0] == nil {
		return o.Child[1]
	}
	if o.Child[1] == nil {
		return o.Child[0]
	}
	d := 0
	if o.Child[0].Priority > o.Child[1].Priority {
		d = 1
	}
	o = o.rotate(d)
	o.Child[d] = t._delete(o.Child[d], v)
	return o
}

func (t *Treap) Ceil(v int) (res *Node) {
	return t._find(v, 0)
}

func (t *Treap) Floor(v int) (res *Node) {
	return t._find(v, 1)
}

func (t *Treap) _find(v, cmp int) (res *Node) {
	cur := t.root
	for cur != nil {
		d := cur.cmp(v)
		if d == -1 {
			return cur
		}
		if d == cmp {
			res = cur
		}
		cur = cur.Child[d]
	}
	return
}
