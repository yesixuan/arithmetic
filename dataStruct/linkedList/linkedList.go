package linkedList

import "fmt"

type node struct {
	e    interface{}
	next *node
}

type LinkedList struct {
	size      int
	dummyHead *node
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{0, &node{}}
}

func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

// 找到待添加节点的前一个节点
func (list *LinkedList) Add(index int, e interface{}) {
	// todo index 防空
	prevNode := list.dummyHead
	for i := 0; i < index; i++ {
		prevNode = prevNode.next
	}
	prevNode.next = &node{e, prevNode.next}
	list.size++
}

func (list *LinkedList) AddFirst(e interface{}) {
	list.Add(0, e)
}

func (list *LinkedList) AddLast(e interface{}) {
	list.Add(list.size, e)
}

func (list *LinkedList) Get(index int) interface{} {
	currNode := list.dummyHead.next
	for i := 0; i < index; i++ {
		currNode = currNode.next
	}
	return currNode.e
}

func (list *LinkedList) Set(index int, e interface{}) {
	currNode := list.dummyHead.next
	for i := 0; i < index; i++ {
		currNode = currNode.next
	}
	currNode.e = e
}

func (list *LinkedList) Contains(e interface{}) bool {
	currNode := list.dummyHead.next
	for currNode.next != nil {
		if currNode.e == e {
			return true
		}
		currNode = currNode.next
	}
	return false
}

func (list *LinkedList) Print() {
	res := make([]interface{}, 0)
	currNode := list.dummyHead.next
	for currNode != nil {
		res = append(res, currNode.e)
		currNode = currNode.next
	}

	fmt.Println(res, list.size)
}
