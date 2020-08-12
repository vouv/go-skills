package main

type Node struct {
	Left  *Node
	Right *Node
}

func visit(node *Node) {} // 表示遍历这个节点

func preOrder(root *Node) {
	if root == nil {
		return
	}
	visit(root)
	preOrder(root.Left)
	preOrder(root.Right)
}

func inOrder(root *Node) {
	if root == nil {
		return
	}
	preOrder(root.Left)
	visit(root)
	preOrder(root.Right)
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	preOrder(root.Left)
	preOrder(root.Right)
	visit(root)
}

func main() {

}
