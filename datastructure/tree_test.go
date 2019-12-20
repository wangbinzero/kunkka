package datastructure

import (
	"fmt"
	"testing"
)

func TestTree_Add(t *testing.T) {
	tree := Tree{}
	tree.Add(0)
	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
	tree.Add(4)
	tree.Add(5)
	tree.Add(6)
	tree.Add(7)
	tree.Add(8)
	tree.Add(9)
	tree.Add(10)
	tree.Add(11)
	tree.Add(12)
	//fmt.Println("深度遍历开始")
	//tree.PreOrder(tree.RootNode)
	//fmt.Println("深度遍历结束")

	//fmt.Println("广度遍历开始")
	//tree.BreadthTravel()
	//fmt.Println("广度遍历结束")

	fmt.Println("中序遍历开始")
	tree.InOrder(tree.RootNode)
	fmt.Println("中序遍历结束")

}
