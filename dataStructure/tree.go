package dataStructure

type TreeNode struct {
	Data     interface{}
	Children []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *Tree) AddNode(data interface{}) {
	if t.Root == nil {
		t.Root = &TreeNode{Data: data}
	} else {
		t.Root.Children = append(t.Root.Children, &TreeNode{Data: data})
	}
}

func InitTree() *Tree {
	var tree *Tree = &Tree{Root: InitTreeNode(nil)}
	return tree
}

func InitTreeNode(data interface{}) *TreeNode {
	return &TreeNode{Data: data}
}
