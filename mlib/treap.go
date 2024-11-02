package mlib

import "math/rand"

type Node struct {
	ch       [2]*Node
	val      int
	priority int
}

func (o *Node) cmp(v int) int {
	if o.val > v {
		return 0
	} else if o.val < v {
		return 1
	}
	return -1
}

func (o *Node) rotate(d int) *Node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
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
		return &Node{val: v, priority: rand.Int()}
	}
	d := o.cmp(v)
	o.ch[d] = t._put(o.ch[d], v)
	if o.ch[d].priority > o.priority {
		o = o.rotate(d ^ 1)
	}
	return o
}

func (t *Treap) Delete(v int) {
	t.root = t._delete(t.root, v)
}

func (t *Treap) _delete(o *Node, v int) *Node {
	if d := o.cmp(v); d >= 0 {
		o.ch[d] = t._delete(o.ch[d], v)
		return o
	}
	if o.ch[0] == nil {
		return o.ch[1]
	}
	if o.ch[1] == nil {
		return o.ch[0]
	}
	d := 0
	if o.ch[0].priority > o.ch[1].priority {
		d = 1
	}
	o = o.rotate(d)
	o.ch[d] = t._delete(o.ch[d], v)
	return o
}

func (t *Treap) Ceil(v int) (res *Node) {
	cur := t.root
	for cur != nil {
		d := cur.cmp(v)
		if d == 0 {
			res = cur
			cur = cur.ch[0]
		} else if d == 1 {
			cur = cur.ch[1]
		} else {
			return cur
		}
	}
	return
}

func (t *Treap) Floor(v int) (res *Node) {
	cur := t.root
	for cur != nil {
		d := cur.cmp(v)
		if d == 0 {
			cur = cur.ch[0]
		} else if d == 1 {
			res = cur
			cur = cur.ch[1]
		} else {
			return cur
		}
	}
	return
}
