package datastructure

import "fmt"

type Object interface {
}

// indicate tree node
type TreeNode struct {
	Data       Object
	LeftChild  *TreeNode
	RightChild *TreeNode
}

// indicate the tree of root node
type Tree struct {
	RootNode *TreeNode
}

func (this *Tree) Add(o Object) {
	node := &TreeNode{Data: o}

	// if root node is nil then add current node to root node
	if this.RootNode == nil {
		this.RootNode = node
		return
	}

	queue := []*TreeNode{this.RootNode}
	for len(queue) != 0 {
		c_node := queue[0]
		queue = queue[1:]
		if c_node.LeftChild == nil {
			c_node.LeftChild = node
			return
		} else {
			queue = append(queue, c_node.LeftChild)
		}

		if c_node.RightChild == nil {
			c_node.RightChild = node
			return
		} else {
			queue = append(queue, c_node.RightChild)
		}
	}
}

//广度遍历
func (this *Tree) BreadthTravel() {
	if this.RootNode == nil {
		return
	}

	queue := []*TreeNode{this.RootNode}
	//queue = append(queue, this.RootNode)
	for len(queue) != 0 {
		c_node := queue[0]
		queue = queue[1:]
		fmt.Printf("current node data is: %v\n", c_node.Data)

		if c_node.LeftChild != nil {
			queue = append(queue, c_node.LeftChild)
		}
		if c_node.RightChild != nil {
			queue = append(queue, c_node.RightChild)
		}
	}
}

//深度遍历
//先序遍历：root -> left -> right
//中序遍历：left -> center -> right
//后序遍历：left -> right -> root
func (this *Tree) PreOrder(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("current node data is: %v\n", node.Data)
	this.PreOrder(node.LeftChild)
	this.PreOrder(node.RightChild)
}

func (this *Tree) InOrder(node *TreeNode) {
	if node == nil {
		return
	}
	this.InOrder(node.LeftChild)
	fmt.Printf("current node data is: %v\n", node.Data)
	this.InOrder(node.RightChild)
}

func (this *Tree) PostOrder(node *TreeNode) {
	if node == nil {
		return
	}
	this.PostOrder(node.LeftChild)
	this.PostOrder(node.LeftChild)
	fmt.Printf("current node data is: %v\n", node.Data)
}
