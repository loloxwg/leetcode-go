package __14

import "fmt"

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}

func (node *TreeNode) Print() {
	fmt.Printf("%d ", node.Val)
}

func (node *TreeNode) SetValue(value int) {
	node.Val = value
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

func sumNumbers(root *TreeNode) int {
	var (
		dfs func(node *TreeNode, pre int) int
	)
	dfs = func(node *TreeNode, pre int) int {
		if node == nil {
			return 0
		}
		sum := pre*10 + node.Val
		if node.left == nil && node.right == nil {
			return sum
		}
		return dfs(node.left, sum) + dfs(node.right, sum)
	}
	return dfs(root, 0)
}

func main() {
	root := CreateNode(5)
	root.left = CreateNode(2)
	root.right = CreateNode(4)
	//root.left.right = CreateNode(7)
	//root.left.right.left = CreateNode(6)
	//root.right.left = CreateNode(8)
	//root.right.right = CreateNode(9)
	num := sumNumbers(root)
	println(num)
}
