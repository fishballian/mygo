package mlib

import (
	"math"
	"math/rand"
)

const MaxL = 32
const S = 0xFFFF
const P = 4
const PS = S / P
const INVALID = math.MaxInt

func randomLevel() int {
	lv := 1
	for float64(rand.Int()&S) < PS {
		lv++
	}
	return min(MaxL, lv)
}

type SkipList struct {
	level  int
	length int
	head   *SkipListNode
	tail   *SkipListNode
}

type SkipListNode struct {
	key     int
	val     int
	level   int
	forward []*SkipListNode
}

func NewSkipList() *SkipList {
	tail := &SkipListNode{key: INVALID}
	head := &SkipListNode{key: INVALID, level: MaxL, forward: make([]*SkipListNode, MaxL+1)}
	for i := 0; i <= MaxL; i++ {
		head.forward[i] = tail
	}
	return &SkipList{head: head, tail: tail}
}

func (s *SkipList) Find(key int) int {
	p := s.head
	for i := s.level; i >= 0; i-- {
		for p.forward[i].key < key {
			p = p.forward[i]
		}
	}
	p = p.forward[0]
	if p.key == key {
		return p.val
	}
	return s.tail.val
}

func (s *SkipList) Insert(key int, val int) {
	update := make([]*SkipListNode, MaxL+1)
	p := s.head
	for i := s.level; i >= 0; i-- {
		for p.forward[i].key < key {
			p = p.forward[i]
		}
		update[i] = p
	}
	p = p.forward[0]

	if p.key == key {
		p.val = val
		return
	}

	lv := randomLevel()
	if lv > s.level {
		s.level++
		lv = s.level
		update[lv] = s.head
	}

	newNode := &SkipListNode{key: key, val: val, level: lv}
	newNode.forward = make([]*SkipListNode, lv+1)
	for i := lv; i >= 0; i-- {
		p = update[i]
		newNode.forward[i] = p.forward[i]
		p.forward[i] = newNode
	}
	s.length++
}

func (s *SkipList) Delete(key int) bool {
	update := make([]*SkipListNode, MaxL+1)
	p := s.head
	for i := s.level; i >= 0; i-- {
		for p.forward[i].key < key {
			p = p.forward[i]
		}
		update[i] = p
	}
	p = p.forward[0]
	if p.key != key {
		return false
	}
	for i := 0; i <= s.level; i++ {
		if update[i].forward[i] != p {
			break
		}
		update[i].forward[i] = p.forward[i]
	}

	for s.level > 0 && s.head.forward[s.level] == s.tail {
		s.level--
	}
	s.length--
	return true
}
