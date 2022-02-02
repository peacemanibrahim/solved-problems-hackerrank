// Link to problem statement on hackerrank: https://www.hackerrank.com/challenges/reverse-a-doubly-linked-list/problem
// This problem is solved using Doubly "Linked List" data structure. All test cases passed on hackerrank.com.
package main

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

func reverse(llist *DoublyLinkedListNode) *DoublyLinkedListNode {
	head := llist
	var lastElementOfList *DoublyLinkedListNode
	for head != nil {
		storedMemoryAddressOfNextNode := head.next
		storedMemoryAddressOfPrevNode := head.prev
		head.prev = head.next
		head.next = storedMemoryAddressOfPrevNode
		lastElementOfList = head
		head = storedMemoryAddressOfNextNode
	}
	return lastElementOfList
}

// Note that this reverse function also works like the other reverse function above.
// func reverse(llist *DoublyLinkedListNode) *DoublyLinkedListNode {
// 	head := llist
// 	var lastElementOfList *DoublyLinkedListNode
// 	for head != nil {
// 		storedMemoryAddressOfNextNode := head.next
// 		head.next = head.prev
// 		lastElementOfList = head
// 		head = storedMemoryAddressOfNextNode
// 	}
// 	return lastElementOfList
// }

func main() {
	number1 = 1
	number2 = 2
	number3 = 3
	number4 = 4
	number5 = 5
	var nodeInsert DoublyLinkedList
	nodeInsert.insertNodeIntoDoublyLinkedList(number1)
	nodeInsert.insertNodeIntoDoublyLinkedList(number2)
	nodeInsert.insertNodeIntoDoublyLinkedList(number3)
	nodeInsert.insertNodeIntoDoublyLinkedList(number4)
	nodeInsert.insertNodeIntoDoublyLinkedList(number5)

	// result := nodeInsert.reverse()

	// fmt.Println(result.data)
	// fmt.Println(result.next.data)
	// fmt.Println(result.next.next.data)
	// fmt.Println(result.next.next.next.data)
	// fmt.Println(result.next.next.next.next.data)
}
