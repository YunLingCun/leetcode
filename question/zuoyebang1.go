package question

func MergeArray(a []int, b []int) []int {
	var c = make([]int, 0, len(a)+len(b))
	for len(a) > 0 && len(b) > 0 {
		if a[0] < b[0] {
			c = append(c, a[0])
			a = a[1:]
		} else {
			c = append(c, b[0])
			b = b[1:]
		}
	}
	c = append(c, a...)
	c = append(c, b...)
	return c
}
