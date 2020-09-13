package bst

import (
	"fmt"
)

type node struct {
	e           int
	left, right *node
}

type BST struct {
	root *node
	size int
}

func CreateBST() *BST {
	return &BST{nil, 0}
}

func (bst *BST) IsEmpty() bool {
	return bst.size == 0
}

func add(root *node, e int) *node {
	// 	递归终止条件
	if root == nil {
		return &node{e, nil, nil}
	}
	if root.e > e {
		root.left = add(root.left, e)
	} else {
		root.right = add(root.right, e)
	}
	return root
}

func (bst *BST) Add(e int) {
	bst.root = add(bst.root, e)
	bst.size++
}

func contains(node *node, e int) bool {
	if node == nil {
		return false
	}
	if node.e == e {
		return true
	}
	if node.e > e {
		return contains(node.left, e)
	} else {
		return contains(node.right, e)
	}
}

func (bst *BST) Contains(e int) bool {
	return contains(bst.root, e)
}

// 前序遍历（中序与后续就不再赘述）中序遍历就将二分搜索树排序了
func preOrder(node *node) {
	if node == nil {
		return
	}
	fmt.Println(node.e)
	preOrder(node.left)
	preOrder(node.right)
}

func (bst *BST) PreOrder() {
	preOrder(bst.root)
}

// 非遍历实现（利用栈的特性）
func (bst *BST) PreOrderNR() {
	var currNode *node
	stack := make([]*node, 0)
	// 根节点入栈
	stack = append(stack, bst.root)
	for len(stack) > 0 {
		currNode = stack[len(stack)-1]
		fmt.Println(currNode.e)
		stack = stack[:len(stack)-1]
		if currNode.right != nil {
			// 右节点入栈
			stack = append(stack, currNode.right)
		}
		if currNode.left != nil {
			// 左节点入栈
			stack = append(stack, currNode.left)
		}
	}
}

// 借助队列实现层序遍历
func (bst *BST) LeveOrder() {
	var currNode *node
	queue := make([]*node, 0)
	queue = append(queue, bst.root)
	for len(queue) > 0 {
		currNode = queue[0]
		queue = queue[1:]
		fmt.Println(currNode.e)
		if currNode.left != nil {
			queue = append(queue, currNode.left)
		}
		if currNode.right != nil {
			queue = append(queue, currNode.right)
		}
	}
}

// 寻找最小元素
func minimum(node *node) *node {
	if node.left == nil {
		return node
	}
	return minimum(node.left)
}

// 删除以传入节点为根节点的子树的最小值，返回该节点
func removeMin(node *node) *node {
	if node.left == nil {
		return node.right
	}
	node.left = removeMin(node.left)
	return node
}

func (bst *BST) RemoveMin() *node {
	minNode := minimum(bst.root)
	removeMin(bst.root)
	bst.size--
	return minNode
}

func remove(node *node, e int) *node {
	if node == nil {
		panic("没有找到目标元素")
	}
	if node.e > e {
		node.left = remove(node.left, e)
		return node
	} else if node.e < e {
		node.right = remove(node.right, e)
		return node
	} else { // node.e == e
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			res := node
			res.left = node.left
			res.right = removeMin(node.right)
			return res
		}
	}

}

// 删除任意节点
// 1. 找到目标节点 2. 找到目标节点的后继（也就是其右子树的最小节点，前驱也是可以的） 3. 用后继替换目标节点
func (bst *BST) Remove(e int) {
	bst.root = remove(bst.root, e)
	bst.size--
}
