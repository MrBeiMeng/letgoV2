package common

import (
	"fmt"
	"reflect"
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

type SplitUtilI interface {
	// GetLevelElements level from 0
	GetLevelElements(level int, withChild bool) []string
	LevelNum() (min, max int)
}

type SplitUtil struct {
	data []interface{}
}

func iterateSlice(minLevel, maxLevel *int, slice []interface{}) {

	// 如果所有元素都是数组类型，则minLevel加一
	// 有一个是数组类型，maxLevel加一
	hasSubArr, allSubArr := false, true

	for i := 0; i < len(slice); i++ {
		if value, ok := slice[i].([]interface{}); ok {
			iterateSlice(minLevel, maxLevel, value)
			hasSubArr = true
		} else {
			allSubArr = false
		}
	}

	if hasSubArr {
		*maxLevel += 1
	}

	if allSubArr {
		*minLevel += 1
	}

}

func (s *SplitUtil) LevelNum() (min, max int) {

	iterateSlice(&min, &max, s.data)

	return min, max
}

func sliceLevelElements(targetLevel, currentLevel int, slice []interface{}, withChild bool, result *[]string) {

	for i := 0; i < len(slice); i++ {
		tempValue := slice[i]
		if currentLevel < targetLevel {
			if value, ok := tempValue.([]interface{}); ok {
				sliceLevelElements(targetLevel, currentLevel+1, value, withChild, result)
			}

			continue
		}

		if _, ok := tempValue.([]interface{}); !ok || withChild {
			if reflect.TypeOf(tempValue).Kind() == reflect.Uint8 {
				*result = append(*result, fmt.Sprintf("%c", tempValue))
			} else {
				*result = append(*result, fmt.Sprintf("%s", tempValue))
			}
		}

	}

}

func (s *SplitUtil) GetLevelElements(level int, withChild bool) []string {
	tempLevelElements := make([]string, 0)

	sliceLevelElements(level, 0, s.data, withChild, &tempLevelElements)

	return tempLevelElements
}

func isSwapBySquareBrackets(str string) bool {
	if len(str) < 2 {
		return false
	}

	return str[0] == '[' && str[len(str)-1] == ']'
}

func test1(strData *[]byte) (result []interface{}) {

	// 先检测是否有特殊标识
	if (*strData)[0] == '[' {
		*strData = (*strData)[1:] // 去除外面括号
	}

	for len(*strData) > 0 {
		temChar := (*strData)[0]
		if temChar == '@' {
			tempElement := strings.Builder{}
			for len(*strData) > 0 {
				subTempChar := (*strData)[0]

				tempElement.WriteByte(subTempChar)

				*strData = (*strData)[1:]
				if subTempChar == ']' {
					result = append(result, tempElement.String())
					break
				}

			}

			continue
		}

		if temChar == '[' {

			result = append(result, test1(strData))
			continue
		}

		if !strings.Contains("[],", string(temChar)) {
			result = append(result, temChar)
		}

		*strData = (*strData)[1:]
		if temChar == ']' {
			break
		}
	}

	return result
}

// NewSplitUtil 这个类负责拆分字符串，像是[[1,2,3],2,3]
// 将他们拆分成层级
// 如果任意层级包含@则不进行拆分
// 前提：[]是分组标识，, 是间隔字符
func NewSplitUtil(strData []byte) *SplitUtil {
	obj := &SplitUtil{}
	obj.data = make([]interface{}, 0)

	obj.data = test1(&strData)

	//// 先检测是否有特殊标识
	//if isSwapBySquareBrackets(string(strData)) {
	//	strData = strData[1 : len(strData)-1] // 去除外面括号
	//}
	//
	//for len(strData) > 0 {
	//	tempChar := strData[0]
	//	if tempChar == '[' {
	//		obj.data = append(obj.data, test1(&strData))
	//		continue
	//	}
	//	if !strings.Contains("[],", string(tempChar)) {
	//		obj.data = append(obj.data, tempChar)
	//	}
	//
	//	strData = strData[1:]
	//
	//	if tempChar == ']' {
	//		break
	//	}
	//}

	return obj
}
