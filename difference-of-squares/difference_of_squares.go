package diffsquares

import "math"

func SquareOfSum(input int) int {
	sum := 0

	for i := 1; i <= input; i++ {
		sum += i
	}

	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(input int) int {
	sum := 0

	for i := 1; i <= input; i++ {
		sum += int(math.Pow(float64(i), 2))
	}

	return sum
}

func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}

// Solutions that use algebra instead of iteration:
// func SquareOfSums(n int) int {
// 	ret := (n * (n + 1)) / 2
// 	return ret * ret
// }

// func SumOfSquares(n int) int {
// 	return (n * (n + 1) * (2*n + 1)) / 6
// }

// func Difference(n int) int {
// 	return SquareOfSums(n) - SumOfSquares(n)
// }
