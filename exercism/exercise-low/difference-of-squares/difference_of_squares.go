// Package diffsquares is a library for determine diff between square of sum and sum of squares
package diffsquares

// SquareOfSum return square of sum natural digits
func SquareOfSum(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

// SumOfSquares return sum of squares natural digits
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference return difference between square of sum and sum of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
