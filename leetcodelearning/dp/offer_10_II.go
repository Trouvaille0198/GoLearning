package dp

func numWays(n int) int {
	const mod = 1e9 + 7
	a, b, sum := 0, 1, 1
	for i := 2; i <= n; i++ {
		a = b
		b = sum
		sum = (a + b) % mod
	}
	return sum
}
