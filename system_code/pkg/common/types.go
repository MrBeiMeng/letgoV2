package common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) LoadFromStr(str string) error {
	tree := StrToBinaryTree(str)
	t.Val = tree.Val
	t.Left = tree.Left
	t.Right = tree.Right

	return nil
}

func (t *TreeNode) String() string {
	return "[" + LevelOrderTraversal(t, true) + "]"
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return ListNodeToString(l)
}
