package dp

const Mod int = 1e9 + 7

// 递归
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return (fib(n-1) + fib(n-2)) % Mod
}

// DP
func fib2(n int) int {
	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	a, b, sum := 0, 0, 1
	for i := 2; i <= n; i++ {
		a = b
		b = sum
		sum = (a + b) % mod
	}
	return sum
}
