package array

type queue struct {
	Q []int
}

func (q *queue) back() int {
	return q.Q[len(q.Q)-1]
}

func (q *queue) front() int {
	return q.Q[0]
}

func (q *queue) len() int {
	return len(q.Q)
}

func (q *queue) push(num int) {
	for q.len() != 0 && num > q.back() {
		q.Q = q.Q[:q.len()-1]
	}
	q.Q = append(q.Q, num)
}

func (q *queue) pop(num int) {
	if q.len() != 0 && num == q.front() {
		q.Q = q.Q[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)
	q := &queue{Q: make([]int, 0)} // 从队头到队尾递减的单调队列
	for i := 0; i < k; i++ {
		q.push(nums[i])
	}
	result = append(result, q.front())

	i, j := 1, k
	for j < len(nums) {
		q.pop(nums[i-1])
		q.push(nums[j])
		result = append(result, q.front())
		i++
		j++
	}
	return result
}
