package array

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix))

	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix))
}

func TestGetModifiedArray(t *testing.T) {
	updates := [][]int{
		{1,3,2},
		{2,4,3},
		{0,2,-2},
	}
	fmt.Println(getModifiedArray(5, updates))
}

func TestCorpFlightBookings(t *testing.T) {
	bookings := [][]int{
		{1,2,10},
		{2,3,20},
		{2,5,25},
	}
	fmt.Println(corpFlightBookings(bookings, 5))

	bookings = [][]int{
		{1,2,10},
		{2,2,15},
	}
	fmt.Println(corpFlightBookings(bookings, 2))
}

func TestCarPooling(t *testing.T) {
	trips := [][]int{
		{2,1,5},
		{3,3,7},
	}
	fmt.Println(carPooling(trips, 4))

	trips = [][]int{
		{2,1,5},
		{3,3,7},
	}
	fmt.Println(carPooling(trips, 5))

	trips = [][]int{
		{9,0,1},
		{3,3,7},
	}
	fmt.Println(carPooling(trips, 4))
}

func TestRotate(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)

	matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}

func TestNextPermutation(t *testing.T) {
	nums := []int{1,2,3}
	nextPermutation(nums)
	fmt.Println(nums)
}