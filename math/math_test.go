package math

import (
	"fmt"
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	fmt.Println(trailingZeroes(125))
}

func TestMedianFinder(t *testing.T) {
	f := MedianFinderConstructor()
	f.AddNum(2)
	f.AddNum(1)
	fmt.Println(f.FindMedian())
	f.AddNum(3)
	fmt.Println(f.FindMedian())
}
