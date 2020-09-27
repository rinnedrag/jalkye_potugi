package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	list3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	var lists = make([]*ListNode, 0)
	lists = append(lists, list1, list2, list3)

	printList(mergeKLists(lists))
}

func printList(n *ListNode) {
	for n != nil {
		fmt.Printf("%d ", n.Val)
		n = n.Next
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	var result *ListNode
	for _, list := range lists {
		result = merge(result, list)
	}
	return result
}

func merge(ListNode1 *ListNode, ListNode2 *ListNode) *ListNode {
	if ListNode1 == nil {
		return ListNode2
	}
	if ListNode2 == nil {
		return ListNode1
	}

	if ListNode1.Val < ListNode2.Val {
		ListNode1.Next = merge(ListNode1.Next, ListNode2)
		return ListNode1
	} else {
		ListNode2.Next = merge(ListNode1, ListNode2.Next)
		return ListNode2
	}
}
