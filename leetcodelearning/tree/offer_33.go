package tree

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	rootVal := postorder[len(postorder)-1]
	var i int
	for i = len(postorder) - 1; i >= 0; i-- {
		// 找到第一个小于rootVal的下标
		if postorder[i] < rootVal {
			break
		}
	}

	for j := i; j >= 0; j-- {
		if postorder[j] > rootVal {
			return false
		}
	}

	return verifyPostorder(postorder[:i+1]) && verifyPostorder(postorder[i+1:len(postorder)-1])
}
