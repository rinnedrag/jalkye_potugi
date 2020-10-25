package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	a := levelOrder(tree)
	fmt.Println(a)
}

type LevelNode struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var list = make([]*LevelNode, 0)
	list = append(list, &LevelNode{
		node:  root,
		level: 0,
	})
	var output = make([][]int, 0)
	for len(list) != 0 {
		node := list[0]
		if len(output) <= node.level {
			output = append(output, []int{node.node.Val})
		} else {
			output[node.level] = append(output[node.level], node.node.Val)
		}
		list = list[1:]
		if node.node.Left != nil {
			list = append(list, &LevelNode{
				node:  node.node.Left,
				level: node.level + 1,
			})
		}
		if node.node.Right != nil {
			list = append(list, &LevelNode{
				node:  node.node.Right,
				level: node.level + 1,
			})
		}
	}
	return output
}
