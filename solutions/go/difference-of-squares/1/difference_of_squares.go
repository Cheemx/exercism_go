package diffsquares

func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
    return sum * sum
}

func SumOfSquares(n int) int {
    sum := 0
	for i := 1; i <= n; i++ {
        sum += i * i
    }
    return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
