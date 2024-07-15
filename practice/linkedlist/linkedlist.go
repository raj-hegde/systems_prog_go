package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func addNode(n *Node, v int) int {
	if root == nil {
		n = &Node{v, nil}
		root = n
		return 0
	}

	if v == n.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if n.Next == nil {
		n.Next = &Node{v, nil}
		return -2
	}

	return addNode(n.Next, v)
}

func traverse(n *Node) {
	if n == nil {
		fmt.Println("-> Emptly list!")
		return
	}
	for n != nil {
		fmt.Printf("%d -> ", n.Value)
		n = n.Next
	}
	fmt.Println()
}

var root = new(Node)

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
}
