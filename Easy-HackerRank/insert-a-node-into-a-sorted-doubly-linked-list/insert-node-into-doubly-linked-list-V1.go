// Link to problem statement on hackerrank:
// https://www.hackerrank.com/challenges/insert-a-node-into-a-sorted-doubly-linked-list/problem
// This problem is solved using a Doubly "Linked List" data structure.
package main

import "fmt"

var (
	number1 int32
	number2 int32
	number3 int32
	number4 int32
	number5 int32
)

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

func (doublyLinkedList *DoublyLinkedList) insertNodeIntoDoublyLinkedList(nodeData int32) {
	node := &DoublyLinkedListNode{
		next: nil,
		prev: nil,
		data: nodeData,
	}

	if doublyLinkedList.head == nil {
		doublyLinkedList.head = node
	} else {
		doublyLinkedList.tail.next = node
		node.prev = doublyLinkedList.tail
	}
	doublyLinkedList.tail = node
}

func (doublyLinkedList DoublyLinkedList) sortedInsert(data int32) {
	var lastElementOfTheList *DoublyLinkedListNode
	var shouldBreak bool
	shouldBreak = false
	for doublyLinkedList.head != nil {
		if doublyLinkedList.head.next != nil && data > doublyLinkedList.head.data &&
			data < doublyLinkedList.head.next.data || data == doublyLinkedList.head.data {
			node := &DoublyLinkedListNode{
				data: data,
				prev: doublyLinkedList.head,
				next: doublyLinkedList.head.next,
			}
			doublyLinkedList.head.next = node
			shouldBreak = true
			break
		}
		lastElementOfTheList = doublyLinkedList.head
		doublyLinkedList.head = doublyLinkedList.head.next
	}
	if shouldBreak == false {
		node := &DoublyLinkedListNode{
			data: data,
			prev: lastElementOfTheList,
			next: nil,
		}
		lastElementOfTheList.next = node
	}
}

func (doublyLinkedList *DoublyLinkedList) Display() {
	for doublyLinkedList.head != nil {
		fmt.Printf("%v <-> ", doublyLinkedList.head.data)
		doublyLinkedList.head = doublyLinkedList.head.next
	}

}

func main() {
	number1 = 1
	number2 = 2
	number3 = 3
	number4 = 4
	number5 = 5
	var nodeInsert DoublyLinkedList
	nodeInsert.insertNodeIntoDoublyLinkedList(number1)
	nodeInsert.insertNodeIntoDoublyLinkedList(number2)
	nodeInsert.insertNodeIntoDoublyLinkedList(number4)
	nodeInsert.insertNodeIntoDoublyLinkedList(number5)

	nodeInsert.sortedInsert(number3)
	nodeInsert.Display()
}
