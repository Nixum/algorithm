package monotonic_stack

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	nums := []int{}
	nums = []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(nums))

	nums = []int{4,2,0,3,2,5}
	fmt.Println(trap(nums))
}
