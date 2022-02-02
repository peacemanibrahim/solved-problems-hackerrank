// Link to problem statement on hackerrank: https://www.hackerrank.com/challenges/balanced-brackets/problem
// This problem is solved using "stacks" data structure. All test cases passed on hackerrank.com.
package main

import (
	"fmt"
	"strings"
)

var (
	bracketIsBalanced bool
)

type Stack []string

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Push(str string) {
	*stack = append(*stack, str)
}

func (stack *Stack) Pop() {
	index := len(*stack) - 1
	*stack = (*stack)[:index]
}

func (stack *Stack) Peek() string {
	if stack.IsEmpty() {
		return ""
	} else {
		indexOfTopElement := len(*stack) - 1
		topElement := (*stack)[indexOfTopElement]
		return topElement
	}
}

func (stack *Stack) CompareTheTypesOfTheBrackets(str string) bool {
	if stack.IsEmpty() {
		return false
	}
	topElementInStack := stack.Peek()
	if str == ")" && topElementInStack == "(" {
		return true
	} else if str == "]" && topElementInStack == "[" {
		return true
	} else if str == "}" && topElementInStack == "{" {
		return true
	} else {
		return false
	}
}

func (stack *Stack) isBalanced(s string) string {
	arrayRepOfS := strings.Split(s, "")
	bracketIsBalanced = false
	for i := range arrayRepOfS {
		if arrayRepOfS[i] == "{" || arrayRepOfS[i] == "[" || arrayRepOfS[i] == "(" {
			stack.Push(arrayRepOfS[i])
		} else if arrayRepOfS[i] == "}" || arrayRepOfS[i] == "]" || arrayRepOfS[i] == ")" {
			bracketsAreOfSameType := stack.CompareTheTypesOfTheBrackets(arrayRepOfS[i])
			if bracketsAreOfSameType == true {
				bracketIsBalanced = true
				stack.Pop()
			} else if bracketsAreOfSameType == false {
				bracketIsBalanced = false
				break
			}
		}
	}
	if bracketIsBalanced == true && stack.IsEmpty() {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	var stack Stack
	s := "{(([])[])[]}"

	result := stack.isBalanced(s)
	fmt.Println(result)
}
