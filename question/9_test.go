package question

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Logf("100 : %v", IsPalindrome(100))
	t.Logf("101 : %v", IsPalindrome(101))
	t.Logf("102 : %v", IsPalindrome(102))
	t.Logf("202 : %v", IsPalindrome(202))
	t.Logf("-101 : %v", IsPalindrome(-101))
	t.Logf("0 : %v", IsPalindrome(0))
	t.Logf("111 : %v", IsPalindrome(111))
	t.Logf("110000000011 : %v", IsPalindrome(110000000011))
	t.Logf("1100001110000111 : %v", IsPalindrome(1100001110000111))
}

func BenchmarkIsPalindrome(b *testing.B) {
	b.Run("1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IsPalindrome(i)
		}
	})
	b.Run("2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IsPalindrome2(i)
		}
	})

}