package differenceofsquares

func SquareOfSum(n int) int {
	x := 0
	for i := 1; i <= n; i++ {
         x += i
    }
    return x * x
}

func SumOfSquares(n int) int {
	x := 0
    for i := 1; i <= n; i++ {
        x += i * i
    }
    return x
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
