package link

import (
	"fmt"
	"testing"
)

func TestD2(t *testing.T)  {
	tr := BuildList([]int{1,2,3,4,5})
	res := removeNthFromEnd(tr, 2)
	fmt.Println(res)
}