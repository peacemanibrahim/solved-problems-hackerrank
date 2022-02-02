// Link to problem statement on hackerrank: https://www.hackerrank.com/challenges/largest-rectangle/problem
// This problem is solved using both "stacks" and "queues" data structure.
// All test cases was passed on hackerrank.
package main

import "fmt"

var (
	s                           Stack
	q                           Queue
	lengthOfRectangle           int64
	currentLargestRectangleArea int64
	mainLargestRectangleArea    int64
)

type Stack []int32
type Queue []int32

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(element int32) {
	*s = append(*s, element)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	} else {
		index := len(*s) - 1
		*s = (*s)[:index]
	}
}

func (s *Stack) Peek() (int32, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		indexOfTopElement := len(*s) - 1
		topElement := (*s)[indexOfTopElement]
		return topElement, true
	}
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(element int32) {
	*q = append(*q, element)
}

func (q *Queue) Dequeue() {
	if len(*q) == 0 {
		return
	}
	*q = (*q)[1:]
}

func (q *Queue) Front() (int32, bool) {
	if len(*q) == 0 {
		return 0, false
	} else {
		frontElement := (*q)[0]
		return frontElement, true
	}
}

func FindTheNoOfHousesToTheLeftOfCurrentHouse(currentHouse int32) int64 {
	var housesToTheLeftOfCurrentHouseCount int64
	for len(s) != 0 {
		peekValue, ok := s.Peek()
		if peekValue >= currentHouse && ok == true {
			housesToTheLeftOfCurrentHouseCount++
			s.Pop()
		} else {
			break
		}
	}
	return housesToTheLeftOfCurrentHouseCount
}

func FindTheNoOfHousesToTheRightOfCurrentHouse(currentHouse int32) int64 {
	var housesToTheRightOfCurrentHouseCount int64
	for len(q) != 0 {
		frontValue, ok := q.Front()
		if frontValue >= currentHouse && ok == true {
			housesToTheRightOfCurrentHouseCount++
			q.Dequeue()
		} else {
			break
		}
	}
	return housesToTheRightOfCurrentHouseCount
}

func LargestRectangle(h []int32) int64 {
	lengthOfRectangle = 1
	for i := 1; i < len(h); i++ {
		q.Enqueue(h[i])
	}
	for i := 0; i < len(h); i++ {
		storedElementsOfStackS := s
		housesToTheLeftOfCurrentHouseCount := FindTheNoOfHousesToTheLeftOfCurrentHouse(h[i])
		lengthOfRectangle += housesToTheLeftOfCurrentHouseCount
		storedElementsOfQueueQ := q
		housesToTheRightOfCurrentHouseCount := FindTheNoOfHousesToTheRightOfCurrentHouse(h[i])
		lengthOfRectangle += housesToTheRightOfCurrentHouseCount
		currentLargestRectangleArea = int64(h[i]) * lengthOfRectangle
		if currentLargestRectangleArea > mainLargestRectangleArea {
			mainLargestRectangleArea = currentLargestRectangleArea
		}
		lengthOfRectangle = 1
		currentLargestRectangleArea = 0
		s = storedElementsOfStackS
		q = storedElementsOfQueueQ
		s.Push(h[i])
		q.Dequeue()
	}
	return mainLargestRectangleArea
}

func main() {
	h := []int32{6, 5, 2, 6, 5, 8, 7, 3, 4, 7, 9, 7, 6, 3}

	mainLargestRectangleArea = LargestRectangle(h)
	fmt.Println(mainLargestRectangleArea)
}
