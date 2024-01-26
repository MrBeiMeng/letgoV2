package common

import "fmt"

func treeToString(root *TreeNode) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return fmt.Sprintf("%d", root.Val)
	}
	if root.Right == nil {
		return fmt.Sprintf("%d,%s", root.Val, treeToString(root.Left))
	}
	if root.Left == nil {
		return fmt.Sprintf("%d,nil,%s", root.Val, treeToString(root.Right))
	}
	return fmt.Sprintf("%d,%s,%s", root.Val, treeToString(root.Left), treeToString(root.Right))
}
