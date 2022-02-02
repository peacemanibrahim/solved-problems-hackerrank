// Link to problem statement on hackerrank: https://www.hackerrank.com/challenges/min-max-riddle/problem
// This problem is solved using the "queue" data structure.
package main

import "fmt"

var (
	q                       Queue
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

func EnqueueArrayElements(arr []int64) {
	maximumOfMinimums = arr[0]
	for i := 0; i < len(arr); i++ {
		q.Enqueue(arr[i])
	}
}

func Riddle(arr []int64) []int64 {
	EnqueueArrayElements(arr)
	listOfMinimumsOfPreviousWindowSize := arr
	listOfMinimumsOfCurrentWindowSize := []int64{}
	for len(q) != 0 {
		newArray := q
		for i := 0; i < len(newArray); i++ {
			if newArray[i] < listOfMinimumsOfPreviousWindowSize[i] {
				listOfMinimumsOfCurrentWindowSize = append(listOfMinimumsOfCurrentWindowSize, newArray[i])
				minimumOfCurrentWindow = newArray[i]
				if minimumOfCurrentWindow > maximumOfMinimums {
					maximumOfMinimums = minimumOfCurrentWindow
				}
			} else if listOfMinimumsOfPreviousWindowSize[i] < newArray[i] {
				listOfMinimumsOfCurrentWindowSize = append(listOfMinimumsOfCurrentWindowSize,
					listOfMinimumsOfPreviousWindowSize[i])
				minimumOfCurrentWindow = listOfMinimumsOfPreviousWindowSize[i]
				if minimumOfCurrentWindow > maximumOfMinimums {
					maximumOfMinimums = minimumOfCurrentWindow
				}
			} else {
				listOfMinimumsOfCurrentWindowSize = append(listOfMinimumsOfCurrentWindowSize,
					listOfMinimumsOfPreviousWindowSize[i])
				minimumOfCurrentWindow = listOfMinimumsOfPreviousWindowSize[i]
				if minimumOfCurrentWindow > maximumOfMinimums {
					maximumOfMinimums = minimumOfCurrentWindow
				}
			}
		}
		listOfMinimumsOfPreviousWindowSize = listOfMinimumsOfCurrentWindowSize
		listOfMinimumsOfCurrentWindowSize = []int64{}
		listOfMaximumOfMinimums = append(listOfMaximumOfMinimums, maximumOfMinimums)
		maximumOfMinimums = 0
		q.Dequeue()
	}
	return listOfMaximumOfMinimums
}

func main() {
	arr := []int64{353239933, 893398600, 587953779, 431392757, 716678677, 762548333, 446062962, 369841449,
		542149249, 593633419, 885154210, 473307020, 34770270, 983889405, 235786145, 539925885, 345978777,
		860225911, 86068305, 779954892, 427613068, 619071383, 797913168, 648016415, 249995795, 226132930,
		863490700, 500515140, 586661617, 233034431,
	}

	listOfMaximumOfMinimums = Riddle(arr)
	fmt.Println(listOfMaximumOfMinimums)
}
