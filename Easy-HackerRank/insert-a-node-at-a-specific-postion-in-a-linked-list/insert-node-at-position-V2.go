// Link to the problem statement on hackerrank:
// https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list/problem
// This problem is solved using "Linked List" data structure. All the test cases passed on hackerrank.com.
package main

import "fmt"

var (
	number1 int32
	number2 int32
	number3 int32
	number4 int32
	number5 int32
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}
	singlyLinkedList.tail = node
}

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	var currentPosition int32
	currentPosition = 0
	if position == 0 {
		node := &SinglyLinkedListNode{
			data: data,
			next: llist,
		}
		return node
	} else {
		head := llist
		for head != nil {
			if position-1 == currentPosition {
				node := &SinglyLinkedListNode{
					data: data,
					next: head.next,
				}
				head.next = node
				break
			}
			currentPosition++
			head = head.next
		}
		return llist
	}
}

func (singlyLinkedList SinglyLinkedList) Display() {
	for singlyLinkedList.head != nil {
		fmt.Printf("%v -> ", singlyLinkedList.head.data)
		singlyLinkedList.head = singlyLinkedList.head.next
	}

}

func main() {
	number1 = 1
	number2 = 2
	number3 = 3
	number4 = 4
	number5 = 5
	var nodeInsert SinglyLinkedList
	nodeInsert.insertNodeIntoSinglyLinkedList(number1)
	nodeInsert.insertNodeIntoSinglyLinkedList(number2)
	nodeInsert.insertNodeIntoSinglyLinkedList(number3)
	nodeInsert.insertNodeIntoSinglyLinkedList(number4)
	nodeInsert.insertNodeIntoSinglyLinkedList(number5)

	nodeInsert.Display()
}
