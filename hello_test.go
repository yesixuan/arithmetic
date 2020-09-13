package test

import (
	"arithmetic/dataStruct/bst"
	"arithmetic/dataStruct/linkedList"
	"testing"
)

func TestList(t *testing.T) {
	list := linkedList.CreateLinkedList()
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)
	list.AddLast(4)
	list.Print()
}

func TestBST(t *testing.T) {
	bst := bst.CreateBST()
	bst.Add(2)
	bst.Add(3)
	bst.Add(1)
	bst.Add(4)
	bst.Add(0)
	// r1 := bst.Contains(0)
	// r2 := bst.Contains(3)
	// r3 := bst.Contains(1)
	// fmt.Println(r1, r2, r3)
	// bst.PreOrder()
	// bst.PreOrderNR()
	// bst.LeveOrder()
	// bst.RemoveMin()
	bst.Remove(3)
	bst.PreOrderNR()
}
