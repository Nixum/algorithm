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
