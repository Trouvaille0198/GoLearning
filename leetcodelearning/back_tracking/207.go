package back_tracking

func canFinish(numCourses int, prerequisites [][]int) bool {
	count := 0
	// 入度表
	inDegrees := make([]int, numCourses)
	for _, p := range prerequisites {
		inDegrees[p[0]]++
	}
	// 邻接表
	edgeMap := map[int][]int{}
	for _, p := range prerequisites {
		if _, ok := edgeMap[p[1]]; !ok {
			edgeMap[p[1]] = []int{p[0]}
		} else {
			edgeMap[p[1]] = append(edgeMap[p[1]], p[0])
		}
	}
	// 队列
	queue := make([]int, 0)
	// 入度为0的节点入队
	for i, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		// 出队 代表学过
		i := queue[0]

		count++
		if edgeNodes, ok := edgeMap[i]; ok && len(edgeNodes) > 0 {
			// i有后续节点
			for _, edgeNode := range edgeNodes {
				// 后续节点入度-1
				inDegrees[edgeNode]--
				if inDegrees[edgeNode] == 0 {
					// 若后续节点入度为0，入队
					queue = append(queue, edgeNode)
				}
			}
		}
	}
	return count == numCourses
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	count := 0
	// 入度表
	inDegrees := make([]int, numCourses)
	for _, p := range prerequisites {
		inDegrees[p[0]]++
	}
	// 邻接表
	edgeMap := map[int][]int{}
	for _, p := range prerequisites {
		if _, ok := edgeMap[p[1]]; !ok {
			edgeMap[p[1]] = []int{p[0]}
		} else {
			edgeMap[p[1]] = append(edgeMap[p[1]], p[0])
		}
	}

	var dfs func(node int)
	dfs = func(node int) {
		count++
		if edgeNodes, ok := edgeMap[node]; ok && len(edgeNodes) > 0 {
			// 遍历后续节点
			for _, edgeNode := range edgeNodes {
				inDegrees[edgeNode]--
				if inDegrees[edgeNode] == 0 {
					// 入度为0 可以学了
					dfs(edgeNode)
				}
			}
		}
	}
	// 记录一下入度为0的节点 因为之后dfs的时候会变
	zeroInDegrees := make([]int, 0)
	for i, inDegree := range inDegrees {
		if inDegree == 0 {
			zeroInDegrees = append(zeroInDegrees, i)
		}
	}
	for _, zeroInDegree := range zeroInDegrees {
		dfs(zeroInDegree)
	}
	return count == numCourses
}
