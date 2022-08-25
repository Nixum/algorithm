package queue

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1,1,1,2,2,3,3,3,3}
	fmt.Println(topKFrequent(nums, 2))
}
