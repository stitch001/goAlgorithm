package tree

import (
	"GoAlgo/myUtils"
	"fmt"
	"math"
)

//树节点
type treeNode struct {
	Val    interface{}
	LNode  *treeNode
	RNode  *treeNode
	Parent *treeNode
}

func createBTree(nodesVal []interface{}) *treeNode {
	if len(nodesVal) == 0 {
		return nil
	}
	level := math.Log2(float64(len(nodesVal))) + 1
	var root = &treeNode{Val: nodesVal[0]}
	nodes := []*treeNode{root}
	for i := 1; i < int(level); i++ {
		// 2**0=1 2**1=2 2**2=4
		levelLen := int(math.Pow(2, float64(i)))
		for j := 0; j < levelLen; j++ {
			// 父节点 node / 2
			parent := ((levelLen + j) / 2) - 1
			if levelLen+j > len((nodesVal)) {
				break
			}
			//    1    2**0+0
			//  2   3  2**1+0 2**1+1
			// 4 5 6 7 2**2+0 +1 +2 +3  对应数组元素还需要减一
			node := &treeNode{Val: nodesVal[levelLen+j-1]}
			node.Parent = nodes[parent]
			nodes = append(nodes, node)
			if j%2 == 0 {
				nodes[parent].LNode = node

			} else {
				nodes[parent].RNode = node
			}

		}
	}
	return root
}

func (root *treeNode) treeHeight() int {
	if root == nil {
		return 0
	}

	return int(1 + math.Max(float64(root.LNode.treeHeight()), float64(root.RNode.treeHeight())))
}

//递归先序遍历
func (root *treeNode) PreOrder(traverseFunc func(val interface{})) {
	if root == nil {
		return
	}
	fmt.Printf("%v ", root.Val)
	root.LNode.PreOrder(traverseFunc)
	root.RNode.PreOrder(traverseFunc)
}

// 层次遍历
func (root *treeNode) LevelTraverse(traverseFunc func(val interface{})) {
	if root == nil {
		return
	}
	tlist := []*treeNode{root}
	level := 1

	for len(tlist) > 0 {
		fmt.Println("level :", level)
		for i := 0; i < len(tlist); i++ {
			treeNode := tlist[0]
			tlist = tlist[1:]
			traverseFunc(treeNode.Val)
			if treeNode.Parent != nil {
				fmt.Printf(" Parent: %v ", treeNode.Parent.Val)
			}
			if treeNode.LNode != nil {
				tlist = append(tlist, treeNode.LNode)
			}
			if treeNode.RNode != nil {
				tlist = append(tlist, treeNode.RNode)
			}
		}
		fmt.Println("")
		level++
	}
}

func traverse(val interface{}) {
	fmt.Printf("TreeNode val %v ", val)
	fmt.Print("hello", val, "\n")
}

func Run() {

	root := createBTree(myUtils.ToInterfaceArray(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println("先序遍历")
	root.PreOrder(traverse)
	fmt.Println("")
	root.LevelTraverse(traverse)
	fmt.Println("Tree Height:", root.treeHeight())
}
