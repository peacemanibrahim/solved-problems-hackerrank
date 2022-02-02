// Link to problem statement on hackerrank: https://www.hackerrank.com/challenges/ctci-queue-using-two-stacks/problem
// This is the implementation of the "queues" data structure using stacks. All test cases was passed on hackerrank.com
// A queue can be implemented using two stacks.
// METHOD 2: In this method, in Enqueue operation, the new element is entered at the top of stack1.
// In Dequeue operation, if stack2 is empty then all the elements are moved from stack1 to stack2 and finally
// the oldest elment which is now at the top of stack2 is returned.
// Note that Method 2 has a lower time complexity than Method 1. Therefore Method 2 is better than Method 1.
// The time complexity of the Enqueue operation using this method is O(1). The time complexity of the Dequeue
// operation using this method is O(n).
package main

import (
	"bufio"
	"fmt"	"io"
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
	s1.PushToStack1(element)
}

func Dequeue() int {
	if s2.IsEmpty() {
		for len(s1) != 0 {
			s2.PushToStack2(s1[len(s1)-1])
			s1.Pop()
		}
	}
	return s2.Pop()
}

// This function returns the element at the front of the queue
func Front() int {
	if s2.IsEmpty() {
		for len(s1) != 0 {
			s2.PushToStack2(s1[len(s1)-1])
			s1.Pop()
		}
	}
	return s2.Peek()
}

func (s2 *Stack2) IsEmpty() bool {
	return len(*s2) == 0
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
		fmt.Println("No Pop value: Stack 1 is empty")
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
		fmt.Println("No Pop value: Stack 2 is empty")
	} else {
		index := len(*s2) - 1
		element = (*s2)[index]
		*s2 = (*s2)[:index]
	}
	return element
}

func (s2 *Stack2) Peek() int {
	var topElement int
	if s2.IsEmpty() {
		fmt.Println("No Peek value: Stack 2 is empty")
	} else {
		indexOfTopElement := len(*s2) - 1
		topElement = (*s2)[indexOfTopElement]
	}
	return topElement
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
    
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

// Explanation of the code:
// 1. Enqueue() function: you push all elements to stack 1 normally. You just keep pushing elements normally.
// 2. Dequeue() function: this function starts by first checking if stack 2 is empty. We know that stack 2 is
// 		empty by Default. If empty is true then all the elements of stack 1 are pushed into stack 2 in such a
// 		way that we have the oldest element at the top of stack 2.
// 		Then after the elements of stack 1 have been moved to stack 2, now stack 2 is no more empty. So the
// 		next thing now is, since we now have the oldest added element at the top of stack 2, we just pop() from
// 		stack 2 and we get the result of the Dequeue() function.
// 		Now let us note that the only time elements from stack 1 will be moved to stack 2 is when stack 2 is empty.
// 		Now let's say we have dequeued all elements from stack 2 and now stack 2 is empty. Now that empty is true
// 		the function checks if new elements have been added to stack 1 and then pushes to stack 2 such that the
// 		oldest element is at the top of stack 2 and the process continues like that.

// Note Strongly: This method makes sure that oldest entered element is always at the top of stack 2,
// so that deQueue operation just pops from stack2.
