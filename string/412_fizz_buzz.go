package string

import "fmt"

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		if (i+1) % 5 == 0 && (i+1) % 3 == 0 {
			res[i] = "FizzBuzz"
		} else if (i+1) % 5 == 0 {
			res[i] = "Buzz"
		} else if (i+1) % 3 == 0 {
			res[i] = "Fizz"
		} else {
			res[i] = fmt.Sprintf("%d", i+1)
		}
	}
	return res
}
