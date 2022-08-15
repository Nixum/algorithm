package hash

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}

func TestIntersection(t *testing.T) {
	var a1 []int
	var a2 []int

	a1 = []int{1, 2, 2, 1}
	a2 = []int{2, 2}
	fmt.Println(intersection(a1, a2))

	a1 = []int{4,9,5}
	a2 = []int{9,4,9,8,4}
	fmt.Println(intersection(a1, a2))
}

func TestIsHappy(t *testing.T) {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(111))
}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
}

func TestFourSumCount(t *testing.T) {
	var a []int
	var b []int
	var c []int
	var d []int

	a = []int{1, 2}
	b = []int{-2, -1}
	c = []int{-1, 2}
	d = []int{0, 2}
	fmt.Println(fourSumCount(a, b, c, d))

	a = []int{0}
	b = []int{0}
	c = []int{0}
	d = []int{0}
	fmt.Println(fourSumCount(a, b, c, d))
}

func TestCanConstruct(t *testing.T) {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
