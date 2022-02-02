// Link to problem statement on hackerrank: https://www.hackerrank.com/challenges/min-max-riddle/problem
// This problem is solved using the "queue" data structure. All test cases Not passed on hackerrank.
package main

import "fmt"

var (
	q                       Queue
	windowSize              int64
	minimumOfCurrentWindow  int64
	maximumOfMinimums       int64
	listOfMaximumOfMinimums []int64
)

type Queue []int64

func (q *Queue) Enqueue(element int64) {
	*q = append(*q, element)
}

func (q *Queue) Dequeue() {
	if len(*q) == 0 {
		return
	}
	*q = (*q)[1:]
}

func (q *Queue) Front() (int64, bool) {
	if len(*q) == 0 {
		return 0, false
	} else {
		frontElement := (*q)[0]
		return frontElement, true
	}
}

func EnqueueArrayElementsAndFindMaximumElement(arr []int64) int64 {
	maximumOfMinimums = arr[0]
	for i := 1; i < len(arr); i++ {
		q.Enqueue(arr[i])
		if arr[i] > maximumOfMinimums {
			maximumOfMinimums = arr[i]
		}
	}
	return maximumOfMinimums
}

func FindTheMinimumInAWindow(firstElementOfAWindow, windowSize int64) int64 {
	minimumOfCurrentWindow = firstElementOfAWindow
	for j := 0; j < int(windowSize); j++ {
		qFrontValue, ok := q.Front()
		if qFrontValue < minimumOfCurrentWindow && ok == true {
			minimumOfCurrentWindow = qFrontValue
		}
		q.Dequeue()
	}
	return minimumOfCurrentWindow
}

func riddle(arr []int64) []int64 {
	windowSize = 1
	maximumOfMinimums = EnqueueArrayElementsAndFindMaximumElement(arr)
	listOfMaximumOfMinimums = append(listOfMaximumOfMinimums, maximumOfMinimums)
	maximumOfMinimums = 0
	defaultStoredElementsOfQ := q
	for windowSize < int64(len(arr)) {
		q = defaultStoredElementsOfQ
		for i := range arr {
			storedElementsOfQ := q
			if len(q) < int(windowSize) {
				break
			}
			minimumOfCurrentWindow = FindTheMinimumInAWindow(arr[i], windowSize)
			if minimumOfCurrentWindow > maximumOfMinimums {
				maximumOfMinimums = minimumOfCurrentWindow
			}
			minimumOfCurrentWindow = 0
			q = storedElementsOfQ
			q.Dequeue()
		}
		listOfMaximumOfMinimums = append(listOfMaximumOfMinimums, maximumOfMinimums)
		maximumOfMinimums = 0
		windowSize++
	}
	return listOfMaximumOfMinimums
}

func main() {
	arr := []int64{3, 5, 4, 7, 6, 2}

	listOfMaximumOfMinimums = riddle(arr)
	fmt.Println(listOfMaximumOfMinimums)
}
