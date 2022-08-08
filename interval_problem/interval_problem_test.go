package interval_problem

import (
	"fmt"
	"testing"
)

func TestIntervalIntersection(t *testing.T) {
	first := [][]int{
		{0, 2}, {5, 10}, {13, 23}, {24, 25},
	}
	second := [][]int{
		{1, 5}, {8, 12}, {15, 24}, {25, 26},
	}
	//t.Fatalf()
	fmt.Println(intervalIntersection(first, second))
	first = [][]int{
		{1, 3}, {5, 9},
	}
	second = [][]int{}
	fmt.Println(intervalIntersection(first, second))

	first = [][]int{
		{1, 3},
	}
	second = [][]int{
		{1, 3},
	}
	fmt.Println(intervalIntersection2(first, second))
}

func TestRemoveCoveredIntervals(t *testing.T) {
	first := [][]int{
		{1, 4}, {3, 6}, {2, 8},
	}
	fmt.Println(removeCoveredIntervals(first))
	first = [][]int{
		{1, 4}, {2, 3},
	}
	fmt.Println(removeCoveredIntervals(first))

	first = [][]int{
		{1, 3},
	}
	fmt.Println(removeCoveredIntervals(first))

	first = [][]int{
		{3, 10}, {4, 10}, {5, 11},
	}
	fmt.Println(removeCoveredIntervals(first))

	first = [][]int{
		{1, 2}, {1, 4}, {3, 4},
	}
	fmt.Println(removeCoveredIntervals(first))

	first = [][]int{
		//{34335,39239},{15875,91969},{29673,66453},{53548,69161},{40618,93111},
		{3, 4}, {2, 6}, {4, 7}, {1, 8}, {3, 9},
	}
	fmt.Println(removeCoveredIntervals(first))

	first = [][]int{
		{34335,39239},{15875,91969},{29673,66453},{53548,69161},{40618,93111},
		//{3, 4}, {2, 6}, {4, 7}, {1, 8}, {3, 9},
	}
	fmt.Println(removeCoveredIntervals(first))
}

func TestMerge(t *testing.T) {
	var intervals [][]int
	intervals = [][]int{
		{1, 3}, {2, 6}, {8, 10}, {15,18},
	}
	fmt.Println(merge(intervals))

	intervals = [][]int{
		{1, 3}, {3, 6},
	}
	fmt.Println(merge(intervals))

	intervals = [][]int{
		{1, 3},
	}
	fmt.Println(merge(intervals))

	intervals = [][]int{
		{1, 3}, {1, 3},
	}
	fmt.Println(merge(intervals))

	intervals = [][]int{
		{1, 4}, {2, 3},
	}
	fmt.Println(merge(intervals))
}

func TestEraseOverlapIntervals(t *testing.T) {
	var intervals [][]int
	intervals = [][]int{
		{1, 2}, {2, 3}, {3, 4},{1,3},
	}
	fmt.Println(eraseOverlapIntervals2(intervals))

	intervals = [][]int{
		{1, 2}, {1, 2}, {1, 2},
	}
	fmt.Println(eraseOverlapIntervals2(intervals))

	intervals = [][]int{
		{1, 2}, {2, 3},
	}
	fmt.Println(eraseOverlapIntervals2(intervals))

	intervals = [][]int{
		{-52,31},{-73,-26},{82,97},{-65,-11},{-62,-49},{95,99},{58,95},{-31,49},{66,98},{-63,2},{30,47},{-40,-26},
	}
	fmt.Println(eraseOverlapIntervals2(intervals))
}

func TestFindMinArrowShots(t *testing.T) {
	var intervals [][]int
	intervals = [][]int{
		{10, 16}, {2, 8}, {1, 6},{7,12},
	}
	fmt.Println(findMinArrowShots(intervals))

	intervals = [][]int{
		{1, 2}, {3, 4}, {5, 6}, {7, 8},
	}
	fmt.Println(findMinArrowShots(intervals))

	intervals = [][]int{
		{1, 2}, {2, 3},{3, 4}, {4, 5},
	}
	fmt.Println(findMinArrowShots(intervals))

	intervals = [][]int{
		{2, 3}, {2, 3},
	}
	fmt.Println(findMinArrowShots(intervals))

	intervals = [][]int{
		{2, 3},
	}
	fmt.Println(findMinArrowShots(intervals))
}