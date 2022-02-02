package main

import "fmt"

func main() {
	s := []string{"m", "o", "m", "m", "y"}
	mapOfS := make(map[string]int)

	for _, v := range s {
		mapOfS[v]++
	}
	fmt.Println(mapOfS)
}
