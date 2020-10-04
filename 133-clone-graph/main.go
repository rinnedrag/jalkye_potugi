package main

import "fmt"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Clone Graph.
//Memory Usage: 2.7 MB, less than 83.86% of Go online submissions for Clone Graph.

type Node struct {
	Val       int
	Neighbors []*Node
}

func main() {
	node1 := &Node{
		Val: 1,
	}
	node2 := &Node{
		Val: 2,
	}
	node3 := &Node{
		Val: 3,
	}
	node4 := &Node{
		Val: 4,
	}

	node1.Neighbors = []*Node{
		node2, node4,
	}
	node2.Neighbors = []*Node{
		node1, node3,
	}
	node3.Neighbors = []*Node{
		node2, node4,
	}
	node4.Neighbors = []*Node{
		node1, node3,
	}

	var nodes = make(map[int]*Node, 0)
	setNewNodes(nodes, node1)
	setConnections(nodes, node1)
	fmt.Println(node1)
}

func setNewNodes(nodes map[int]*Node, node *Node) {
	if _, ok := nodes[node.Val]; !ok {
		nodes[node.Val] = &Node{
			Val:       node.Val,
			Neighbors: make([]*Node, 0),
		}
	} else {
		return
	}
	for _, neighbor := range node.Neighbors {
		setNewNodes(nodes, neighbor)
	}
}

func setConnections(nodes map[int]*Node, node *Node) {
	if len(nodes[node.Val].Neighbors) != 0 {
		return
	}
	for _, neighbor := range node.Neighbors {
		nodes[node.Val].Neighbors = append(nodes[node.Val].Neighbors, nodes[neighbor.Val])
		setConnections(nodes, neighbor)
	}
}
