package mlib

import (
	"math/rand"
	"testing"
)

var s *SkipList
var size = 1000000

func TestMain(m *testing.M) {
	s = NewSkipList()
	for i := 0; i < size; i++ {
		v := rand.Int()%size + 1
		s.Insert(v, v)
	}
	m.Run()
}

func TestSkipList(t *testing.T) {
	s := NewSkipList()
	s.Insert(1, 1)
	v := s.Find(1)
	if v != 1 {
		t.Fatal("want 1 but", v)
	}
	s.Delete(1)
	v = s.Find(1)
	if v == 1 {
		t.Fatal("1 not deleted")
	}
	for i := 0; i < 1000000; i++ {
		v = rand.Int()%1000000 + 1
		s.Insert(v, v)
		if v != s.Find(v) {
			t.Fatal("want v", v, "but got", s.Find(v))
		}
	}
}

func TestRandomLevel(t *testing.T) {
	m := make([]int, MaxL+1)
	for i := 0; i < 1000000; i++ {
		lv := randomLevel()
		m[lv]++
	}
	t.Log(m)
}

func BenchmarkSkipListFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := rand.Int()%size + 1
		s.Find(v)
	}
}
