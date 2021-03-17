package str

import (
	"fmt"
	"testing"
)

func TestRandInt(t *testing.T)  {
	var x rune
	x = 'b' - 'a'
	fmt.Println(x)
	 arr := []int{0, 0, 0}
	arr[x] = 11
	fmt.Println(arr)
}

func TestIsAnagram(t *testing.T)  {
	s := "A man, a plan, a canal: Panama"
	isRight := isPalindrome(s)
	fmt.Println(isRight)
}

func TestStrstr(t *testing.T)  {
	i := strStr("hello", "ll")
	fmt.Println(i)
}