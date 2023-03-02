package dataStructure

type TreeNode struct {
	Data     interface{}
	Children []*TreeNode
}

func (tn *TreeNode) AddNode(data interface{}) bool {
	// Root 없으면 false, 있으면 자식 추가하고 true
	if tn.Data == nil {
		return false
	} else {
		tn.Children = append(tn.Children, &TreeNode{Data: data})
		return true
	}
}

func MakeTreeNode(data interface{}) *TreeNode {
	return &TreeNode{Data: data}
}

// preorder
func DFSWithRecursive(startNode *TreeNode) []interface{} {
	var result []interface{}
	result = append(result, startNode.Data)
	for _, child := range startNode.Children {
		result = append(result, DFSWithRecursive(child)...)
	}
	return result
}

func DFSWithStack(startNode *TreeNode) []interface{} {
	var stack []*TreeNode
	var result []interface{}
	stack = append(stack, startNode)
	for len(stack) > 0 {
		var last *TreeNode
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		result = append(result, last.Data)
		stack = append(stack, last.Children...)
	}
	return result
}

func BFS(startNode *TreeNode) []interface{} {
	var queue []*TreeNode
	var result []interface{}
	queue = append(queue, startNode)
	for len(queue) > 0 {
		var first *TreeNode
		first, queue = queue[0], queue[1:]
		result = append(result, first.Data)
		queue = append(queue, first.Children...)
	}
	return result
}
