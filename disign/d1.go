package disign

import (
	"math/rand"
)

//var res [][]int

type Solution struct {
	source []int
}


func Constructor2(nums []int) Solution {
	obj := Solution {
		source: nums,
	}
	return obj
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.source
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	temp := make([]int, len(this.source))
	copy(temp, this.source)
	cnt := len(temp)
	for i := 0; i < cnt; i++ {
		n := rand.Intn(cnt)
		temp[i], temp[n] = temp[n], temp[i]
	}
	return temp
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
