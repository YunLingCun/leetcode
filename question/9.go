package question

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var vl = make([]int, 0, 0)
	for x > 0 {
		vl = append(vl, x%10)
		x /= 10
	}
	var length = len(vl)
	for idx := range vl[:length/2] {
		if vl[idx] != vl[length-idx-1] {
			return false
		}
	}
	return true
}

func IsPalindrome2(x int) bool {
	if x < 0 {return false}
	if x < 10 {return true}
	if x % 10 == 0 {return false}
	var r = 0
	for  {
		e := x % 10
		x = x/10
		r = r * 10 + e
		if x < r {
			return false
		}
		if  x/10 == r || x == r {
			return true
		}
	}
}