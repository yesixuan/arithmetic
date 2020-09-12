package main

import "arithmetic/dataStruct/linkedList"

func main() {
	list := linkedList.CreateLinkedList()
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)
	list.AddLast(4)
	list.Print()
}
