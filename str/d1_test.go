package str

import (
	"fmt"
	"testing"
)

func TestFirstUniqChar(t *testing.T)  {
	str := "leetcode"
	index := firstUniqChar(str)
	fmt.Println(index)
}
