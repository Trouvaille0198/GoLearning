package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	hash := make(map[int]int)
	// 构建中序遍历的哈希表
	for i, val := range inorder {
		hash[val] = i
	}

	var bt func(preorder []int, inorder []int, offset int) *TreeNode
	bt = func(preorder []int, inorder []int, offset int) *TreeNode {
		// 为了方便哈希表 添加offset参数为相对0位置的偏移量
		if len(preorder) == 0 {
			return nil
		}
		root := TreeNode{Val: preorder[0]} // 找到根节点
		i := hash[root.Val] - offset       // 找到根节点在中序遍历中的索引位置

		root.Left = bt(preorder[1:i+1], inorder[:i], offset)
		root.Right = bt(preorder[i+1:], inorder[i+1:], offset+i+1)
		return &root
	}
	root := bt(preorder, inorder, 0)
	return root
}
