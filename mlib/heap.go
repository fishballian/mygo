package mlib

type Heap struct {
	data []int
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Less(i, j int) bool {
	return h.data[i] < h.data[j]
}

func (h *Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

/*
func (h *Heap) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}

func (h *Heap) Pop() interface{} {
	l := h.Len()
	v := h.data[l-1]
	h.data = h.data[:l-1]
	return v
}
*/

func (h *Heap) Push(x int) {
	t := h.Len()
	h.data = append(h.data, x)
	for t > 0 {
		p := (t - 1) / 2
		if h.Less(t, p) {
			h.Swap(t, p)
			t = p
		} else {
			break
		}
	}
}

func (h *Heap) Top() int {
	return h.data[0]
}

func (h *Heap) Pop() int {
	v := h.data[0]
	l := h.Len()
	h.Swap(0, l-1)
	h.data = h.data[:l-1]
	l--
	t := 0
	left := 1
	for left < l {
		if left+1 < l && h.Less(left+1, left) {
			left += 1
		}
		if h.Less(left, t) {
			h.Swap(left, t)
			t = left
			left = t*2 + 1
		} else {
			break
		}
	}
	return v
}
