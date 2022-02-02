// Link to problem statement on hackerrank:
// https://www.hackerrank.com/challenges/find-the-merge-point-of-two-joined-linked-lists/problem
// This problem is solved using Singly "Linked Lists" data structure.
package main

import "fmt"

var (
	number1Add *int
	number2Add *int
	number3Add *int
	number4Add *int
	number5Add *int
	number6Add *int
	number7Add *int

	number1         int
	number2         int
	number3         int
	number4         int
	number5         int
	number6         int
	number7         int
	nodeInsertList1 SinglyLinkedList1
	nodeInsertList2 SinglyLinkedList2
)

type SinglyLinkedListNode struct {
	data *int
	next *SinglyLinkedListNode
}
type SinglyLinkedList1 struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

type SinglyLinkedList2 SinglyLinkedList1

func (singlyLinkedList *SinglyLinkedList1) insertNodeIntoSinglyLinkedList1(nodeData *int) {
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

func (singlyLinkedList *SinglyLinkedList2) insertNodeIntoSinglyLinkedList2(nodeData *int) {
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

func (singlyLinkedList SinglyLinkedList1) DisplayList1() {
	for singlyLinkedList.head != nil {
		fmt.Printf("%v -> ", *singlyLinkedList.head.data)
		singlyLinkedList.head = singlyLinkedList.head.next
	}

}

func findMergeNode(head1, head2 *SinglyLinkedListNode) int {
	var dataValueOfTheCommonNode int
	shouldBreak := false
	head1 = nodeInsertList1.head
	head2 = nodeInsertList2.head
	storedMemAddressOfHead2 := head2
	for head1 != nil {
		for head2 != nil {
			if head1.data == head2.data {
				dataValueOfTheCommonNode = *head1.data
				shouldBreak = true
				break
			}
			head2 = head2.next
		}
		if shouldBreak == true {
			break
		}
		head1 = head1.next
		head2 = storedMemAddressOfHead2
	}
	return dataValueOfTheCommonNode
}

func (singlyLinkedList SinglyLinkedList2) DisplayList2() {
	for singlyLinkedList.head != nil {
		fmt.Printf("%v -> ", *singlyLinkedList.head.data)
		singlyLinkedList.head = singlyLinkedList.head.next
	}

}

func main() {
	number1 = 1
	number2 = 2
	number3 = 3
	number4 = 4
	number5 = 5
	number6 = 6
	number7 = 7
	number1Add = &number1
	number2Add = &number2
	number3Add = &number3
	number4Add = &number4
	number5Add = &number5
	number6Add = &number6
	number7Add = &number7
	// var nodeInsertList1 SinglyLinkedList1
	nodeInsertList1.insertNodeIntoSinglyLinkedList1(number1Add)
	nodeInsertList1.insertNodeIntoSinglyLinkedList1(number2Add)
	nodeInsertList1.insertNodeIntoSinglyLinkedList1(number3Add)
	nodeInsertList1.insertNodeIntoSinglyLinkedList1(number4Add)
	nodeInsertList1.insertNodeIntoSinglyLinkedList1(number5Add)

	// var nodeInsertList2 SinglyLinkedList2
	nodeInsertList2.insertNodeIntoSinglyLinkedList2(number6Add)
	nodeInsertList2.insertNodeIntoSinglyLinkedList2(number7Add)
	nodeInsertList2.insertNodeIntoSinglyLinkedList2(number3Add)
	nodeInsertList2.insertNodeIntoSinglyLinkedList2(number4Add)
	nodeInsertList2.insertNodeIntoSinglyLinkedList2(number5Add)

	nodeInsertList1.DisplayList1()
	fmt.Println()
	nodeInsertList2.DisplayList2()
	fmt.Println()
	fmt.Println(*nodeInsertList1.head.next.data)
	fmt.Println(*nodeInsertList2.head.next.data)

	dataValueOfTheCommonNode := findMergeNode(nodeInsertList1.head, nodeInsertList2.head)
	fmt.Printf("the data value of the common node is: %d\n", dataValueOfTheCommonNode)
}
