package main

import "fmt"

func PrintFun(test int) {
	if test < 1 {
		return
	} else {
		fmt.Println(test)
		PrintFun(test - 1)
		fmt.Println(test)
	}
	return
}

func main() {
	test := 3
	PrintFun(test)
}

// Explanation: statement 9 and 10 above works by repeatedly calling and going into the PrintFun function and untill
// the base case (the statement on line 6 and 7) is reached. Note that if a "base case" is not specified the function
// run infinitely. When the base case is reached, note that there is still a line of code that has not been called in
// all the function - line 11. Therefore line 11 is called from the last call on the PrintFun function, Note: line 11
// is called from the last call on the PrintFun function (NOT first call)
