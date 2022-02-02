// Link to problem statement on hackerrank: https://www.hackerrank.com/challenges/ctci-queue-using-two-stacks/problem
// This is the implementation of the "queues" data structure using stacks. All test cases NOT passed on hackerrank.com
// A queue can be implemented using two stacks.
// METHOD 1: This method makes sure that oldest entered element is always at the top of stack 1, so that
// deQueue operation just pops from stack1.
// The time complexity of the Enqueue operation using this method is O(n). The time complexity of the Dequeue
// operation using this method is O(1)
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	s1 Stack1
	s2 Stack2
)

type Stack1 []int
type Stack2 []int

func Enqueue(element int) {
	// move all elements from s1 to s2
	for len(s1) != 0 {
		s2.PushToStack2(s1[len(s1)-1])
		s1.Pop()
	}
	// push new incoming item into s1
	s1.PushToStack1(element)
	// push everything back to s1
	for len(s2) != 0 {
		s1.PushToStack1(s2[len(s2)-1])
		s2.Pop()
	}
}

func Dequeue() int {
	if len(s1) == 0 {
		fmt.Println("Queue is empty")
	}
	element := s1[len(s1)-1]
	s1.Pop()
	return element
}

// This function returns the element at the front of the queue
func Front() int {
	if len(s1) == 0 {
		fmt.Println("Queue is empty")
	}
	element := s1[len(s1)-1]
	return element
}

func (s1 *Stack1) PushToStack1(element int) {
	*s1 = append(*s1, element)
}

func (s2 *Stack2) PushToStack2(element int) {
	*s2 = append(*s2, element)
}

func (s1 *Stack1) Pop() int {
	var element int
	if len(*s1) == 0 {
		fmt.Println("Stack 1 is empty")
	} else {
		index := len(*s1) - 1
		element = (*s1)[index]
		*s1 = (*s1)[:index]
	}
	return element
}

func (s2 *Stack2) Pop() int {
	var element int
	if len(*s2) == 0 {
		fmt.Println("Stack 2 is empty")
	} else {
		index := len(*s2) - 1
		element = (*s2)[index]
		*s2 = (*s2)[:index]
	}
	return element
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	for i := 0; i < int(qTemp); i++ {
		queryTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		typeTemp, err := strconv.ParseInt(queryTemp[0], 10, 64)
		checkError(err)
		if int(typeTemp) == 1 {
			xTemp, err := strconv.ParseInt(queryTemp[1], 10, 64)
			checkError(err)
			Enqueue(int(xTemp))
		} else if int(typeTemp) == 2 {
			Dequeue()
		} else {
			result := Front()

			fmt.Fprintf(writer, "%d\n", result)
		}
	}
	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Explanation of the Enqueue() function: Firstly you push all elements from s1 into s2 and then empty s1. This is
// why you have a for loop because you are pushing more than one element.
// Then push the new incoming element to s1.
// Then after that you move back all the elements you pushed out of s1 earlier, you move them back into s1. This is
// why you have another for loop because you are pushing more than one element.

// Note Strongly: This method makes sure that oldest entered element is always at the top of stack 1,
// so that deQueue operation just pops from stack1.
