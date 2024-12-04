package mlib

type DQueue struct {
	a []int
	b []int
}

func (q *DQueue) Len() int {
	return len(q.a) + len(q.b)
}

func (q *DQueue) PushBack(v int) {
	q.b = append(q.b, v)
}

func (q *DQueue) PopBack() int {
	if len(q.b) > 0 {
		v := q.b[len(q.b)-1]
		q.b = q.b[:len(q.b)-1]
		return v
	}
	v := q.a[0]
	q.a = q.a[1:]
	return v
}

func (q *DQueue) PushFront(v int) {
	q.a = append(q.a, v)
}

func (q *DQueue) PopFront() int {
	if len(q.a) > 0 {
		v := q.a[len(q.a)-1]
		q.a = q.a[:len(q.a)-1]
		return v
	}
	v := q.b[0]
	q.b = q.b[1:]
	return v
}

func (q *DQueue) IsEmpty() bool {
	return len(q.a) == 0
}
