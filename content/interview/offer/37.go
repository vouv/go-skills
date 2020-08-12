// +build OMIT

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Serialize(node *TreeNode) string {
	sb := strings.Builder{}
	inOrder(node, &sb)
	return strings.TrimRight(sb.String(), "null,")
}

func inOrder(node *TreeNode, buf *strings.Builder) {
	if node == nil {
		buf.WriteString("null,")
		return
	}
	buf.WriteString(strconv.Itoa(node.Val) + ",")
	inOrder(node.Left, buf)
	inOrder(node.Right, buf)
}

func DeSerialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	if len(vals) == 0 {
		return nil
	}
	_, root := reBuild(vals)
	return root
}

func reBuild(vals []string) ([]string, *TreeNode) {
	if len(vals) == 0 {
		return nil, nil
	}
	if vals[0] == "null" {
		return vals[1:], nil
	}
	val, _ := strconv.Atoi(vals[0])
	node := &TreeNode{Val: val}
	vals = vals[1:]
	vals, node.Left = reBuild(vals)
	vals, node.Right = reBuild(vals)
	return vals, node
}

func main() {
	var root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	data := Serialize(root)
	fmt.Println(data)

	res := DeSerialize(data)
	fmt.Println(Serialize(res))

	var root2 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	data2 := Serialize(root2)
	fmt.Println(data2)

	res2 := DeSerialize(data2)
	fmt.Println(Serialize(res2))

}
