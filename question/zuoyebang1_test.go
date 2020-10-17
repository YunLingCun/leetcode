package question

import "testing"

func TestMergeArray(t *testing.T) {
	var a = []int{1,2,3,4,5}
	var b = []int{1,2,3,4,5}
	t.Logf("1: %+v\n",MergeArray(a,b))
	b = append(b, 6,7)
	t.Logf("2: %+v\n",MergeArray(a,b))
	a = []int{}
	t.Logf("3: %+v\n",MergeArray(a,b))
	b=[]int{}
	t.Logf("4: %+v\n",MergeArray(a,b))
}

