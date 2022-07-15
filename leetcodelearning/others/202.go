package others

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	hash := make(map[int]bool)
	sum := n
	for !hash[sum] {
		hash[sum] = true
		sum = getSum(sum)
		if sum == 1 {
			return true
		}
	}
	return false
}

// 非常妙的快慢指针
func isHappy(n int) bool {
	slowSum, fastSum := n, n
	for {
		slowSum = getSum(slowSum)
		fastSum = getSum(fastSum)
		fastSum = getSum(fastSum)
		if slowSum == fastSum {
			break
		}
	}
	return slowSum == 1
}
