package mlib

func CallM(m int) (res []int) {
	if m <= 1 || m >= 100 {
		return res
	}
	size := 100
	c := size
	kick := make([]bool, size)
	curPos := 0
	curNum := 1
	for ; c >= m; curPos = (curPos + 1) % size {
		if kick[curPos] {
			continue
		}
		if curNum == m {
			kick[curPos] = true
			c--
			curNum = 0
		}
		curNum++
	}
	for k, v := range kick {
		if v == false {
			res = append(res, k+1)
		}
	}
	return res
}
