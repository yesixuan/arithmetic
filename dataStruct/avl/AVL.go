package avl

import (
	"fmt"
	"math"
)

type node struct {
	e           int
	left, right *node
	height      float64
}

type AVL struct {
	root *node
	size int
}

func getHeight(node *node) float64 {
	if node == nil {
		return 0
	}
	return node.height
}

// 计算平衡因子
func getBalanceFactor(node *node) float64 {
	return getHeight(node.left) - getHeight(node.right)
}

// 左旋
func leftRotate(node *node) *node {
	rightNode := node.right // 这个将来要作为根节点返回
	leftTree := rightNode.left
	rightNode.left = node
	node.right = leftTree
	// 重新计算高度
	node.height = math.Max(getHeight(node.left), getHeight(node.right))
	rightNode.height = math.Max(getHeight(rightNode.left), getHeight(rightNode.right))
	return rightNode
}

// 右旋
func rightRotate(node *node) *node {
	leftNode := node.left
	rightTree := leftNode.right
	leftNode.right = node
	node.left = rightTree
	// 重计算高度
	node.height = math.Max(getHeight(node.left), getHeight(node.right)) + 1
	leftNode.height = math.Max(getHeight(leftNode.left), getHeight(leftNode.right)) + 1
	return leftNode
}

func CreateAVL() *AVL {
	return &AVL{nil, 0}
}

func (avl *AVL) IsEmpty() bool {
	return avl.size == 0
}

func add(root *node, e int) *node {
	// 	递归终止条件
	if root == nil {
		return &node{e, nil, nil, 0}
	}
	if root.e > e {
		root.left = add(root.left, e)
	} else {
		root.right = add(root.right, e)
	}
	// 更新 height
	root.height = 1 + math.Max(getHeight(root.left), getHeight(root.right))
	// 计算平衡因子
	balanceFactor := getBalanceFactor(root)
	// LL
	if balanceFactor > 1 && getBalanceFactor(root.left) >= 0 {
		return rightRotate(root)
	}
	// RR
	if balanceFactor < -1 && getBalanceFactor(root.right) >= 0 {
		return leftRotate(root)
	}
	// LR
	if balanceFactor > 1 && getBalanceFactor(root.left) < 0 {
		root.left = leftRotate(root.left)
		return rightRotate(root)
	}
	// RL
	if balanceFactor < -1 && getBalanceFactor(root.right) > 0 {
		root.right = rightRotate(root.right)
		return leftRotate(root)
	}
	return root
}

func (avl *AVL) Add(e int) {
	avl.root = add(avl.root, e)
	avl.size++
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

func (avl *AVL) Contains(e int) bool {
	return contains(avl.root, e)
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

func (avl *AVL) PreOrder() {
	preOrder(avl.root)
}

// 非遍历实现（利用栈的特性）
func (avl *AVL) PreOrderNR() {
	var currNode *node
	stack := make([]*node, 0)
	// 根节点入栈
	stack = append(stack, avl.root)
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
func (avl *AVL) LeveOrder() {
	var currNode *node
	queue := make([]*node, 0)
	queue = append(queue, avl.root)
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

func (avl *AVL) RemoveMin() *node {
	minNode := minimum(avl.root)
	removeMin(avl.root)
	avl.size--
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
func (avl *AVL) Remove(e int) {
	avl.root = remove(avl.root, e)
	avl.size--
}
