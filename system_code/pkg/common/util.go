package common

import (
	"fmt"
	"strconv"
	"strings"
)

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

// LevelOrderTraversal 层序遍历打印
func LevelOrderTraversal(root *TreeNode, output bool) string {
	if root == nil {
		return "[]"
	}

	nums := make([]string, 0)

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			nums = append(nums, fmt.Sprintf("null"))
		} else {
			nums = append(nums, fmt.Sprintf("%d", node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == "null" {
			nums = nums[:i]
		} else {
			break
		}
	}
	if output {
		return fmt.Sprintf("%v", nums)
	} else {
		fmt.Printf("%v\n", nums)
		return fmt.Sprintf("%v", nums)
	}
}

// StrToBinaryTree 层序遍历的字符串转二叉树
func StrToBinaryTree(str string) *TreeNode {
	if str == "" {
		return nil
	}

	vals := parseStr(str)
	if len(vals) == 0 {
		return nil
	}

	root := &TreeNode{Val: str2int(vals[0])}
	queue := []*TreeNode{root}
	for i := 1; i < len(vals); {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != "null" {
			val, _ := strconv.Atoi(vals[i])
			node.Left = &TreeNode{Val: val}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != "null" {
			val, _ := strconv.Atoi(vals[i])
			node.Right = &TreeNode{Val: val}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// 解析字符串，返回节点值列表
func parseStr(str string) []string {
	return split(str)
}

func str2int(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

// 按照逗号分隔字符串，返回字符串数组
func split(str string) []string {
	return strings.Split(strings.Trim(str, "[]"), ",")
}

func StringToListNode(s string) *ListNode {
	s = strings.Trim(s, "[]")     // Remove brackets
	nums := strings.Split(s, ",") // Split string by comma

	var head, prev *ListNode

	for _, numStr := range nums {
		num, _ := strconv.Atoi(numStr) // Convert string to integer
		node := &ListNode{Val: num}

		if head == nil {
			head = node
		} else {
			prev.Next = node
		}
		prev = node
	}

	return head
}

func ListNodeToString(head *ListNode) string {
	var sb strings.Builder
	sb.WriteString("[")
	if head != nil && head.Val != 0 {
		for head != nil {
			sb.WriteString(strconv.Itoa(head.Val)) // Convert integer to string
			if head.Next != nil {
				sb.WriteString(",")
			}
			head = head.Next
		}
	}

	sb.WriteString("]")
	return sb.String()
}
