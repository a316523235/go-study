package disign

import (
	"fmt"
	"testing"
	"time"
)

func TestD1(t *testing.T)  {
	for i := 0; i < 100000; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Minute)
	}
}

func TestD2(t *testing.T)  {
	nums := []int{1,2,3,4,5,6,7,8,9,10,11,12}
	changeArr(nums)
	fmt.Println(nums)
}

func changeArr(nums []int)  {
	nums[2] = 5
}
